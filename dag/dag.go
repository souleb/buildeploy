package dag

import "fmt"

// Vertex of the graph
type Vertex interface{}

// Graph represents a directed acyclic graph
type Graph struct {
	adjacencyMap map[Vertex]LinkedList
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyMap: map[Vertex]LinkedList{},
	}
}

func (g *Graph) AddEdge(source Vertex, target Vertex) {

	// Make sure that every used vertex shows up in our map keys.
	if _, ok := g.adjacencyMap[source]; !ok {
		g.adjacencyMap[source] = LinkedList{}
	}

	if _, ok := g.adjacencyMap[target]; !ok {
		g.adjacencyMap[target] = LinkedList{}
	}

	g.addEdgeUniquify(source, target)
}

func (g *Graph) addEdgeUniquify(source Vertex, target Vertex) {
	tempList := g.adjacencyMap[source]
	if (tempList != LinkedList{}) {
		tempList.Remove(target)
	}

	tempList.Append(target)
	g.adjacencyMap[source] = tempList

}

func (g *Graph) PrintEdges() {
	for _, list := range g.adjacencyMap {
		fmt.Println("The vertex has an edge towards: ")
		if (list != LinkedList{}) {
			fmt.Println(list)
		}
	}
}

func (g *Graph) HasEdge(source Vertex, target Vertex) bool {
	if list, ok := g.adjacencyMap[source]; ok {
		if (list != LinkedList{}) {
			if err := list.Get(target); err != nil {
				return false
			}
		}
	}

	return true
}
