package main

import (
	"fmt"

	"grokalgos/pkg/ch04-graph"
)

func main() {

	fmt.Println("Hello Dijkstra!")

	dgraph := ch04.NewGraph(map[string]*ch04.Node{
		"S": ch04.NewNode("S", []ch04.Edge{ch04.NewEdge("A", 6), ch04.NewEdge("B", 2)}),
		"A": ch04.NewNode("A", []ch04.Edge{ch04.NewEdge("F", 1)}),
		"B": ch04.NewNode("B", []ch04.Edge{ch04.NewEdge("A", 3), ch04.NewEdge("F", 5)}),
		"F": ch04.NewNode("F", nil),
	})
	path := ch04.Dijkstra(dgraph)
	fmt.Printf("path: %v, cost: %.2f\n", path, dgraph.Nodes["F"].Cost)

}
