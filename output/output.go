package output

import (
	"encoding/csv"
	"log"
	"os"
)

type Output interface {
	writeRow(map[string]string) error
}

type AbstractOutput struct {
	file   *os.File
	format main.Format
}

func openCreateOrFail(file string) *os.File {
	f, err := os.OpenFile(file, os.O_RDWR, 066)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func NewOutput(config *main.Config) *Output {
	f, err := os.OpenFile(config.OutFile, os.O_RDWR, 066)
	if err != nil {
		log.Fatal(err)
	}

	out := &Output{
		file:   f,
		format: config.Format,
	}

	switch config.Format {
	case main.JSON: // TODO
		break
	case main.CSV:
		out.csvWriter = csv.NewWriter(out.file)
	}

}

func (o *Output) writeRow(row map[string]string) {

}
