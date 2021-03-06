package strtree

import "github.com/yuhangch/gots/geom"

type NodeBase interface {
	Children() *Nodes
}

type Node struct {
	boundaries *Nodes
	boundary   *geom.Envelope
	level      int
}

// Children Get Node's Children
func (n *Node) Children() *Nodes {
	return n.boundaries
}

func NewNode(level int) *Node {
	return &Node{
		level: level,
	}
}

func (n *Node) computeBounds() *geom.Envelope {
	return geom.NewEmptyEnvelope()
}

func (n *Node) Size() int {
	return len(*(n.boundaries))
}

type Nodes []*NodeBase

func (n *Nodes) lastNode() *NodeBase {
	return (*n)[len(*n)-1]
}

type Leaf struct {
	Node
	Items geom.GeometryCollection
}

func (l Leaf) Children() *Nodes {
	return l.boundaries
}
