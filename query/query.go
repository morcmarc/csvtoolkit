package query

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/morcmarc/csvtoolkit/converter"
	"github.com/morcmarc/csvtoolkit/inferer"
	"github.com/morcmarc/csvtoolkit/query/parser"
	"github.com/morcmarc/csvtoolkit/utils"
)

type Row map[string]interface{}

type Query struct {
	reader     *utils.DefaultCSVReader
	schema     *converter.Schema
	typeMap    Row
	output     []Row
	dataBuffer []Row
}

const (
	FuncKeys = "keys"
)

func NewQuery(csv *os.File) *Query {
	q := &Query{
		reader:     utils.NewDefaultCSVReader(csv),
		typeMap:    Row{},
		dataBuffer: []Row{},
		output:     []Row{},
	}

	fields := Keys(q.reader)
	typeMap, err := inferer.Infer(q.reader, fields, 10)
	if err != nil {
		log.Fatalf("Could not infer types: %s", err)
	}
	q.typeMap = typeMap
	q.schema = converter.NewSchema(fields, typeMap)

	return q
}

func (q *Query) Run(qs string) {
	tree := parser.ParseFromString("query", qs)

	// First, load in data
	for {
		line, err := q.reader.Read()
		if err == io.EOF {
			break
		}
		q.dataBuffer = append(q.dataBuffer, q.schema.Convert(line))
	}

	// Run query against the data
	for _, n := range tree {
		q.evalNode(n)
	}

	q.output = q.dataBuffer
	fmt.Printf("%s\n", q.output)
}

func (q *Query) evalNode(node parser.Node) {
	switch node.Type() {
	case parser.NodeCall:
		q.evalFuncCall(node.(*parser.CallNode))
		break
	case parser.NodeIndex:
		q.evalIndex(node.(*parser.IndexNode))
		break
	case parser.NodePipe:
		q.evalPipe()
		break
	}
}

func (q *Query) evalPipe() {
	q.output = q.dataBuffer
}

func (q *Query) evalFuncCall(node *parser.CallNode) {
	switch {
	case isKeysFunc(node):
		// q.dataBuffer = Keys(r)
		break
	}
}

func (q *Query) evalIndex(node *parser.IndexNode) {
	if node.Container.Type() != parser.NodeIdent {
		q.evalNode(node.Container)
	}
	if node.Index.Type() == parser.NodeNumber {
		i := node.Index.(*parser.NumberNode)
		idx, err := strconv.Atoi(i.Value)
		if err != nil {
			panic("Invalid index")
		}
		v := q.dataBuffer[idx]
		q.dataBuffer = []Row{v}
	}
}

func isKeysFunc(node *parser.CallNode) bool {
	if node.Callee.String() == FuncKeys {
		return true
	}
	return false
}
