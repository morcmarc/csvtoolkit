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
	f := p[0].(*CallNode)
	a := f.Args[0].(*StringNode)
	if a.Value != `"argument"` {
		t.Errorf("Unexpected: %s", a)
	}
}

func TestParserParsesIndexesOnIdentifiers(t *testing.T) {
	p := ParseFromString("test", `.[1]`)
	vn := p[0].(*IndexNode)
	_ = vn.Container.(*IdentNode)
	i := vn.Index.(*NumberNode)
	if i.Value != `1` {
		t.Errorf(`Was expecting "a", got: %s`, i.Value)
	}
}

func TestParserParsesIndexesOnFunctions(t *testing.T) {
	p := ParseFromString("test", `map("prop-name")[123]`)
	vn := p[0].(*IndexNode)
	_ = vn.Container.(*CallNode)
	i := vn.Index.(*NumberNode)
	if i.Value != `123` {
		t.Errorf(`Was expecting "a", got: %s`, i.Value)
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
