package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// Special tokens
	ItemError ItemType = iota
	ItemEOF

	// Identifiers and type literals
	ItemIdent
	ItemString
	ItemInt
	ItemFloat

	// Operators
	ItemPlus        // +
	ItemMinus       // -
	ItemMultiply    // *
	ItemDivide      // /
	ItemModulo      // %
	ItemEquals      // =
	ItemLessThan    // <
	ItemGreaterThan // >

	// Delimiters
	ItemBra        // [
	ItemKet        // ]
	ItemLeftParen  // (
	ItemRightParen // )

	// Keywords
	ItemPipe
	// -- probably have to add "." and "keys" as keywords and differentiate
	// them from identifiers.
)

const EOF = -1

// Lex returns a new Lexer
func Lex(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		items: make(chan Item),
	}
	go l.run()
	return l
}

func (l *Lexer) NextItem() Item {
	item := <-l.items
	l.lastPos = item.Pos
	return item
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// emit passes an Item back to the client.
func (l *Lexer) emit(t ItemType) {
	l.items <- Item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Item{ItemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func (l *Lexer) run() {
	for l.state = lexWhitespace; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

func lexWhitespace(l *Lexer) stateFn {
	for r := l.next(); isSpace(r) || r == '\n'; l.next() {
		r = l.peek()
	}
	l.backup()
	l.ignore()

	switch r := l.next(); {
	case r == EOF:
		l.emit(ItemEOF)
		return nil
	case r == '|':
		return lexPipe
	case r == '[':
		return lexBra
	case r == ']':
		return lexKet
	case r == '(':
		return lexLeftParen
	case r == ')':
		return lexRightParen
	case r == '"':
		return lexString
	case r == '+':
		return lexPlus
	case r == '-':
		return lexMinus
	case r == '=':
		return lexEquals
	case l.scanNumber() && (r == '+' || r == '-' || ('0' <= r && r <= '9')):
		return lexNumber
	case isIdentifier(r):
		return lexIdentifier
	default:
		panic(fmt.Sprintf("don't know what to do with: %q", r))
	}
}

func lexEquals(l *Lexer) stateFn {
	l.emit(ItemEquals)
	return lexWhitespace
}

func lexPlus(l *Lexer) stateFn {
	l.emit(ItemPlus)
	return lexWhitespace
}

func lexMinus(l *Lexer) stateFn {
	l.emit(ItemMinus)
	return lexWhitespace
}

func lexPipe(l *Lexer) stateFn {
	l.emit(ItemPipe)
	return lexWhitespace
}

func lexLeftParen(l *Lexer) stateFn {
	l.emit(ItemLeftParen)
	return lexWhitespace
}

func lexRightParen(l *Lexer) stateFn {
	l.emit(ItemRightParen)
	return lexWhitespace
}

func lexBra(l *Lexer) stateFn {
	l.emit(ItemBra)
	return lexWhitespace
}

func lexKet(l *Lexer) stateFn {
	l.emit(ItemKet)
	return lexWhitespace
}

func lexString(l *Lexer) stateFn {
	for r := l.next(); r != '"'; r = l.next() {
		if r == '\\' {
			r = l.next()
		}
		if r == EOF {
			return l.errorf("unterminated quoted string")
		}
	}
	l.emit(ItemString)
	return lexWhitespace
}

func lexNumber(l *Lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}

	if sign := l.peek(); sign == '+' || sign == '-' {
		if !l.scanNumber() {
			return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
		}
		l.emit(ItemFloat)
	} else if strings.ContainsRune(l.input[l.start:l.pos], '.') {
		l.emit(ItemFloat)
	} else {
		l.emit(ItemInt)
	}

	return lexWhitespace
}

func lexIdentifier(l *Lexer) stateFn {
	for r := l.next(); isIdentifier(r); r = l.next() {
	}
	l.backup()

	l.emit(ItemIdent)
	return lexWhitespace
}

func (l *Lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	digits := "0123456789"
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	// Next thing mustn't be alphanumeric.
	if r := l.peek(); unicode.IsLetter(r) {
		l.next()
		return false
	}
	return true
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isIdentifier reports whether r is a valid rune for an identifier.
func isIdentifier(r rune) bool {
	return r == '_' || r == '.' || unicode.IsLetter(r)
}

func debug(msg string) {
	fmt.Println(msg)
}
