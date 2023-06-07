package main

import (
	"fmt"

	"grokking/pkg/ch01-recursion"
)

func main() {

	fmt.Println("Hello Recursion!")

	fmt.Println("Fibonacci Numbers")
	fibos := make([]int, 10)
	for i, _ := range fibos {
		fibos[i] = ch01.Fibo(i)
	}
	fmt.Println(fibos)

	fmt.Println("Factorial")
	facts := make([]int, 10)
	for i, _ := range facts {
		facts[i] = ch01.Fact(i)
	}
	fmt.Println(facts)

	fmt.Println("Recursive Maximum Function")
	items := []int{2, 1, 9, 3, 10, 6, 13, 2, 20, 0}
	fmt.Println(ch01.RecMax[int](items))

}