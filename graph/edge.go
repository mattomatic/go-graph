package graph

type Weight float64

type Edge struct {
	head *Node
	tail *Node
	weight Weight
}

func NewEdge(head, tail *Node, weight Weight) *Edge {
    return &Edge{head, tail, weight}
}