package query

import (
	"encoding/csv"
	"os"
)

type CsvReader interface {
	Read() ([]string, error)
}

type DefaultCSVReader struct {
	CsvReader
	builtinReader *csv.Reader
}

func NewDefaultCSVReader(i *os.File) *DefaultCSVReader {
	d := &DefaultCSVReader{
		builtinReader: csv.NewReader(i),
	}
	return d
}

func (d *DefaultCSVReader) Read() ([]string, error) {
	return d.builtinReader.Read()
}
