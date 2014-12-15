package parser

import (
	"testing"
)

func TestThatItParsesSimpleCommands(t *testing.T) {
	p := ParseFromString("test", `ident "string"`)
	if p[0].Type() != 0 {
		t.Errorf("Was expecting NodeIdent, got: %s", p[0].Type())
	}
	if p[1].Type() != 1 {
		t.Errorf("Was expecting NodeString, got: %s", p[1].Type())
	}
}
