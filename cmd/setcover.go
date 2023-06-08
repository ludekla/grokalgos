package main

import (
	"fmt"

	"grokking/pkg/ch03-hashtb"
)

func main() {

	fmt.Println("Hello Set Covering!")

	// Federal states.
	states := ch03.NewSet[string]([]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})
	// Radio stations with the states they cover.
    stations := map[int]ch03.Set[string]{
        0: ch03.NewSet[string]([]string{"id", "nv", "ut"}),
        1: ch03.NewSet[string]([]string{"wa", "id", "mt"}),
        2: ch03.NewSet[string]([]string{"or", "nv", "ca"}),
        3: ch03.NewSet[string]([]string{"nv", "ut"}),
        4: ch03.NewSet[string]([]string{"ca", "az"}),
    }
	selected := ch03.SetCovering[string](states, stations)
	fmt.Println(selected)

}

	