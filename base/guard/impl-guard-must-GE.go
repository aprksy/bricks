package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardGE[int])(nil)

func NewSimpleGuardGE[T cmp.Ordered](id string, reference ReferenceGetter[T]) SimpleGuardGE[T] {
	return SimpleGuardGE[T]{
		SimpleGuardBase: NewSimpleGuardBase[T](id, reference),
	}
}

type SimpleGuardGE[T cmp.Ordered] struct {
	SimpleGuardBase[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardGE[T]) Evaluate(value T) bool {
	if s.reference == nil {
		return true
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true
	}

	if !(value >= *ref) {
		return false
	}

	return true
}

// Evaluate implements Guard.
func (s *SimpleGuardGE[T]) EvaluateWithErr(value T) (bool, error) {
	if s.reference == nil {
		return true, nil
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true, nil
	}

	if !(value >= *ref) {
		return false, fmt.Errorf("%s: %s", s.id, ErrRefValueNotGE)
	}

	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardGE[T]) GetConstraint() (map[string]T, error) {
	if s.reference == nil {
		return nil, fmt.Errorf(ErrRefNotSet)
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return nil, err
	}
	return map[string]T{s.id: *ref}, nil
}
