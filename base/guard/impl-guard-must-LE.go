package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardLE[int])(nil)

func NewSimpleGuardLE[T cmp.Ordered](id string, reference ReferenceGetter[T]) SimpleGuardLE[T] {
	return SimpleGuardLE[T]{
		SimpleGuardBase: NewSimpleGuardBase[T](id, reference),
	}
}

type SimpleGuardLE[T cmp.Ordered] struct {
	SimpleGuardBase[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardLE[T]) Evaluate(value T) bool {
	if s.reference == nil {
		return true
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true
	}

	if !(value <= *ref) {
		return false
	}

	return true
}

// Evaluate implements Guard.
func (s *SimpleGuardLE[T]) EvaluateWithErr(value T) (bool, error) {
	if s.reference == nil {
		return true, nil
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true, nil
	}

	if !(value <= *ref) {
		return false, fmt.Errorf("%s: %s", s.id, ErrRefValueNotLE)
	}

	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardLE[T]) GetConstraint() (map[string]T, error) {
	if s.reference == nil {
		return nil, fmt.Errorf(ErrRefNotSet)
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return nil, err
	}
	return map[string]T{s.id: *ref}, nil
}
