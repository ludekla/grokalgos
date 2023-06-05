package main

import (
	"testing"
	"sort"
	"strings"
)

func TestBinarySearchFail(t *testing.T) {
	N, j := 100_000, 100_001

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	ix, err := BinarySearch(items, j)
	if err == nil {
		t.Errorf("expected error, got %v\n", ix)
	}
}

func TestBinarySearchLeft(t *testing.T) {
	N, j := 100_000, 8

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	ix, _ := BinarySearch(items, j)
	if ix != 7 {
		t.Errorf("expected 7, got %d\n", ix)
	}
}

func TestBinarySearchRight(t *testing.T) {
	N, j := 100_000, 99992

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	ix, _ := BinarySearch(items, j)
	if ix != 99991 {
		t.Errorf("expected 99991, got %d\n", ix)
	}
}

func TestSelectionSort(t *testing.T) {
	items := []int{8, 7, 1, 2, 10, 12, 6, 1, 91}
	sorted := SelectionSort(items)
	sort.Ints(items)
	for i, item := range items {
		if item != sorted[i] {
			t.Errorf("expected %d, got %d", item, sorted[i])
		}
	}
}

func TestFibo(t *testing.T) {
	fibos := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, fib := range fibos {
		if fib != Fibo(i) {
			t.Errorf("expected Fibonacci number %d, got %d", fib, Fibo(i))
		} 
	}
}

func TestFact(t *testing.T) {
	facts := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	for i, f := range facts {
		if f != Fact(i) {
			t.Errorf("expected %d (factorial of %d), got %d", f, i, Fact(i))
		} 
	}
}

func TestRecMax(t *testing.T) {
	items := []int{1, 2, 40320, 362880,6, 24, 120, 720, 5040, 1}
	expected := 362880
	got := RecMax(items)
	if got != expected {
		t.Errorf("expected %d to be the maximum, got %d", expected, got)
	}
}

func TestQuickSort(t *testing.T) {
	items := []int{8, 7, 1, 2, 10, 12, 6, 1, 91}
	sorted := QuickSort(items)
	sort.Ints(items)
	for i, item := range items {
		if item != sorted[i] {
			t.Errorf("expected %d, got %d", item, sorted[i])
		}
	}
}

func TestBreadthFirst(t *testing.T) {
	graph := map[string][]string{
		"you": []string{"alice", "bob", "claire"}, 
        "claire": []string{"thom", "jonny"},
        "alice": []string{"peggy"},
        "bob": []string{"peggy", "anuj"},
        "peggy": nil, "anuj": nil, "jonny": nil, "thom": nil,
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

func TestSetCovering(t *testing.T) {
	states := NewSet([]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})
    stations := map[int]Set{
        0: NewSet([]string{"id", "nv", "ut"}),
        1: NewSet([]string{"wa", "id", "mt"}),
        2: NewSet([]string{"or", "nv", "ca"}),
        3: NewSet([]string{"nv", "ut"}),
        4: NewSet([]string{"ca", "az"}),
    }
	selected := SetCovering(states, stations)
	expected := map[int]bool{0: true, 2: true, 4: true, 1: true}
	for key, _ := range expected {
		if _, ok := selected[key]; !ok {
			t.Errorf("expected %d, but did not find it in set", key)
		}
	}
	for key, _ := range selected {
		if _, ok := expected[key]; !ok {
			t.Errorf("did not expect to find %d in set", key)
		}
	}
}

func TestDynamicProg(t *testing.T) {
	got := LongCommSubsequence("clues", "blue")
	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
	got = LongCommSubstring("clues", "blue")
	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
	got = LongCommSubsequence("fish", "fosh")
	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
	got = LongCommSubstring("fish", "fosh")
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

func TestNeighboursRegression(t *testing.T) {
	points := []Point{
		Point{5, 1, 0}, Point{1, 1, 0}, Point{3, 1, 1}, 
		Point{4, 0, 0}, Point{4, 0, 1}, Point{2, 0, 0},
	}
	vals := []float64{300, 75, 225, 150, 200, 50}
	nbs := NewNeighboursRegressor(4)
	nbs.Fit(points, vals)
	pt := Point{4, 1, 0}
	expected := 218.75
	got := nbs.Predict(pt)
	if got != expected {
		t.Errorf("expected %f, got %f", expected, got)
	}
}