package main

import (
	"fmt"
	"math/rand"
)


func BreadthFirstSearch(graph map[string][]string) (string, error) {
	queue := make(NameList, 0, 10)
	checked := make(NameList, 0, 10)
	queue = append(queue, graph["you"]...)
	for !queue.Empty() {
		person := queue.PopLeft()
		if checked.Contains(person) {
			continue
		}
		if person[len(person)-1] == 'm' {
			return person, nil
		} else {
			queue = append(queue, graph[person]...)
			checked = append(checked, person)
		}
	}
	return "", fmt.Errorf("target not found")
}

func Dijkstra(g *Graph) []string {
	for {
		chnode := g.Cheapest()
		if chnode == nil {
			break
		} else if chnode.Edges == nil {
			g.Mark(chnode.Tag)
			continue
		}
		for _, edge := range chnode.Edges {
			node := g.Nodes[edge.Target]
			newCost := chnode.Cost + edge.Weight
			if newCost < node.Cost {
				node.Cost = newCost
				node.Parent = chnode.Tag
			}
		}
		g.Mark(chnode.Tag)
	}
	path := make([]string, 0, 10)
	node := g.Nodes["F"]
	path = append(path, node.Tag)
	for node.Parent != "" {
		path = append(path, node.Parent)
		node = g.Nodes[node.Parent]
	}
	n := len(path) - 1
	reversed := make([]string, n + 1)
	for i, tag := range path {
		reversed[n - i] = tag
	}
	return reversed	
}

func SetCovering(states Set, stations map[int]Set) map[int]bool {
	needed := states.Copy()
	selected := make(map[int]bool, len(stations))
	coveredStates := NewSet(nil)
	for len(needed) > 0 {
		var bestStation int
		var covered Set
		for station, states := range stations {
			covered = needed.Intersect(states)
			if len(covered) > len(coveredStates) {
				bestStation = station
				coveredStates.Clear()
				coveredStates.Update(covered)
			}
		}
		needed.Minus(coveredStates)
		coveredStates.Clear()
		selected[bestStation] = true
	}
	return selected
}

func LongCommSubsequence(leftword, rightword string) int {
	nrows := len(leftword) + 1
	ncols := len(rightword) + 1
	cellgrid := NewCellGrid(nrows, ncols)
	nrows--
	ncols--
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if leftword[r] == rightword[c] {
				cellgrid[r + 1][c + 1] = cellgrid[r][c] + 1
			} else {
				if cellgrid[r + 1][c] > cellgrid[r][c + 1] {
					cellgrid[r + 1][c + 1] = cellgrid[r + 1][c]
				} else {
					cellgrid[r + 1][c + 1] = cellgrid[r][c + 1]
				}
			}	
		} 
	}
	// fmt.Println(cellgrid)
	return cellgrid.Max()
}

func LongCommSubstring(leftword, rightword string) int {
	nrows := len(leftword) + 1
	ncols := len(rightword) + 1
	cellgrid := NewCellGrid(nrows, ncols)
	nrows--
	ncols--
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if leftword[r] == rightword[c] {
				cellgrid[r + 1][c + 1] = cellgrid[r][c] + 1
			}	
		} 
	}
	// fmt.Println(cellgrid)
	return cellgrid.Max()
}

func main() {

	fmt.Printf("1. Binary Search: ")
	N, j := 10_000_000, 10_000_001

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	ix, err := BinarySearch(items, j)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println(ix)
	}

	fmt.Println("2. Selection Sort")
	N = 100
	items = make([]int, 10)
	for i, _ := range items {
		items[i] = rand.Intn(N)
	}
	fmt.Println(items)
	items = SelectionSort(items)
	fmt.Println(items)
	
	fmt.Println("3. Fibonacci Numbers")
	fibos := make([]int, 10)
	for i, _ := range fibos {
		fibos[i] = Fibo(i)
	}
	fmt.Println(fibos)

	fmt.Println("4. Factorial")
	facts := make([]int, 10)
	for i, _ := range facts {
		facts[i] = Fact(i)
	}
	fmt.Println(facts)

	fmt.Println("5. Recursive Maximum Function")
	items = []int{2, 1, 9, 3, 10, 6, 13, 2, 20, 0}
	fmt.Println(RecMax(items))

	fmt.Println("6. Quick Sort")
	N = 100
	items = make([]int, 10)
	for i, _ := range items {
		items[i] = rand.Intn(N)
	}
	fmt.Println(items)
	items = QuickSort(items)
	fmt.Println(items)

	fmt.Println("7. Breadth-first Search")
	graph := map[string][]string{
		"you": []string{"alice", "bob", "claire"}, 
        "claire": []string{"thom", "jonny"},
        "alice": []string{"peggy"},
        "bob": []string{"peggy", "anuj"},
        "peggy": nil, "anuj": nil, "jonny": nil, "thom": nil,
	}
	fmt.Println(BreadthFirstSearch(graph))

	fmt.Println("8. Dijkstra")
	dgraph := NewGraph(map[string]*Node{
        "S": NewNode("S", []Edge{NewEdge("A", 6), NewEdge("B", 2)}),
        "A": NewNode("A", []Edge{NewEdge("F", 1)}),
        "B": NewNode("B", []Edge{NewEdge("A", 3), NewEdge("F", 5)}),
        "F": NewNode("F", nil),
    })
	path := Dijkstra(dgraph)
	fmt.Printf("path: %v, cost: %.2f\n", path, dgraph.Nodes["F"].Cost)

	fmt.Println("9. Set Covering")
	states := NewSet([]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})
    stations := map[int]Set{
        0: NewSet([]string{"id", "nv", "ut"}),
        1: NewSet([]string{"wa", "id", "mt"}),
        2: NewSet([]string{"or", "nv", "ca"}),
        3: NewSet([]string{"nv", "ut"}),
        4: NewSet([]string{"ca", "az"}),
    }
	selected := SetCovering(states, stations)
	fmt.Println(selected)

	fmt.Println("10. Dynamic Programming")
	r := LongCommSubsequence("clues", "blue")
	fmt.Printf("Longest common subsequence of 'clues' and 'blue': %d\n", r)
	r = LongCommSubstring("fish", "fosh")
	fmt.Printf("Longest common substring of 'fish' and 'fosh':    %d\n", r)
	r = LongCommSubsequence("fish", "fosh")
	fmt.Printf("Longest common subsequence of 'fish' and 'fosh':  %d\n", r)

	fmt.Println("11. K-Nearest Neighbour Regression")
	points := []Point{
		Point{5, 1, 0}, Point{1, 1, 0}, Point{3, 1, 1}, 
		Point{4, 0, 0}, Point{4, 0, 1}, Point{2, 0, 0},
	}
	vals := []float64{300, 75, 225, 150, 200, 50}
	nbs := NewNeighboursRegressor(4)
	nbs.Fit(points, vals)
	pt := Point{4, 1, 0}
	fmt.Printf("%v: %0.3f\n", pt, nbs.Predict(pt))



}
