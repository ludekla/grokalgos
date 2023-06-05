package ch02

func QuickSort[T Ordered](items []T) []T {
	citems := make([]T, 0, len(items))
	if len(items) < 2 {
		copy(citems, items)
		return citems
	}
	pivot := items[0]
	for _, item := range items[1:] {
		if item <= pivot {
			citems = append(citems, item)
		}
	}
	less := QuickSort(citems)
	copy(citems, less)
	citems = append(citems, pivot)
	lastix := len(citems)
	for _, item := range items[1:] {
		if item > pivot {
			citems = append(citems, item)
		}
	}
	greater := QuickSort(citems[lastix:])
	copy(citems[lastix:], greater)
	return citems
}
