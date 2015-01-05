package query

import (
	"github.com/morcmarc/csvtoolkit/query/parser"
)

const (
	FuncKeys = "keys"
	FuncHas  = "has"
	FuncDot  = "."
)

// Returns the name of all attributes that the given Row object has
func Keys(row Row) []string {
	keys := []string{}
	for p, _ := range row {
		keys = append(keys, p)
	}
	return keys
}

// `HasProperty` returns whether the input object has the given key
func HasProperty(row Row, idx string) bool {
	_, ok := row[idx]
	return ok
}

// `HasIndex` returns whether the input array has an element at the given index.
func HasIndex(rows []Row, idx int) bool {
	if len(rows) > idx && idx > -1 {
		return true
	}
	return false
}

func At(rows []Row, idx int) Row {
	return rows[idx]
}

func Property(row Row, prop string) interface{} {
	return row[prop]
}

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

func isHasFunc(node *parser.CallNode) bool {
	if node.Callee.String() == FuncHas {
		return true
	}
	return false
}
