package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_csvCreation(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "csv*.csv")
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		count   int
		outFile string
		format  string
		mode    string
		fields  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"should write all fields", args{10, tmpFile.Name(), "", "overwrite", "ALL"}},
		{"should write 10 records + header", args{10, tmpFile.Name(), "", "overwrite", "ALL"}},
	}
	// TODO write assertions
	fmt.Println(tests)
}