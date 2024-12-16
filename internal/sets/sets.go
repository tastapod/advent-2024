package sets

import (
	"iter"
)

type Set[T comparable] struct {
	values map[T]bool // values are keys of underlying set
}

func NewSet[T comparable](values ...T) (s Set[T]) {
	s = Set[T]{
		values: make(map[T]bool, len(values)),
	}
	for _, v := range values {
		s.values[v] = true
	}
	return
}

func (s *Set[T]) Len() int {
	return len(s.values)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Add stores a value in the set
// Returns `true` if a value was added, `false` if it was already there
func (s *Set[T]) Add(value T) {
	s.values[value] = true
}

// All provides an iterator function compatible with [range]
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.values {
			if !yield(v) {
				return
			}
		}
	}
}
