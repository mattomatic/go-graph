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
    g := NewGraph(rome, london)
    g.Connect(rome, london, 100)
}