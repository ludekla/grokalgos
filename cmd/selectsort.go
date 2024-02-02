package main

import (
	"fmt"
	"math/rand"

	"grokalgos/pkg/ch02-sort"
)

func main() {

	fmt.Println("2. Selection Sort")
	N := 100
	items := make([]int, 10)
	for i, _ := range items {
		items[i] = rand.Intn(N)
	}
	fmt.Println(items)
	items = ch02.SelectionSort[int](items)
	fmt.Println(items)

}
