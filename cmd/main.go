package main

import (
	"fmt"
	"math/rand"
)




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

}
