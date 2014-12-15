package parser

import (
	"fmt"
)

type IndexNode struct {
	NodeType
	Container Node
	Index     Node
}

func NewIndexNode(container, index Node) *IndexNode {
	return &IndexNode{
		NodeType:  NodeIndex,
		Container: container,
		Index:     index,
	}
}

func (this *IndexNode) Copy() Node {
	vect := &IndexNode{
		NodeType:  this.Type(),
		Container: this.Container.Copy(),
		Index:     this.Index.Copy(),
	}
	return vect
}

func (this *IndexNode) String() string {
	return fmt.Sprint("%s[%s]", this.Container, this.Index)
}
