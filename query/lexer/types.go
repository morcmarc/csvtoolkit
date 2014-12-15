package lexer

// ItemType identifies the type of lex items.
type ItemType int

// Position vector
type Pos int

// Item represents a token returned from the scanner.
type Item struct {
	Typ ItemType
	Pos Pos
	Val string
}

// Lexer holds the state of the scanner.
type Lexer struct {
	name    string    // used for error reports
	input   string    // scanner target
	state   stateFn   // state function
	pos     Pos       // current position in the input
	start   Pos       // start position of the item
	width   Pos       // width of last rune read from input
	lastPos Pos       // last scanned position
	items   chan Item // channel of scanned items

	// parenDepth int
	// vectDepth  int
}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*Lexer) stateFn
