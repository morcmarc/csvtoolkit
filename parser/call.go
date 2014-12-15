package parser

import (
	"fmt"
)

type CallNode struct {
	NodeType
	Callee Node
	Args   []Node
}

func NewCallNode(args []Node) Node {
	return &CallNode{
		NodeType: NodeCall,
		Callee:   args[0],
		Args:     args[1:],
	}
}

func (this *CallNode) Copy() Node {
	call := &CallNode{
		NodeType: this.Type(),
		Callee:   this.Callee.Copy(),
		Args:     make([]Node, len(this.Args)),
	}
	for i, v := range this.Args {
		call.Args[i] = v.Copy()
	}
	return call
}

func (this *CallNode) String() string {
	args := fmt.Sprint(this.Args)
	return fmt.Sprintf("(%s %s)", this.Callee, args[1:len(args)-1])
}
