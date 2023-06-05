package ch01

type Ordered interface {
	~string | ~float64 | ~int
}

func RecMax[T Ordered](items []T) T {
	if len(items) == 1 {
		return items[0]
	} else if othermax := RecMax[T](items[1:]); items[0] > othermax {
		return items[0]
	} else {
		return othermax
	}
}
