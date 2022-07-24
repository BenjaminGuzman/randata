package output

import (
	"github.com/BenjaminGuzman/randata/config"
	"log"
	"os"
)

type Output interface {
	WriteRow(map[string]string) error
	Close() error
}

func openCreateOrFail(file string, mode config.Mode) *os.File {
	var f *os.File
	var err error
	if mode == config.OVERWRITE {
		f, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
		if err == nil {
			err = os.Truncate(file, 0) // delete file contents (if any)
		}
	} else {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_RDWR, 0644)
	}
	if err != nil {
		log.Fatal(err)
	}

	return f
}
