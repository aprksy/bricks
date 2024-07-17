package guard

import (
	"cmp"
)

var _ Guard[bool] = (*SimpleGuardBase[bool])(nil)

func NewSimpleGuardBase[T bool | cmp.Ordered](id string, reference ReferenceGetter[T]) SimpleGuardBase[T] {
	return SimpleGuardBase[T]{
		id:        id,
		reference: reference,
	}
}

type SimpleGuardBase[T bool | cmp.Ordered] struct {
	id        string
	reference ReferenceGetter[T]
}

// Id implements Guard.
func (s *SimpleGuardBase[T]) Id() string {
	return s.id
}

// Reference implements Guard.
func (s *SimpleGuardBase[T]) Reference() ReferenceGetter[T] {
	return s.reference
}

// Evaluate implements Guard.
func (s *SimpleGuardBase[T]) Evaluate(value T) bool {
	return true
}

// EvaluateWithErr implements Guard.
func (s *SimpleGuardBase[T]) EvaluateWithErr(value T) (bool, error) {
	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardBase[T]) GetConstraint() (map[string]T, error) {
	return map[string]T{}, nil
}
