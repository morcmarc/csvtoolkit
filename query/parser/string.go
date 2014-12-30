package parser

import (
	"regexp"
)

type StringNode struct {
	NodeType
	Value string
}

func NewStringNode(val string) *StringNode {
	re := regexp.MustCompile(`"`)
	pv := re.ReplaceAllString(val, "")
	return &StringNode{
		NodeType: NodeString,
		Value:    pv,
	}
}

func (this *StringNode) Copy() Node {
	return NewStringNode(this.Value)
}

func (this *StringNode) String() string {
	return this.Value
}
