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
	dataBuffer []Row
	cursor     Row
}

func NewQuery(csv *os.File) *Query {
	q := &Query{
		reader:     utils.NewDefaultCSVReader(csv),
		typeMap:    Row{},
		dataBuffer: []Row{},
	}

	fields, err := q.reader.Read()
	if err != nil {
		log.Fatalf("Could not read: %s", err)
	}
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

	var prev interface{} = q.dataBuffer
	// Run query against the data
	for _, n := range tree {
		prev = q.evalNode(n, prev)
	}

	fmt.Printf("%v\n", prev)
}

func (q *Query) evalNode(node parser.Node, prev interface{}) interface{} {
	switch node.Type() {
	case parser.NodeCall:
		return q.evalFuncCall(node.(*parser.CallNode), prev)
	case parser.NodeIndex:
		return q.evalIndex(node.(*parser.IndexNode), prev)
	case parser.NodePipe:
		return prev
	}
	return prev
}

func (q *Query) evalFuncCall(node *parser.CallNode, prev interface{}) interface{} {
	switch {
	case isKeysFunc(node):
		return Keys(prev.(Row))
	case isDotFunc(node):
		return processIndex(prev.([]Row), node.Args[0])
	case isHasFunc(node):
		return processHas(prev, node.Args[0])
	}
	return nil
}

func (q *Query) evalIndex(node *parser.IndexNode, prev interface{}) interface{} {
	var p interface{} = prev
	// Resolve "nested" indexing, i.e.: .[0]["Property"] first have to get the
	// value of .[0] and then execute ["Property"] on that result instead
	if node.Container.Type() != parser.NodeIdent {
		p = q.evalNode(node.Container, prev)
	}
	if p == nil {
		return nil
	}
	return processIndex(p, node.Index)
}

// @TODO: figure out how to merge processIndex and processHas nicely
func processIndex(data interface{}, idx parser.Node) interface{} {
	if idx.Type() == parser.NodeString {
		idx := idx.(*parser.StringNode).Value
		return Property(data.(Row), idx)
	}
	if idx.Type() == parser.NodeNumber {
		idx, err := strconv.Atoi(idx.(*parser.NumberNode).Value)
		if err != nil {
			panic("Invalid index")
		}
		return At(data.([]Row), idx)
	}
	return data
}

// @TODO: figure out how to merge processIndex and processHas nicely
func processHas(data interface{}, idx parser.Node) bool {
	if idx.Type() == parser.NodeString {
		idx := idx.(*parser.StringNode).Value
		return HasProperty(data.(Row), idx)
	}
	if idx.Type() == parser.NodeNumber {
		idx, err := strconv.Atoi(idx.(*parser.NumberNode).Value)
		if err != nil {
			panic("Invalid index")
		}
		return HasIndex(data.([]Row), idx)
	}
	return false
}
