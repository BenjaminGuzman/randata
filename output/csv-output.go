package output

import (
	"encoding/csv"
)

type CSVOutput struct {
	writer *csv.Writer
}

func NewCSVOutput(config *main.Config) *CSVOutput {

}

func (out *CSVOutput) WriteRow(row map[string]string) error {
	// TODO
	return nil
}
