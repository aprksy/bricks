package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardLT[int])(nil)

func NewSimpleGuardLT[T cmp.Ordered](reference ReferenceGetter[T]) (*SimpleGuardLT[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardLT[T]{
		reference: reference,
	}, nil
}

type SimpleGuardLT[T cmp.Ordered] struct {
	reference ReferenceGetter[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardLT[T]) Evaluate(actnCtx string, value T) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref <= value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueOutOfRange)
	}

	return true, nil
}
