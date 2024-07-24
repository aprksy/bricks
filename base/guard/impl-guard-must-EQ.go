package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[bool] = (*SimpleGuardEQ[bool])(nil)

func NewSimpleGuardEQ[T bool | cmp.Ordered](id string, reference ReferenceGetter[T]) SimpleGuardEQ[T] {
	return SimpleGuardEQ[T]{
		SimpleGuardBase: NewSimpleGuardBase[T](id, reference),
	}
}

type SimpleGuardEQ[T bool | cmp.Ordered] struct {
	SimpleGuardBase[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardEQ[T]) Evaluate(value T) bool {
	if s.reference == nil {
		return true
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true
	}

	if !(value == *ref) {
		return false
	}

	return true
}

// Evaluate implements Guard.
func (s *SimpleGuardEQ[T]) EvaluateWithErr(value T) (bool, error) {
	if s.reference == nil {
		return true, nil
	}

	ref, err := s.reference.Get(s.id)

	if err != nil {
		return true, nil
	}

	if !(value == *ref) {
		return false, fmt.Errorf("%s: %s", s.id, ErrRefValueNotEQ)
	}

	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardEQ[T]) GetConstraint() (map[string]T, error) {
	if s.reference == nil {
		return nil, fmt.Errorf(ErrRefNotSet)
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return nil, err
	}
	return map[string]T{s.id: *ref}, nil
}
