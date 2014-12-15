package parser

type StringNode struct {
	NodeType
	Value string
}

func NewStringNode(val string) *StringNode {
	return &StringNode{
		NodeType: NodeString,
		Value:    val,
	}
}

func (this *StringNode) Copy() Node {
	return NewStringNode(this.Value)
}

func (this *StringNode) String() string {
	return this.Value
}
