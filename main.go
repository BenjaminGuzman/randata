package main

import (
	"flag"
	"fmt"
	"github.com/BenjaminGuzman/randata/config"
	"github.com/BenjaminGuzman/randata/generator"
	"github.com/BenjaminGuzman/randata/output"
	"log"
	"strings"
)

func configFlags(count *int, outFile, format, mode, fields *string) {
	flag.IntVar(count, "count", 100, "Number of random records (rows) to generate")
	flag.StringVar(
		format,
		"format",
		"",
		"Output format. If not given it'll be inferred from the output file extension (as long as it's valid).\n"+
			"Valid formats: csv, json",
	)
	flag.StringVar(
		outFile,
		"out",
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
		"Comma separated fields to project. These are ignored in append mode.\n "+
			"ALL is a special value, and it is equivalent to selecting all fields.\n "+
			"If you need a field that is not one of the pre-defined, you can provide its name followed by its format\n "+
			"(separated by a colon, e.g. customField:### ??**)\n "+
			"To specify the format you can use the wildcards '?' for a random letter, '#' for a random number, "+
			"and '*' for a random ASCII char",
	)
}

// calls configFlags(), parses, validates and creates a new Config object.
// if some config is invalid, a fatal error will be logged
func initConfig() *config.Config {
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
		format = config.InferFormat(outFile)
		if format == "" {
			log.Fatalf("couldn't infer format from file %s\n", outFile)
		}
	}

	// validate args
	if err := config.ValidateConfig(count, outFile, format, mode, fields); err != nil {
		log.Fatalln(err)
	}

	return config.NewConfig(count, outFile, format, mode, fields)
}

func main() {
	conf := initConfig()

	fmt.Println("Using configuration:", conf)

	// prepare output depending on the format
	var out output.Output
	switch conf.Format {
	case config.CSV:
		csvOut := output.NewCSVOutput(conf)
		if conf.Mode == config.APPEND {
			conf.SetProjectedFields(csvOut.Headers)
		}
		out = csvOut
		break
	case config.JSON:
		out = output.NewJSONOutput(conf)
		break
	}
	defer out.Close()

	// generate and write random data
	gen := generator.NewGenerator()
	for i := 0; i < conf.Count; i++ {
		err := out.WriteRow(gen.GenerateRow(conf.ProjectedFields))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Successfully generated %d random records. Saved to %s\n", conf.Count, conf.OutFile)
}
