package ch03

import "testing"

func TestSetCovering(t *testing.T) {
	states := NewSet[string]([]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})
	stations := map[int]Set[string]{
		0: NewSet[string]([]string{"id", "nv", "ut"}),
		1: NewSet[string]([]string{"wa", "id", "mt"}),
		2: NewSet[string]([]string{"or", "nv", "ca"}),
		3: NewSet[string]([]string{"nv", "ut"}),
		4: NewSet[string]([]string{"ca", "az"}),
	}
	selected := SetCovering[string](states, stations)
	expected := map[int]bool{0: true, 2: true, 4: true, 1: true}
	for key, _ := range expected {
		if _, ok := selected[key]; !ok {
			t.Errorf("expected %d, but did not find it in set", key)
		}
	}
	for key, _ := range selected {
		if _, ok := expected[key]; !ok {
			t.Errorf("did not expect to find %d in set", key)
		}
	}
}
