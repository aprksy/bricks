package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardGT[int])(nil)

func NewSimpleGuardGT[T cmp.Ordered](reference ReferenceGetter[T]) (*SimpleGuardGT[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardGT[T]{
		reference: reference,
	}, nil
}

type SimpleGuardGT[T cmp.Ordered] struct {
	reference ReferenceGetter[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardGT[T]) Evaluate(actnCtx string, value T) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref >= value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueOutOfRange)
	}

	return true, nil
}
