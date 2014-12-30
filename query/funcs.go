package query

import (
	"reflect"

	"github.com/morcmarc/csvtoolkit/query/parser"
)

const (
	FuncKeys = "keys"
	FuncDot  = "."
)

func isKeysFunc(node *parser.CallNode) bool {
	if node.Callee.String() == FuncKeys {
		return true
	}
	return false
}

func isDotFunc(node *parser.CallNode) bool {
	if node.Callee.String() == FuncDot {
		return true
	}
	return false
}

func Keys(row Row) []string {
	keys := []string{}
	for p, _ := range row {
		keys = append(keys, p)
	}
	return keys
}

func Dot(data, idx interface{}) interface{} {
	switch reflect.TypeOf(idx) {
	case reflect.TypeOf(int(1)):
		return At(data.([]Row), idx.(int))
	case reflect.TypeOf(string("a")):
		return Property(data.(Row), idx.(string))
	case nil:
		return data
	}
	return nil
}

func At(rows []Row, idx int) Row {
	return rows[idx]
}

func Property(row Row, prop string) interface{} {
	return row[prop]
}
