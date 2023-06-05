package ch00

import "fmt"

func BinarySearch(itemlist []int, item int) (int, error) {
	low := 0
	high := len(itemlist) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := itemlist[mid]
		if guess == item {
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, fmt.Errorf("cannot find %d", item)
}
