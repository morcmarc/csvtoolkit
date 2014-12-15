package parser

import (
	"fmt"
)

type VectorNode struct {
	NodeType
	Nodes []Node
}

func NewVectNode(content []Node) *VectorNode {
	return &VectorNode{
		NodeType: NodeVector,
		Nodes:    content,
	}
}

func (this *VectorNode) Copy() Node {
	vect := &VectorNode{
		NodeType: this.Type(),
		Nodes:    make([]Node, len(this.Nodes)),
	}
	for i, v := range this.Nodes {
		vect.Nodes[i] = v.Copy()
	}
	return vect
}

func (this *VectorNode) String() string {
	return fmt.Sprint(this.Nodes)
}
