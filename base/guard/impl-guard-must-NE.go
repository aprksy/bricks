package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[bool] = (*SimpleGuardNE[bool])(nil)

func NewSimpleGuardNE[T bool | cmp.Ordered](reference ReferenceGetter[bool]) (*SimpleGuardNE[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardNE[T]{
		reference: reference,
	}, nil
}

type SimpleGuardNE[T bool | cmp.Ordered] struct {
	reference ReferenceGetter[bool]
}

// Evaluate implements Guard.
func (s *SimpleGuardNE[T]) Evaluate(actnCtx string, value bool) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref == value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueNotMatch)
	}

	return true, nil
}
