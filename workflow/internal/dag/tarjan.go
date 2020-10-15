package dag

import "fmt"

// Cycles detect and return cycles in the graph
// It implements the tarjan algorithm
func (g *Graph) Cycles() {
	g2 := g.Copy()

	index := 0

	stack := make([]Vertex, 0, len(g2.adjacencyMap))
	indices := make(map[string]int)
	lowlinks := make(map[string]int)
	onStack := make(map[string]bool)

	for _, v := range g2.hashMap {
		fmt.Println(strongConnect(g2, &v, &index, &stack, indices, lowlinks, onStack))
	}

}

func strongConnect(g *Graph, source *Vertex, index *int, stack *[]Vertex, indices map[string]int, lowlinks map[string]int, onstack map[string]bool) []Vertex {
	indices[hashcode(*source).(string)] = *index
	lowlinks[hashcode(*source).(string)] = *index
	*index++
	*stack = append(*stack, *source)
	onstack[hashcode(*source).(string)] = true
	output := []Vertex{}

	src := hashcode(*source).(string)

	for edge := range g.adjacencyMap[hashcode(*source)] {
		target := edge.(Vertex)
		tar := hashcode(edge).(string)
		if _, ok := indices[tar]; !ok {
			strongConnect(g, &target, index, stack, indices, lowlinks, onstack)
			lowlinks[src] = min(lowlinks[src], lowlinks[tar])
		} else if onstack[hashcode(tar).(string)] {
			lowlinks[src] = min(lowlinks[src], indices[tar])
		}
	}

	if lowlinks[src] == indices[src] {
		var w string
		var v Vertex
		for w != src {
			v, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
			output = append(output, v)
			w = hashcode(v).(string)
		}
	}

	return output

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
