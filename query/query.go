package query

import (
	"fmt"
	"log"
	"os"

	"github.com/morcmarc/csvtoolkit/inferer"
	"github.com/morcmarc/csvtoolkit/utils"
)

type Query struct {
	reader  utils.CsvReader
	typeMap map[string]interface{}
}

func NewQuery(csv *os.File) *Query {
	q := &Query{
		reader:  utils.NewDefaultCSVReader(csv),
		typeMap: make(map[string]interface{}),
	}

	fields := Keys(q.reader)
	typeMap, err := inferer.Infer(q.reader, fields, 10)
	if err != nil {
		log.Fatalf("Could not infer types: %s", err)
	}
	q.typeMap = typeMap

	return q
}

func (q *Query) Run(qs string) {
	q.reader.Reset()
	switch qs {
	case "keys":
		r := Keys(q.reader)
		fmt.Println(r)
		break
	}
}
