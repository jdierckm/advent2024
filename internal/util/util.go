package util

type Direction int

const (
	N Direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

type Set[T comparable] struct {
	elements map[T]bool
}

func (s *Set[T]) Add(v T) {
	s.elements[v] = true
}

func (s *Set[T]) Elements() []T {
	ret := []T{}
	for k, _ := range s.elements {
		ret = append(ret, k)
	}
	return ret
}

func NewSet[T comparable]() *Set[T] {
	s := Set[T]{}
	s.elements = map[T]bool{}
	return &s
}
