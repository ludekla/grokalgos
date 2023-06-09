package ch01

import "testing"

func TestFibo(t *testing.T) {
	fibos := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, fib := range fibos {
		if fib != Fibo(i) {
			t.Errorf("expected Fibonacci number %d, got %d", fib, Fibo(i))
		}
	}
}

func TestFact(t *testing.T) {
	facts := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	for i, f := range facts {
		if f != Fact(i) {
			t.Errorf("expected %d (factorial of %d), got %d", f, i, Fact(i))
		}
	}
}

func TestRecMax(t *testing.T) {
	items := []int{1, 2, 40320, 362880, 6, 24, 120, 720, 5040, 1}
	expected := 362880
	got := RecMax(items)
	if got != expected {
		t.Errorf("expected %d to be the maximum, got %d", expected, got)
	}
}
