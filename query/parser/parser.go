package parser

import (
	"fmt"

	"github.com/morcmarc/csvtoolkit/query/lexer"
)

type NodeType int

type Node interface {
	Type() NodeType
	String() string
	Copy() Node
}

const (
	NodeIdent NodeType = iota
	NodeIndex
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
		case lexer.ItemEOF:
			if lookingFor != ' ' {
				panic(fmt.Sprintf("Unexpected end of input, was expecting: %s", lookingFor))
			}
		case lexer.ItemIdent:
			tree = append(tree, NewIdentNode(item.Val))
		case lexer.ItemString:
			tree = append(tree, NewStringNode(item.Val))
		case lexer.ItemInt:
			tree = append(tree, NewIntNode(item.Val))
		case lexer.ItemFloat:
			tree = append(tree, NewFloatNode(item.Val))
		case lexer.ItemLeftParen:
			// Previous node is identifier
			if tree[len(tree)-1].Type() == NodeIdent {
				t := parser(l, make([]Node, 0), ')')
				tree[len(tree)-1] = NewCallNode(tree[len(tree)-1], t)
			} else {
				panic("Was expecting identifier before function call")
			}
		case lexer.ItemBra:
			if len(tree) > 0 {
				t := parser(l, make([]Node, 0), ']')
				if len(t) > 1 {
					panic("Invalid index")
				}
				tree[len(tree)-1] = NewIndexNode(tree[len(tree)-1], t[0])
			} else {
				panic(fmt.Sprintf("unexpected \"[\" [%d]", item.Pos))
			}
		case lexer.ItemRightParen:
			if lookingFor != ')' {
				panic(fmt.Sprintf("unexpected \")\" [%d]", item.Pos))
			}
			return tree
		case lexer.ItemKet:
			if lookingFor != ']' {
				panic(fmt.Sprintf("unexpected \"]\" [%d]", item.Pos))
			}
			return tree
		case lexer.ItemError:
			fmt.Println(item.Val)
		default:
			panic("Bad Item type")
		}
		item = l.NextItem()
	}

	return tree
}
