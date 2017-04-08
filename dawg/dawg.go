package dawg

// DawgNode is a struct that will provide common graph node functionality
type DawgNode struct {
	ID          uint32 `json:"id"`
	Terminating bool   `json:"terminating"`
	Edges       Edges  `json:"edges"`
}

// Edges is simply a map that provides a convenient mapping
type Edges map[string]Edge

// Edge is a structure that stores connections between two nodes
type Edge struct {
	Trigger string    `json:"trigger"`
	Link    *DawgNode `json:"link"`
}

// Dawg is a struct that wraps the root node and provides functionality needed for the router
type Dawg struct {
	Root    DawgNode `json:"root"`
	Counter uint32   `json:"counter"`
}

// NewDawg is a constructor function for new dawg instances
func NewDawg() *Dawg {
	return &Dawg{
		Root: DawgNode{
			ID:    0,
			Edges: Edges{},
		},
		Counter: 1,
	}
}

func newEdge(n *DawgNode, trig string) Edge {
	return Edge{
		Link:    n,
		Trigger: trig,
	}
}

/*****************************
           DAWG
*****************************/

// Not exported Dawg functionality
func (auto *Dawg) nextID() uint32 {
	auto.Counter += 1
	return auto.Counter - 1
}

func (auto *Dawg) NewNode(terminate bool) *DawgNode {
	return &DawgNode{
		ID:          auto.nextID(),
		Terminating: terminate,
		Edges:       make(Edges),
	}
}

/*****************************
           DAWG NODE
*****************************/

func (node *DawgNode) AddEdge(other *DawgNode, trigger string) {
	node.Edges[trigger] = newEdge(other, trigger)
}

func (node *DawgNode) hasEdge(trig string) bool {
	_, ok := node.Edges[trig]
	return ok
}
