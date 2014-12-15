package parser

import (
	"fmt"
)

type CallNode struct {
	NodeType
	Callee Node
	Args   []Node
}

func NewCallNode(callee Node, args []Node) Node {
	return &CallNode{
		NodeType: NodeCall,
		Callee:   callee,
		Args:     args,
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
	return fmt.Sprintf("f: %s, a: (%s)", this.Callee, args[1:len(args)-1])
}
