package query

import (
	"fmt"
	"os"
)

type Query struct {
	reader  CsvReader
	typeMap map[string]interface{}
}

func NewQuery(csv *os.File) *Query {
	q := &Query{
		reader:  NewDefaultCSVReader(csv),
		typeMap: make(map[string]interface{}),
	}
	return q
}

func (q *Query) Run(qs string) {
	switch qs {
	case "keys":
		r := Keys(q.reader)
		fmt.Println(r)
		break
	}
}
