package generator

import (
	"fmt"
	"github.com/jaswdr/faker"
	"log"
	"strings"
	"time"
)

type Generator struct {
	faker faker.Faker
}

func NewGenerator() *Generator {
	return &Generator{
		faker: faker.New(),
	}
}

func parseCustomField(custom string) (string, string, error) {
	splitIdx := strings.Index(custom, ":")
	if splitIdx == -1 {
		return "", "", fmt.Errorf("field \"%s\" couldn't be parsed. ':' to specify format is missing", custom)
	}

	return custom[:splitIdx], custom[splitIdx+1:], nil
}

// modified implementation of Faker.Asciify
func (g *Generator) asciify(str string) (out string) {
	for _, c := range strings.Split(str, "") {
		if c == "*" {
			c = fmt.Sprintf("%c", g.faker.IntBetween(32, 126))
		}
		out = out + c
	}

	return out
}

var PROJECTED_FIELDS_AVAILABLE = []string{
	"firstName",
	"lastName",
	"username",
	"password",
	"address",
	"streetAddress",
	"state",
	"city",
	"country",
	"phone",
	"ssn",
	"dob",
	"zipCode",
}

// returns the field name and field value
func (g *Generator) genField(field string) (string, string) {
	switch field {
	case "firstName":
		return field, g.faker.Person().FirstName()
	case "lastName":
		return field, g.faker.Person().LastName()
	case "username":
		return field, g.faker.Internet().User()
	case "password":
		return field, g.faker.Internet().Password()
	case "address":
		return field, strings.Replace(g.faker.Address().Address(), "%", "#", 1) // faker uses % instead of #
	case "streetAddress":
		return field, strings.Replace(g.faker.Address().StreetAddress(), "%", "#", 1)
	case "state":
		return field, g.faker.Address().State()
	case "city":
		return field, g.faker.Address().City()
	case "country":
		return field, g.faker.Address().Country()
	case "phone":
		return field, g.faker.Phone().Number()
	case "ssn":
		return field, g.faker.Person().SSN()
	case "dob": // date of birth
		minAge := 18
		maxAge := 120
		today := time.Now()
		maxDate := today.AddDate(-minAge, 0, 0)
		minDate := today.AddDate(-maxAge, 0, 0)
		return field, g.faker.Time().TimeBetween(minDate, maxDate).Format(time.RFC3339)
	case "zipCode":
		return field, g.faker.Address().PostCode()
	default:
		fieldName, fieldFormat, err := parseCustomField(field)
		if err != nil {
			log.Fatal(err)
		}
		return fieldName, g.asciify(g.faker.Bothify(fieldFormat))
	}
}

func (g *Generator) GenerateRow(projectedFields map[string]bool) map[string]string {
	row := make(map[string]string)

	for field, b := range projectedFields {
		if !b {
			continue
		}
		fieldName, fieldValue := g.genField(field)
		row[fieldName] = fieldValue
	}

	return row
}
