package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleMaxGuard[int])(nil)

func NewSimpleMaxGuard[T cmp.Ordered](reference ReferenceGetter[T]) (*SimpleMaxGuard[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleMaxGuard[T]{
		reference: reference,
	}, nil
}

type SimpleMaxGuard[T cmp.Ordered] struct {
	reference ReferenceGetter[T]
}

// Evaluate implements Guard.
func (s *SimpleMaxGuard[T]) Evaluate(actnCtx string, value T) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref <= value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueOutOfRange)
	}

	return true, nil
}
