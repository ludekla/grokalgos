package main

import (
	"fmt"

	"grokking/pkg/ch04-graph"
)


func main() {
	fmt.Println("7. Breadth-first Search")
	graph := map[string][]string{
		"you": []string{"alice", "bob", "claire"}, 
        "claire": []string{"thom", "jonny"},
        "alice": []string{"peggy"},
        "bob": []string{"peggy", "anuj"},
        "peggy": nil, "anuj": nil, "jonny": nil, "thom": nil,
	}
	fmt.Println(ch04.BreadthFirstSearch(graph))
}