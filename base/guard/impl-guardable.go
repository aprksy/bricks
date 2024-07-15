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
func (s *SimpleGuardable[T]) Allow(actnCtx string, value T) (bool, error) {
	return s.guard.Evaluate(actnCtx, value)
}
