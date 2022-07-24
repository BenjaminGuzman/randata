package output

import (
	"encoding/json"
	"github.com/BenjaminGuzman/randata/config"
	"os"
)

type JSONOutput struct {
	encoder *json.Encoder
	file    *os.File
}

func NewJSONOutput(conf *config.Config) *JSONOutput {
	file := openCreateOrFail(conf.OutFile, conf.Mode)
	out := &JSONOutput{
		file:    file,
		encoder: json.NewEncoder(file),
	}

	return out
}

func (out *JSONOutput) WriteRow(row map[string]string) error {
	return out.encoder.Encode(row)
}

func (out *JSONOutput) Close() error {
	return out.file.Close()
}
