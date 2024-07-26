package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[bool] = (*SimpleGuardNE[bool])(nil)

func NewSimpleGuardNE[T bool | cmp.Ordered](id string, reference ReferenceGetter[T]) SimpleGuardNE[T] {
	return SimpleGuardNE[T]{
		SimpleGuardBase: NewSimpleGuardBase[T](id, reference),
	}
}

type SimpleGuardNE[T bool | cmp.Ordered] struct {
	SimpleGuardBase[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardNE[T]) Evaluate(value T) bool {
	if s.reference == nil {
		return true
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true
	}

	if *ref == value {
		return false
	}

	return true
}

// Evaluate implements Guard.
func (s *SimpleGuardNE[T]) EvaluateWithErr(value T) (bool, error) {
	if s.reference == nil {
		return true, nil
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true, nil
	}

	if !(value != *ref) {
		return false, fmt.Errorf("%s: %s", s.id, ErrRefValueNotNE)
	}

	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardNE[T]) GetConstraint() (map[string]T, error) {
	if s.reference == nil {
		return nil, fmt.Errorf(ErrRefNotSet)
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return nil, err
	}
	return map[string]T{s.id: *ref}, nil
}
