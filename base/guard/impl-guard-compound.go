package guard

import (
	"cmp"
	"fmt"
)

var _ CompoundGuard[int] = (*SimpleCompoundGuard[int])(nil)

func NewSimpleCompoundGuard[T bool | cmp.Ordered](id string) SimpleCompoundGuard[T] {
	return SimpleCompoundGuard[T]{
		SimpleGuardBase: NewSimpleGuardBase[T](id, nil),
		guards:          map[string]Guard[T]{},
	}
}

type SimpleCompoundGuard[T bool | cmp.Ordered] struct {
	SimpleGuardBase[T]
	guards map[string]Guard[T]
}

// ClearGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) ClearGuard() CompoundGuard[T] {
	s.guards = map[string]Guard[T]{}
	return s
}

// EvaluateWithErr implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) EvaluateWithErr(value T) (bool, error) {
	return true, nil
}

// Evaluate implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) Evaluate(value T) bool {
	return true
}

// ResetGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) ResetGuard(key string) CompoundGuard[T] {
	delete(s.guards, key)
	return s
}

// SetGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) SetGuard(guard Guard[T]) CompoundGuard[T] {
	s.guards[guard.Id()] = guard
	return s
}

// GetGuard implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) GetGuardWithErr(key string) (Guard[T], error) {
	result, exists := s.guards[key]
	if !exists {
		return nil, fmt.Errorf("not found")
	}
	return result, nil
}

// GetGuardWithErr implements CompoundGuard.
func (s *SimpleCompoundGuard[T]) GetGuard(key string) Guard[T] {
	result, exists := s.guards[key]
	if !exists {
		return nil
	}
	return result
}

// GetConstraint implements Guard.
func (s *SimpleCompoundGuard[T]) GetConstraint() (map[string]T, error) {
	return nil, fmt.Errorf("need outer implementation to shadow this method")
}
