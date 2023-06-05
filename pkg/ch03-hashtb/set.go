package ch03

type Ordered interface {
	~string | ~float64 | ~int
}

type Set[T Ordered] map[T]bool

func NewSet[T Ordered](elems []T) Set[T] {
	set := make(Set[T], len(elems))
	for _, elem := range elems {
		set[elem] = true
	}
	return set
}

func (s Set[T]) Copy() Set[T] {
	newset := make(Set[T], len(s))
	newset.Update(s)
	return newset
}

func (s Set[T]) Update(other Set[T]) {
	for key, _ := range other {
		s[key] = true
	}
}

func (s Set[T]) Minus(other Set[T]) {
	for key, _ := range other {
		delete(s, key)
	}
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	newset := make(Set[T], len(s))
	for key, _ := range other {
		if _, ok := s[key]; ok {
			newset[key] = true
		}
	}
	return newset
}

func (s Set[T]) Clear() {
	for key, _ := range s {
		delete(s, key)
	}
}
