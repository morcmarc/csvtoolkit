package query

import (
	"encoding/csv"
	"os"
)

type CsvReader interface {
	Reset()
	Read() ([]string, error)
}

type DefaultCSVReader struct {
	CsvReader
	input         *os.File
	builtinReader *csv.Reader
}

func NewDefaultCSVReader(i *os.File) *DefaultCSVReader {
	d := &DefaultCSVReader{
		input:         i,
		builtinReader: csv.NewReader(i),
	}
	return d
}

func (d *DefaultCSVReader) Read() ([]string, error) {
	return d.builtinReader.Read()
}

func (d *DefaultCSVReader) Reset() {
	d.input.Seek(0, 0)
	d.builtinReader = csv.NewReader(d.input)
}
