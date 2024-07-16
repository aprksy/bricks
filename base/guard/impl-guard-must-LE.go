package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[int] = (*SimpleGuardLE[int])(nil)

func NewSimpleGuardLE[T cmp.Ordered](reference ReferenceGetter[T]) (*SimpleGuardLE[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardLE[T]{
		reference: reference,
	}, nil
}

type SimpleGuardLE[T cmp.Ordered] struct {
	reference ReferenceGetter[T]
}

// Evaluate implements Guard.
func (s *SimpleGuardLE[T]) Evaluate(actnCtx string, value T) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref < value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueOutOfRange)
	}

	return true, nil
}
