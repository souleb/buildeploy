package dag

import "fmt"

// TopOrder is a topological ordering.
type TopOrder []Vertex

// TopologicalSort implements the kahn algorithm and returns a topological sort of the graph.
// The graph must not have any cycle or it will preturn an error.
func TopologicalSort(graph *Graph) {
	topOrder := make([]string, 0, len(graph.adjacencyMap))
	queue := make([]string, 0, len(graph.adjacencyMap))
	indegreeMap := make(map[string]int)
	var visited int
	processIndegrees(&indegreeMap, graph)

	for vertex := range graph.adjacencyMap {
		v := vertex.(string)
		if indegreeMap[v] == 0 {
			queue = append(queue, v)
		}
	}

	for queue != nil {
		source := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = nil
		}

		topOrder = append(topOrder, source)
		visited++
		for target := range graph.adjacencyMap[source] {
			t := target.(string)
			indegreeMap[t]--
			if indegreeMap[t] == 0 {
				queue = append(queue, t)
			}
		}
	}

	if visited != len(graph.adjacencyMap) {
		fmt.Println("There exists a cycle in the graph")
	}

	fmt.Println(topOrder)

}

func processIndegrees(indegreeMap *map[string]int, graph *Graph) {
	/*
			for each node in Nodes
		    indegree[node] = 0;
			for each edge(src, dest) in Edges
				indegree[dest]++
	*/

	for source, edges := range graph.adjacencyMap {
		if _, ok := (*indegreeMap)[source.(string)]; ok == false {
			(*indegreeMap)[source.(string)] = 0
		}
		for target := range edges {
			(*indegreeMap)[target.(string)]++
		}
	}

}
