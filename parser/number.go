package parser

import (
	"go/token"
)

type NumberNode struct {
	NodeType
	Value      string
	NumberType token.Token
}

func NewIntNode(val string) *NumberNode {
	return &NumberNode{
		NodeType:   NodeNumber,
		Value:      val,
		NumberType: token.INT,
	}
}

func NewFloatNode(val string) *NumberNode {
	return &NumberNode{
		NodeType:   NodeNumber,
		Value:      val,
		NumberType: token.FLOAT,
	}
}

func (this *NumberNode) Copy() Node {
	return &NumberNode{
		NodeType:   this.Type(),
		Value:      this.Value,
		NumberType: this.NumberType,
	}
}

func (this *NumberNode) String() string {
	return this.Value
}
