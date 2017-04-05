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
	nu := auto.newNode("b")
	root.AddEdge(nu, "a")
	assert.True(root.hasEdge("a"), true, "Should be able to add edge to new node")
}
