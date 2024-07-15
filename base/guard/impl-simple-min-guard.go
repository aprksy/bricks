package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleMinGuard[int])(nil)

func NewSimpleMinGuard[T cmp.Ordered](reference ReferenceGetter[T]) (*SimpleMinGuard[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleMinGuard[T]{
		reference: reference,
	}, nil
}

type SimpleMinGuard[T cmp.Ordered] struct {
	reference ReferenceGetter[T]
}

// Evaluate implements Guard.
func (s *SimpleMinGuard[T]) Evaluate(actnCtx string, value T) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref > value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueOutOfRange)
	}

	return true, nil
}
