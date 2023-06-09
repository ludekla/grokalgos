package ch04

import (
	"strings"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	graph := map[string][]string{
		"you":    []string{"alice", "bob", "claire"},
		"claire": []string{"thom", "jonny"},
		"alice":  []string{"peggy"},
		"bob":    []string{"peggy", "anuj"},
		"peggy":  nil, "anuj": nil, "jonny": nil, "thom": nil,
	}
	if name, err := BreadthFirstSearch(graph); err != nil {
		t.Errorf("expected no error, got error '%s'", err)
	} else if name != "thom" {
		t.Errorf("expected 'thom', got %s", name)
	}
}

func TestDijkstra(t *testing.T) {
	graphs := []*Graph{
		NewGraph(map[string]*Node{
			"S": NewNode("S", []Edge{NewEdge("A", 5), NewEdge("C", 2)}),
			"A": NewNode("A", []Edge{NewEdge("B", 4), NewEdge("D", 2)}),
			"C": NewNode("C", []Edge{NewEdge("A", 8), NewEdge("D", 7)}),
			"B": NewNode("B", []Edge{NewEdge("F", 3), NewEdge("D", 6)}),
			"D": NewNode("D", []Edge{NewEdge("F", 1)}),
			"F": NewNode("F", nil),
		}),
		NewGraph(map[string]*Node{
			"S": NewNode("S", []Edge{NewEdge("A", 10)}),
			"A": NewNode("A", []Edge{NewEdge("B", 20)}),
			"B": NewNode("B", []Edge{NewEdge("F", 30)}),
			"C": NewNode("C", []Edge{NewEdge("A", 1), NewEdge("B", 1)}),
			"F": NewNode("F", nil),
		}),
	}
	expected := []string{"S A D F", "S A B F"}
	costs := []float64{8, 60}
	for i, graph := range graphs {
		path := Dijkstra(graph)
		exp := strings.Split(expected[i], " ")
		for i, tag := range path {
			if tag != exp[i] {
				t.Errorf("expected %s, got %s in path", exp[i], tag)
			}
		}
		if graph.Nodes["F"].Cost != costs[i] {
			t.Errorf("expected %.2f, got %.2f as cost", costs[i], graph.Nodes["F"].Cost)
		}
	}
}
