package guard

import (
	"cmp"
	"fmt"
)

var (
	_ ReferenceGetter[int] = (*SimpleReference[int])(nil)
	_ ReferenceSetter[int] = (*SimpleReference[int])(nil)
)

func NewSimpleReference[T bool | cmp.Ordered]() *SimpleReference[T] {
	return &SimpleReference[T]{
		values: map[string]T{},
	}
}

type SimpleReference[T bool | cmp.Ordered] struct {
	values map[string]T
}

// Set implements ReferenceSetter.
func (s *SimpleReference[T]) Set(key string, value T) {
	s.values[key] = value
}

// Get implements ReferenceGetter.
func (s *SimpleReference[T]) Get(key string) (*T, error) {
	value, exists := s.values[key]
	if !exists {
		return nil, fmt.Errorf("NOT_FOUND")
	}

	return &value, nil
}
