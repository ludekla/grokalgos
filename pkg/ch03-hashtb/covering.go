package ch03

func SetCovering[T Ordered](states Set[T], stations map[int]Set[T]) map[int]bool {
	needed := states.Copy()
	selected := make(map[int]bool, len(stations))
	coveredStates := NewSet[T](nil)
	for len(needed) > 0 {
		var bestStation int
		var covered Set[T]
		for station, states := range stations {
			covered = needed.Intersect(states)
			if len(covered) > len(coveredStates) {
				bestStation = station
				coveredStates.Clear()
				coveredStates.Update(covered)
			}
		}
		needed.Minus(coveredStates)
		coveredStates.Clear()
		selected[bestStation] = true
	}
	return selected
}