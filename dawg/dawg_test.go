package dawg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAutomaton(t *testing.T) {
	var auto *Dawg = NewDawg()
	t.Log(auto)
}

func TestConnectTwoNodes(t *testing.T) {
	assert := assert.New(t)
	auto := NewDawg()
	// add dawg node
	root := auto.Root
	nu := auto.NewNode(true)
	root.AddEdge(nu, "a")
	assert.True(root.hasEdge("a"), "Should be able to add edge to new node")
}

func TestTerminatingNode(t *testing.T) {
	assert := assert.New(t)
	auto := NewDawg()
	// add Node
	root := auto.Root
	nu := auto.NewNode(false)

	root.AddEdge(nu, "a")
	assert.True(root.hasEdge("a"), "Should be able to add edge to root node")
	nu.AddEdge(auto.NewNode(true), "b")
	bedge, ok := nu.Edges["b"]
	assert.True(ok, "Edges should have b edge")
	// assert that it's terminating
	assert.True(bedge.Link.Terminating, "B edge should be terminating")
}

/*
Termination:
brackets mean the node can be final

a -> (b) -> (c)
b -> b -> (c)
This means that connection a->b is a terminal, but b->b is not
So we need to store the terminability in the edge
*/
