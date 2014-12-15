package parser

type IdentNode struct {
	NodeType
	Ident string
}

func NewIdentNode(name string) *IdentNode {
	return &IdentNode{
		NodeType: NodeIdent,
		Ident:    name,
	}
}

func (this *IdentNode) Copy() Node {
	return NewIdentNode(this.Ident)
}

func (this *IdentNode) String() string {
	return this.Ident
}
