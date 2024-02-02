package main

import (
	"flag"
	"fmt"

	"grokalgos/pkg/ch05-greedy"
)

var verbose = flag.Bool("v", false, "verbose, ie show grid")

func main() {
	flag.Parse()

	fmt.Println("Hello Dynamic Programming!")

	r := ch05.LongCommSubsequence("clues", "blue", *verbose)
	fmt.Printf("Longest common subsequence of 'clues' and 'blue': %d\n", r)

	r = ch05.LongCommSubstring("fish", "fosh", *verbose)
	fmt.Printf("Longest common substring of 'fish' and 'fosh':    %d\n", r)

	r = ch05.LongCommSubsequence("fish", "fosh", *verbose)
	fmt.Printf("Longest common subsequence of 'fish' and 'fosh':  %d\n", r)

}
