package parser

import (
	"testing"
)

func TestThatItParsesSimpleCommands(t *testing.T) {
	p := ParseFromString("test", `ident "string"`)
	if p[0].Type() != NodeIdent {
		t.Errorf("Was expecting NodeIdent, got: %s", p[0].Type())
	}
	if p[1].Type() != NodeString {
		t.Errorf("Was expecting NodeString, got: %s", p[1].Type())
	}
}

func TestParserParsesFunctionCalls(t *testing.T) {
	p := ParseFromString("test", `function("argument")`)
	if p[0].Type() != NodeCall {
		t.Errorf("Was expecting a call, got: %s", p[0].Type())
	}
	if p[0].String() != `function("argument")` {
		t.Errorf("Unexpected: %s", p[0].String())
	}
}
