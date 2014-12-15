package lexer

import (
	"testing"
)

func TestLexerRecognizesEOF(t *testing.T) {
	l := Lex("query", ".")
	l.NextItem()
	eof := l.NextItem()
	if eof.Typ != ItemEOF {
		t.Errorf("Was expecting EOF, got: %s", eof.Typ)
	}
}

func TestLexerRecognizesDotAsIdentifier(t *testing.T) {
	l := Lex("query", ".")
	i := l.NextItem()
	if i.Val != "." {
		t.Errorf("Was expecting '.', got: %s", i.Val)
	}
	if i.Typ != ItemIdent {
		t.Errorf("Was expecting Dot, got: %s", i.Typ)
	}
}

func TestLexerRecognizesKeysAsIdentifier(t *testing.T) {
	l := Lex("query", "keys")
	i := l.NextItem()
	if i.Val != "keys" {
		t.Errorf("Was expecting 'keys', got: %s", i.Val)
	}
	if i.Typ != ItemIdent {
		t.Errorf("Was expecting Keys, got: %s", i.Typ)
	}
}

func TestLexerRecognizesAString(t *testing.T) {
	l := Lex("query", `"I am a string"`)
	i := l.NextItem()
	if i.Val != `"I am a string"` {
		t.Errorf(`Was expecting "I am a string", got: %s`, i.Val)
	}
	if i.Typ != ItemString {
		t.Errorf("Was expecting String, got: %s", i.Typ)
	}
}

func TestLexerThrowsErrorIfStringIsUnterminated(t *testing.T) {
	l := Lex("query", `"I am a string`)
	i := l.NextItem()
	if i.Typ != ItemError {
		t.Errorf("Was expecting error")
	}
	if i.Val != "unterminated quoted string" {
		t.Errorf(`Was expecting unterminated quoted string error, got: %s`, i.Val)
	}
}

func TestLexerRecognizesBraAndKet(t *testing.T) {
	l := Lex("query", `.["Attribute-Name"]`)
	l.NextItem()
	b := l.NextItem()
	l.NextItem()
	r := l.NextItem()
	if b.Typ != ItemBra {
		t.Errorf("Was expecting Bra, got: %s", b.Typ)
	}
	if r.Typ != ItemKet {
		t.Errorf("Was expecting Ket, got: %s", r.Typ)
	}
}

func TestLexerIgnoresWhitespaces(t *testing.T) {
	l := Lex("query", `   .   [ "Attribute-Name"]     `)
	i := l.NextItem()
	if i.Typ != ItemIdent {
		t.Error("Was expecting a Dot")
	}
	b := l.NextItem()
	if b.Typ != ItemBra {
		t.Error("Was expecting a Bra")
	}
}
