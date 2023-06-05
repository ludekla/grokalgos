package ch00

type Ordered interface {
	~string | ~float64 | ~int
}

type List[T Ordered] []T

func (l List[T]) Empty() bool {
	return len(l) == 0
}

func (l List[T]) Contains(obj T) bool {
	for _, elem := range l {
		if elem == obj {
			return true
		}
	}
	return false
}

func (l *List[T]) PopLeft() T {
	purgee := (*l)[0]
	*l = (*l)[1:]
	return purgee
}
