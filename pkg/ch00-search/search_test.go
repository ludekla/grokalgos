package ch00

import (
	"testing"
)

func TestBinarySearchFail(t *testing.T) {
	N, j := 100_000, 100_001

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	if ix, ok := BinarySearch[int](items, j); ok {
		t.Errorf("expected error, got %v\n", ix)
	}
}

func TestBinarySearchLeft(t *testing.T) {
	N, j := 100_000, 8

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	ix, _ := BinarySearch[int](items, j)
	if ix != 7 {
		t.Errorf("expected 7, got %d\n", ix)
	}
}

func TestBinarySearchRight(t *testing.T) {
	N, j := 100_000, 99992

	items := make([]int, N)
	for i, _ := range items {
		items[i] = i + 1
	}

	ix, _ := BinarySearch[int](items, j)
	if ix != 99991 {
		t.Errorf("expected 99991, got %d\n", ix)
	}
}
