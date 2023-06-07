package ch00

func BinarySearch[T Ordered](itemlist []T, item T) (int, bool) {
	low := 0
	high := len(itemlist) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := itemlist[mid]
		if guess == item {
			return mid, true
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
