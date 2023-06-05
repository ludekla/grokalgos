package ch02

type Ordered interface {
	~string | ~float64 | ~int
}

func findSmallest[T Ordered](items []T) int {
	smallest, ix := items[0], 0
	for i, item := range items {
		if item < smallest {
			smallest = item
			ix = i
		}
	}
	return ix
}

func popItem[T Ordered](items []T, ix int) (T, []T) {
	item := items[ix]
	copy(items[ix:], items[ix+1:])
	items = items[:len(items)-1]
	return item, items
}

func SelectionSort[T Ordered](items []T) []T {
	newList := make([]T, 0, len(items))
	citems := make([]T, len(items))
	copy(citems, items)
	var item T
	for len(citems) > 0 {
		ix := findSmallest(citems)
		item, citems = popItem(citems, ix)
		newList = append(newList, item)
	}
	return newList
}
