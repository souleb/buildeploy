package dag

// Vertex of the graph
type Vertex interface{}

// Graph represents a directed acyclic graph
type Graph struct {
	AdjacencyMap map[Vertex]LinkedList
}

func NewGraph() *Graph {
	return &Graph{
		AdjacencyMap: map[Vertex]LinkedList{},
	}
}
