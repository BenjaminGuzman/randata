package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

type Mode int

const (
	APPEND Mode = iota
	OVERWRITE
)

type Format int

const (
	CSV Format = iota
	JSON
)

var validFormatsMap = map[string]Format{
	"csv":  CSV,
	"json": JSON,
}

type Config struct {
	// number of random rows to generate
	Count int

	// output file
	OutFile string

	// output format
	Format Format

	Mode Mode

	ProjectedFields map[string]bool
}

// InferFormat tries to infer the format of the file from its extension.
// An empty string is returned in case format cannot be inferred
func InferFormat(file string) string {
	ext := strings.ToLower(path.Ext(file))

	for formatStr := range validFormatsMap {
		if len(ext) >= len(formatStr) && ext[len(ext)-len(formatStr):] == formatStr {
			return formatStr
		}
	}

	return ""
}

// ValidateConfig validates configuration. nil is returned in case configuration is valid
// Call InferFormat prior to this function in case format is empty. format argument must not be empty
func ValidateConfig(count int, outFile, format, mode, fields string) error {
	// validate count
	if count < 1 {
		return fmt.Errorf("invalid argument count: must be greater than 1")
	}

	// validate mode
	if mode != "overwrite" && mode != "append" {
		return fmt.Errorf("invalid mode: \"%s\". Valid values are overwrite, append", mode)
	}

	// check format is valid
	isFormatValid := false
	for validFormatStr := range validFormatsMap {
		if format == validFormatStr {
			isFormatValid = true
			break
		}
	}
	if !isFormatValid {
		validFormatsStr := make([]string, 0, len(validFormatsMap))
		for validFormatStr := range validFormatsMap {
			validFormatsStr = append(validFormatsStr, validFormatStr)
		}

		return fmt.Errorf(
			"invalid format: \"%s\". Valid values are %s",
			format,
			strings.Join(validFormatsStr, ", "),
		)
	}

	// ensure flag compatibility
	if _, err := os.Stat(outFile); err != nil {
		if mode == "append" && errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("incompatible flag: File \"%s\" doesn't exist, however append mode was selected", outFile)
		}
	}
	// no validation regarding file permissions is made because an error will be thrown later anyway (and I'm lazy)

	return nil
}

// NewConfig creates a new Config object.
// You shall call ValidateConfig before calling this function
func NewConfig(count int, outFile, format, mode, fields string) *Config {
	config := Config{
		Count:           count,
		OutFile:         outFile,
		Format:          validFormatsMap[format], // format should be valid at this time
		Mode:            OVERWRITE,
		ProjectedFields: nil,
	}

	// extract all fields
	splittedFields := strings.Split(fields, ",")
	fieldsMap := make(map[string]bool)
	for _, field := range splittedFields {
		field = strings.TrimSpace(field)
		fieldsMap[field] = true
	}
	config.ProjectedFields = fieldsMap

	// change mode if needed
	if mode == "append" { // at this time mode should be valid
		config.Mode = APPEND
	}

	return &config
}

// Fields returns all the fields to be projected
func (conf *Config) Fields() []string {
	fields := make([]string, 0, len(conf.ProjectedFields))

	for key := range conf.ProjectedFields {
		fields = append(fields, key)
	}

	return fields
}

func (conf *Config) String() string {
	return fmt.Sprintf("Config["+
		"count=%d, "+
		"outFile=%s, "+
		"format=%d, "+
		"fields=%s, "+
		"mode=%d"+
		"]",
		conf.Count,
		conf.OutFile,
		conf.Format,
		strings.Join(conf.Fields(), ","),
		conf.Mode,
	)
}
