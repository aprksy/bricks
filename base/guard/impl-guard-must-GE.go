package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardGE[int])(nil)

func NewSimpleGuardGE[T cmp.Ordered](reference ReferenceGetter[T]) (*SimpleGuardGE[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardGE[T]{
		reference: reference,
	}, nil
}

type SimpleGuardGE[T cmp.Ordered] struct {
	reference ReferenceGetter[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardGE[T]) Evaluate(actnCtx string, value T) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref > value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueOutOfRange)
	}

	return true, nil
}
