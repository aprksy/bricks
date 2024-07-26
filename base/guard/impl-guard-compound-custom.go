package guard

import "cmp"

var _ CustomCompoundGuard[int] = (*SimpleCustomCompoundGuard[int])(nil)

func NewSimpleCustomCompoundGuard[T bool | cmp.Ordered](id string) SimpleCustomCompoundGuard[T] {
	return SimpleCustomCompoundGuard[T]{
		SimpleCompoundGuard: NewSimpleCompoundGuard[T](id),
	}
}

type SimpleCustomCompoundGuard[T bool | cmp.Ordered] struct {
	SimpleCompoundGuard[T]
	onEval          func(value T) bool
	onEvalWithErr   func(value T) (bool, error)
	onGetConstraint func() (map[string]T, error)
}

// EvaluateWithErr implements CustomCompoundGuard.
// Subtle: this method shadows the method (SimpleCompoundGuard).Evaluate of SimpleCustomCompoundGuard.SimpleCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) EvaluateWithErr(value T) (bool, error) {
	if s.onEvalWithErr != nil {
		return s.onEvalWithErr(value)
	}

	return s.SimpleCompoundGuard.EvaluateWithErr(value)
}

// Evaluate implements CustomCompoundGuard.
// Subtle: this method shadows the method (SimpleCompoundGuard).Evaluate of SimpleCustomCompoundGuard.SimpleCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) Evaluate(value T) bool {
	if s.onEval != nil {
		return s.onEval(value)
	}

	return s.SimpleCompoundGuard.Evaluate(value)
}

// Evaluate implements CustomCompoundGuard.
// Subtle: this method shadows the method (SimpleCompoundGuard).Evaluate of SimpleCustomCompoundGuard.SimpleCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) GetConstraint() (map[string]T, error) {
	if s.onGetConstraint != nil {
		return s.onGetConstraint()
	}

	return s.SimpleCompoundGuard.GetConstraint()
}

// SetOnEvaluate implements CustomCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) SetOnEvaluate(evalFunc func(value T) bool) {
	s.onEval = evalFunc
}

// SetOnEvaluate implements CustomCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) SetOnEvaluateWithErr(evalFunc func(value T) (bool, error)) {
	s.onEvalWithErr = evalFunc
}

// SetOnEvaluate implements CustomCompoundGuard.
func (s *SimpleCustomCompoundGuard[T]) SetOnGetConstraint(getConstraint func() (map[string]T, error)) {
	s.onGetConstraint = getConstraint
}
