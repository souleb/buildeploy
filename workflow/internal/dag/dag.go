package dag

import (
	"fmt"
	"sort"
	"strings"
)

// EdgeSet is a set data structure
type EdgeSet map[Vertex]int

// Graph represents a directed acyclic graph
type Graph struct {
	adjacencyMap map[Vertex]EdgeSet
}

// Add a Vertex to the adjacencyMap.
func (g *Graph) Add(v Vertex) {
	g.init()
	// Add the vertex entry
	hash := hashcode(v)
	if _, ok := g.adjacencyMap[hash]; !ok {
		g.adjacencyMap[hash] = make(EdgeSet)
	}
}

// Remove delete a Vertex from the adjacencyMap.
func (g *Graph) Remove(v Vertex) {
	hash := hashcode(v)
	// delete the vertex entry
	delete(g.adjacencyMap, hash)

	// delete all occurence of the Vertex in the sets.
	for _, set := range g.adjacencyMap {
		delete(set, hash)
	}
}

// AddEdge add an edge to the Graph.
func (g *Graph) AddEdge(source, target Vertex, weight int) {
	g.init()

	// Make sure that every used vertex shows up in our map keys.
	hashSource, hashTarget := hashcode(source), hashcode(target)
	if _, ok := g.adjacencyMap[hashSource]; !ok {
		g.adjacencyMap[hashSource] = make(EdgeSet)
	}

	if _, ok := g.adjacencyMap[hashTarget]; !ok {
		g.adjacencyMap[hashTarget] = make(EdgeSet)
	}

	g.adjacencyMap[hashSource][hashTarget] = weight
}

// RemoveEdge delete an edge from the adjacencyMap.
func (g *Graph) RemoveEdge(source, target Vertex) {
	hashSource, hashTarget := hashcode(source), hashcode(target)
	if set, ok := g.adjacencyMap[hashSource]; ok {
		delete(set, hashTarget)
	}
}

// HasEdge check if an edge exist between to vertices.
func (g *Graph) HasEdge(source, target Vertex) bool {
	hashSource, hashTarget := hashcode(source), hashcode(target)
	if set, ok := g.adjacencyMap[hashSource]; ok {
		if _, ok := set[hashTarget]; ok {
			return true
		}
	}
	return false
}

// HasVertex check if a vertex is in the adjacencyMap.
func (g *Graph) HasVertex(v Vertex) bool {
	hash := hashcode(v)
	if _, ok := g.adjacencyMap[hash]; ok {
		return true
	}
	return false
}

func (g *Graph) verticesEdgesMapAndSlice() ([]string, map[string][]string) {
	names := make([]string, 0, len(g.adjacencyMap))
	mapping := make(map[string][]string, len(g.adjacencyMap))

	// Get the vertices and edges in alphabetical orders by using a string sort.
	// having this deterministic behavior make testing easier.
	for v, targets := range g.adjacencyMap {
		names = append(names, VertexName(v))
		deps := make([]string, 0, len(targets))

		for target, weight := range targets {
			deps = append(deps, fmt.Sprintf(
				"%s (%d)", VertexName(target), weight))
		}
		sort.Strings(deps)

		mapping[VertexName(v)] = deps

	}

	sort.Strings(names)

	return names, mapping
}

// String is a human-friendly representation of the graph
func (g *Graph) String() string {
	var buf strings.Builder
	buf.WriteString("\n")

	names, mapping := g.verticesEdgesMapAndSlice()

	for _, name := range names {
		buf.WriteString(fmt.Sprintf("%s\n", name))
		for _, d := range mapping[name] {
			buf.WriteString(fmt.Sprintf("  %s\n", d))
		}
	}

	return buf.String()
}

func (g *Graph) init() {
	if g.adjacencyMap == nil {
		g.adjacencyMap = make(map[Vertex]EdgeSet)
	}
}
