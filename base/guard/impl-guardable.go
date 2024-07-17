package guard

import (
	"cmp"
	"fmt"
)

var _ Guardable[int] = (*SimpleGuardable[int])(nil)

func NewSimpleGuardable[T bool | cmp.Ordered](guard Guard[T]) (*SimpleGuardable[T], error) {
	if guard == nil {
		return nil, fmt.Errorf("guard is nil")
	}
	return &SimpleGuardable[T]{
		guard: guard,
	}, nil
}

type SimpleGuardable[T bool | cmp.Ordered] struct {
	guard Guard[T]
}

// Allow implements Guardable.
func (s *SimpleGuardable[T]) Allow(value T) bool {
	return s.guard.Evaluate(value)
}

// Allow implements Guardable.
func (s *SimpleGuardable[T]) AllowWithErr(value T) (bool, error) {
	return s.guard.EvaluateWithErr(value)
}

// GetConstraint implements Guardable.
func (s *SimpleGuardable[T]) GetConstraint() (map[string]T, error) {
	return s.guard.GetConstraint()
}
