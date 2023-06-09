package ch02

import "testing"
import "sort"

func TestSelectionSort(t *testing.T) {
	items := []int{8, 7, 1, 2, 10, 12, 6, 1, 91}
	sorted := SelectionSort(items)
	sort.Ints(items)
	for i, item := range items {
		if item != sorted[i] {
			t.Errorf("expected %d, got %d", item, sorted[i])
		}
	}
}

func TestQuickSort(t *testing.T) {
	items := []int{8, 7, 1, 2, 10, 12, 6, 1, 91}
	sorted := QuickSort(items)
	sort.Ints(items)
	for i, item := range items {
		if item != sorted[i] {
			t.Errorf("expected %d, got %d", item, sorted[i])
		}
	}
}
