package guard

import "cmp"

var _ CustomCompoundGuard[int] = (*SimpleCustomCompoundGuard[int])(nil)

type SimpleCustomCompoundGuard[T bool | cmp.Ordered] struct {
	SimpleCompoundGuard[T]
	onEval func(actnCtx string, value T) (bool, error)
}

// Evaluate implements CustomCompoundGuard.
// Subtle: this method shadows the method (SimpleCompoundGuard).Evaluate of SimpleCustomCompoundGuard.SimpleCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) Evaluate(actnCtx string, value T) (bool, error) {
	if s.onEval != nil {
		return s.onEval(actnCtx, value)
	}

	return s.SimpleCompoundGuard.Evaluate(actnCtx, value)
}

// SetOnEvaluate implements CustomCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) SetOnEvaluate(evalFunc func(actnCtx string, value T) (bool, error)) {
	s.onEval = evalFunc
}
