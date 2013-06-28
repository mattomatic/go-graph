package graph

type Graph struct {
    nodes map[string]*Node
}

func NewGraph(items ...Item) *Graph {
	g := &Graph{make(map[string]*Node)}
	g.AddItems(items...)
	return g
}

func (g *Graph) AddItems(items ...Item) {
	for _, item := range items {
	    g.nodes[item.Key()] = NewNode(item)
	}
}

func (g *Graph) Connect(head Item, tail Item, weight Weight) {
    h := g.Get(head)
    t := g.Get(tail)
    h.connect(t, weight)
}

func (g *Graph) Get(item Item) *Node {
    _, ok := g.nodes[item.Key()]

    if !ok {
        g.nodes[item.Key()] = NewNode(item)
    }
    
    return g.nodes[item.Key()]
}