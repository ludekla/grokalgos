package ch04

import (
	"fmt"

	"grokking/pkg/ch00-search"
)

func BreadthFirstSearch(graph map[string][]string) (string, error) {
	queue := make(ch00.List[string], 0, 10)
	checked := make(ch00.List[string], 0, 10)
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
