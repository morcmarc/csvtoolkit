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

func TestParserParsesVectors(t *testing.T) {
	p := ParseFromString("test", `.["a"]`)
	if p[0].Type() != NodeIdent {
		t.Errorf("Was expecting identifier, got: %s", p[0].Type())
	}
	if p[1].Type() != NodeVector {
		t.Errorf("Was expecting vector, got: %s", p[1].Type())
	}
	vn := p[1].(*VectorNode)
	sn := vn.Nodes[0].(*StringNode)
	if sn.Value != `"a"` {
		t.Errorf(`Was expecting "a", got: %s`, sn.Value)
	}
}

func TestParserHandlesNestedFunctionArguments(t *testing.T) {
	p := ParseFromString("test", `map(map("a"))`)
	c1n := p[0].(*CallNode)
	c2n := c1n.Args[0].(*CallNode)
	a := c2n.Args[0].(*StringNode)
	if a.Value != `"a"` {
		t.Errorf(`Was expecting "a", got: %s`, a.Value)
	}
}
