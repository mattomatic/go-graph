package graph

type status struct {
	discovered bool
	completed  bool
	parent     *Node
}

type Node struct {
	Item
	status
	edges map[string]*Edge
}

func NewNode(item Item) *Node {
	return &Node{item, status{}, make(map[string]*Edge)}
}

func (n *Node) connect(tail *Node, weight Weight) {
	n.edges[tail.Key()] = NewEdge(n, tail, weight)
}

func (n *Node) reset() {
	n.discovered = false
	n.completed = false
	n.parent = nil
}
