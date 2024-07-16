package guard

import (
	"cmp"
	"fmt"
)

var _ CompoundGuard[int] = (*SimpleCompoundGuard[int])(nil)

func NewSimpleCompoundGuard[T bool | cmp.Ordered]() *SimpleCompoundGuard[T] {
	return &SimpleCompoundGuard[T]{
		guards: map[string]Guard[T]{},
	}
}

type SimpleCompoundGuard[T bool | cmp.Ordered] struct {
	guards map[string]Guard[T]
}

// ClearGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) ClearGuard() CompoundGuard[T] {
	s.guards = map[string]Guard[T]{}
	return s
}

// Evaluate implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) Evaluate(actnCtx string, value T) (bool, error) {
	fmt.Println("need outer implementation to shadow this method")
	return true, nil
}

// ResetGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) ResetGuard(key string) CompoundGuard[T] {
	delete(s.guards, key)
	return s
}

// SetGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) SetGuard(key string, guard Guard[T]) CompoundGuard[T] {
	s.guards[key] = guard
	return s
}
