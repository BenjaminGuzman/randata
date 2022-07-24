package output

import (
	"encoding/csv"
	"github.com/BenjaminGuzman/randata/config"
	"log"
	"os"
)

type CSVOutput struct {
	file    *os.File
	writer  *csv.Writer
	Headers []string
}

func NewCSVOutput(conf *config.Config) *CSVOutput {
	file := openCreateOrFail(conf.OutFile, conf.Mode)
	out := &CSVOutput{
		file:   file,
		writer: csv.NewWriter(file),
	}

	if conf.Mode == config.APPEND {
		out.Headers = out.readHeaders()
	} else {
		out.Headers = conf.Fields()

		err := out.writer.Write(out.Headers)
		if err != nil {
			log.Fatal(err)
		}
	}

	return out
}

func (out *CSVOutput) WriteRow(row map[string]string) error {
	values := make([]string, 0, len(row))

	for _, header := range out.Headers {
		values = append(values, row[header])
	}

	err := out.writer.Write(values)
	if err != nil {
		return err
	}

	return nil
}

// function must be called only once. Calling it twice will produce panic.
// This is because this function is non-deterministic
func (out *CSVOutput) readHeaders() []string {
	if out.Headers != nil {
		log.Panicln("Headers were read already!")
	}

	// we could use file.Seek to ensure we always read from the beginning of the file and thus making this deterministic
	// but that function is also non-deterministic in case file was opened with O_APPEND, so...
	headers, err := csv.NewReader(out.file).Read()
	if err != nil {
		log.Fatal(err)
	}
	return headers
}

func (out *CSVOutput) Close() error {
	out.writer.Flush()
	return out.file.Close()
}
