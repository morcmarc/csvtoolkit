package query

import (
	"fmt"
	"os"
)

type Query struct {
	input   *os.File
	typeMap map[string]interface{}
}

func NewQuery(csv *os.File) *Query {
	q := &Query{
		input:   csv,
		typeMap: make(map[string]interface{}),
	}
	return q
}

func (q *Query) Run(qs string) {
	csvReader := NewDefaultCSVReader(q.input)

	switch qs {
	case "keys":
		r := Keys(csvReader)
		fmt.Println(r)
		break
	}
}
