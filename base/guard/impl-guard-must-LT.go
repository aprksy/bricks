package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardLT[int])(nil)

func NewSimpleGuardLT[T cmp.Ordered](id string, reference ReferenceGetter[T]) SimpleGuardLT[T] {
	return SimpleGuardLT[T]{
		SimpleGuardBase: NewSimpleGuardBase[T](id, reference),
	}
}

type SimpleGuardLT[T cmp.Ordered] struct {
	SimpleGuardBase[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardLT[T]) Evaluate(value T) bool {
	if s.reference == nil {
		return true
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true
	}

	if *ref <= value {
		return false
	}

	return true
}

// Evaluate implements Guard.
func (s *SimpleGuardLT[T]) EvaluateWithErr(value T) (bool, error) {
	if s.reference == nil {
		return true, nil
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true, nil
	}

	if !(value < *ref) {
		return false, fmt.Errorf("%s: %s", s.id, ErrRefValueNotLT)
	}

	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardLT[T]) GetConstraint() (map[string]T, error) {
	if s.reference == nil {
		return nil, fmt.Errorf(ErrRefNotSet)
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return nil, err
	}
	return map[string]T{s.id: *ref}, nil
}
