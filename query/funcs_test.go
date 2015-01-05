package query

import (
	"testing"

	"github.com/morcmarc/csvtoolkit/query/parser"
)

func TestIsKeyFunc(t *testing.T) {
	a := parser.ParseFromString("query", "keys()")
	if !isKeysFunc(a[0].(*parser.CallNode)) {
		t.Errorf("Was expecting true, got false")
	}
	b := parser.ParseFromString("query", "anotherfunc()")
	if isKeysFunc(b[0].(*parser.CallNode)) {
		t.Errorf("Was expecting true, got false")
	}
}

func TestIsDotFunc(t *testing.T) {
	a := parser.ParseFromString("query", ".(2)")
	if !isDotFunc(a[0].(*parser.CallNode)) {
		t.Errorf("Was expecting true, got false")
	}
	b := parser.ParseFromString("query", "anotherfunc()")
	if isDotFunc(b[0].(*parser.CallNode)) {
		t.Errorf("Was expecting true, got false")
	}
}

func TestIsHasFunc(t *testing.T) {
	a := parser.ParseFromString("query", "has(2)")
	if !isHasFunc(a[0].(*parser.CallNode)) {
		t.Errorf("Was expecting true, got false")
	}
	b := parser.ParseFromString("query", "anotherfunc()")
	if isHasFunc(b[0].(*parser.CallNode)) {
		t.Errorf("Was expecting true, got false")
	}
}

var (
	rA     Row   = Row{"A": 1, "B": "rA"}
	rB     Row   = Row{"A": 2, "B": "rB"}
	rC     Row   = Row{"A": 3, "B": "rC"}
	rD     Row   = Row{"A": 4, "B": "rD"}
	cursor Row   = rB
	rows   []Row = []Row{rA, rB, rC, rD}
)

func TestKeysFunction(t *testing.T) {
	keys := Keys(cursor)
	if len(keys) != len(rB) {
		t.Errorf("Was expecting %d properties, got %d", len(rB), len(keys))
	}
}
