package stack

import (
	"fmt"
)

func isValid(s string) bool {
	bytes := []byte(s)
	fmt.Println(string(bytes[0]))

	t := [...]int{1,2,3}
	fmt.Println(t)



	return true
}
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	fmt.Println(graph[from])
	return graph[from][to]
}