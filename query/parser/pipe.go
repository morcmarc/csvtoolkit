package parser

type PipeNode struct {
	NodeType
	Value string
}

func NewPipeNode(val string) *PipeNode {
	return &PipeNode{
		NodeType: NodeString,
		Value:    val,
	}
}

func (this *PipeNode) Copy() Node {
	return NewPipeNode(this.Value)
}

func (this *PipeNode) String() string {
	return this.Value
}
