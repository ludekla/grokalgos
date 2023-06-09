package ch01

import (
	"grokking/pkg/ch00-search"
)

func RecMax[T ch00.Ordered](items []T) T {
	if len(items) == 1 {
		return items[0]
	} else if othermax := RecMax(items[1:]); items[0] > othermax {
		return items[0]
	} else {
		return othermax
	}
}
