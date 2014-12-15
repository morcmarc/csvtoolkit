package parser

import (
	"github.com/morcmarc/csvtoolkit/lexer"
)

type NodeType int

type Node interface {
	Type() NodeType
	String() string
	Copy() Node
}

const (
	NodeIdent NodeType = iota
	NodeString
	NodeNumber
	NodeCall
)

func (this NodeType) Type() NodeType {
	return this
}

func ParseFromString(name, program string) []Node {
	return Parse(lexer.Lex(name, program))
}

func Parse(l *lexer.Lexer) []Node {
	return parser(l, make([]Node, 0), ' ')
}

func parser(l *lexer.Lexer, tree []Node, lookingFor rune) []Node {
	for item := l.NextItem(); item.Typ != lexer.ItemEOF; {
		switch t := item.Typ; t {
		case lexer.ItemIdent:
			tree = append(tree, NewIdentNode(item.Val))
		case lexer.ItemString:
			tree = append(tree, NewStringNode(item.Val))
		// case lexer.ItemInt:
		// 	tree = append(tree, newIntNode(item.Value))
		// case lexer.ItemFloat:
		// 	tree = append(tree, newFloatNode(item.Value))
		// case lexer.ItemComplex:
		// 	tree = append(tree, newComplexNode(item.Value))
		// case lexer.ItemLeftParen:
		// 	tree = append(tree, newCallNode(parser(l, make([]Node, 0), ')')))
		// case lexer.ItemLeftVect:
		// 	tree = append(tree, newVectNode(parser(l, make([]Node, 0), ']')))
		// case lexer.ItemRightParen:
		// 	if lookingFor != ')' {
		// 		panic(fmt.Sprintf("unexpected \")\" [%d]", item.Pos))
		// 	}
		// 	return tree
		// case lexer.ItemRightVect:
		// 	if lookingFor != ']' {
		// 		panic(fmt.Sprintf("unexpected \"]\" [%d]", item.Pos))
		// 	}
		// 	return tree
		case lexer.ItemError:
			println(item.Val)
		default:
			panic("Bad Item type")
		}
		item = l.NextItem()
	}

	return tree
}
