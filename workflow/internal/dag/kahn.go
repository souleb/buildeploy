package dag

import "fmt"

// TopoOrder is a topological ordering.
type TopoOrder []Vertex

// TopologicalSort implements the kahn algorithm and returns a topological generateTopo of the graph.
// The graph must not have any cycle or it will preturn an error.
// TO DO: return a TopoOrder
func TopologicalSort(graph *Graph) {
	indegreeMap := make(map[string]int)
	processIndegrees(indegreeMap, graph)

	topOrder, visited := generateTopo(graph, indegreeMap)

	if visited != len(graph.adjacencyMap) {
		// TO DO: add a cycle detection
		fmt.Println("There exists a cycle in the graph")
	}

	fmt.Println(topOrder)

}

func initQueue(graph *Graph, indegreeMap map[string]int) []string {
	queue := make([]string, 0, len(graph.adjacencyMap))
	for vertex := range graph.adjacencyMap {
		v := vertex.(string)
		if indegreeMap[v] == 0 {
			queue = append(queue, v)
		}
	}
	return queue
}

func generateTopo(graph *Graph, indegreeMap map[string]int) ([]string, int) {
	topOrder := make([]string, 0, len(graph.adjacencyMap))
	var visited int

	queue := initQueue(graph, indegreeMap)

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
	return topOrder, visited
}

func processIndegrees(indegreeMap map[string]int, graph *Graph) {
	/*
			for each node in Nodes
		    indegree[node] = 0;
			for each edge(src, dest) in Edges
				indegree[dest]++
	*/

	for source, edges := range graph.adjacencyMap {
		if _, ok := indegreeMap[source.(string)]; ok == false {
			indegreeMap[source.(string)] = 0
		}
		for target := range edges {
			indegreeMap[target.(string)]++
		}
	}

}
