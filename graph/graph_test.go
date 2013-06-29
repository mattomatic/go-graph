package graph

import (
	"testing"
)

type String string

func (s String) Key() string {
	return string(s)
}

func TestGraphSimple(t *testing.T) {
	rome := String("Rome")
	london := String("London")
	paris := String("paris")
	g := NewGraph(rome, london, paris)
	g.Connect(rome, london, 100)
	g.Connect(rome, paris, 150)

	enterNode := func(item Item) bool { return true }
	exitNode := enterNode

	g.DepthFirstSearch(enterNode, exitNode)
}
