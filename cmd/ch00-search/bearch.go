package main

import (
	"fmt"

	"grokalgos/pkg/ch00-intro"
)

func main() {
	fmt.Printf("0. Binary Search: ")
	N, j := 10_000_000, 10_000_001

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	if ix, ok := ch00.BinarySearch[int](items, j); ok {
		fmt.Printf("found element at index %d\n", ix)
	} else {
		fmt.Printf("cannot find %d\n", j)
	}
}
