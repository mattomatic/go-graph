package graph

// Perform a depth-first search on graph g, calling the enterNode
// function after a node is visited for the first time, and also
// calling the exitNode function after all edges of the node have
// been examined. As long as enterNode and exitNode return true
// the search will continue.
func (g *Graph) DepthFirstSearch(enterNode, exitNode func(Item) bool) {
	g.reset()

	for _, node := range g.nodes {
		if depthFirstSearch(node, enterNode, exitNode) == false {
			return
		}
	}
}

// returns true if the search should continue, false otherwise
func depthFirstSearch(node *Node, enterNode, exitNode func(Item) bool) bool {
	// entering node for the first time
	node.discovered = true

	if enterNode(node) == false {
		return false
	}

	for _, edge := range node.edges {
		next := edge.tail

		if next.discovered == false {
			next.parent = node
			depthFirstSearch(next, enterNode, exitNode)
		}
	}

	// exiting node
	node.completed = true

	if exitNode(node) == false {
		return false
	}

	return true
}
