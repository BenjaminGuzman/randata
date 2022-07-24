package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func configFlags(count *int, outFile, format, mode, fields *string) {
	flag.IntVar(count, "count", 100, "Number of random records (rows) to generate")
	flag.StringVar(
		format,
		"format",
		"",
		"Output format. If not given it'll be inferred from the output file extension (as long as it's valid). ",
	)
	flag.StringVar(
		outFile,
		"file",
		"data.csv",
		"Output file. You can choose to append or overwrite with the --mode flag",
	)
	flag.StringVar(
		mode,
		"mode",
		"overwrite",
		"Decide how to update the output file. Valid values are: overwrite, append",
	)
	flag.StringVar(
		fields,
		"fields",
		"ALL",
		"Comma separated fields to project. These are ignored in append mode. "+
			"Special value is ALL which is equivalent to selecting all fields. "+
			"If you need a field that is not one of the pre-defined, you could provide its name followed by its format",
	)
}

// calls configFlags(), parses, validates and creates a new Config object.
// May return nil if some config is invalid. You don't need to log any error to notify user
func initConfig() *Config {
	var count int
	var outFile, format, mode, fields string

	configFlags(&count, &outFile, &format, &mode, &fields)
	flag.Parse()

	// normalize string arguments
	outFile = strings.TrimSpace(outFile)
	format = strings.ToLower(strings.TrimSpace(format))
	mode = strings.ToLower(strings.TrimSpace(mode))
	fields = strings.TrimSpace(fields)

	// infer format in case it was not given
	if format == "" {
		format = InferFormat(outFile)
		if format == "" {
			log.Printf("couldn't infer format from file %s\n", outFile)
			return nil
		}
	}

	// validate args
	if err := ValidateConfig(count, outFile, format, mode, fields); err != nil {
		log.Println(err)
		return nil
	}

	return NewConfig(count, outFile, format, mode, fields)
}

func main() {
	config := initConfig()
	if config == nil {
		return
	}

	// TODO generate random data using the given configuration
	fmt.Println(config)
}
