package main

import (
	"fmt"
	"math/rand"

	"grokalgos/pkg/ch02-sort"
)

func main() {

	fmt.Println("Hello Quick Sort!")
	N := 100
	items := make([]int, 10)
	for i, _ := range items {
		items[i] = rand.Intn(N)
	}
	fmt.Println(items)
	items = ch02.QuickSort[int](items)
	fmt.Println(items)

}
