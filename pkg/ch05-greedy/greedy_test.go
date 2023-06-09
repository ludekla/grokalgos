package ch05

import "testing"

func TestDynamicProg(t *testing.T) {
	got := LongCommSubsequence("clues", "blue", false)
	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
	got = LongCommSubstring("clues", "blue", false)
	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
	got = LongCommSubsequence("fish", "fosh", false)
	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
	got = LongCommSubstring("fish", "fosh", false)
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}
