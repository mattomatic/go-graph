package graph

type Node struct {
	Item
	edges map[string]*Edge
}

func NewNode(item Item) *Node {
	return &Node{item, make(map[string]*Edge)}
}

func (n *Node) connect (tail *Node, weight Weight) {
    n.edges[tail.Key()] = NewEdge(n, tail, weight)
}