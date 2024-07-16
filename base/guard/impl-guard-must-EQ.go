package guard

import (
	"cmp"
	"fmt"
)

var _ Guard[bool] = (*SimpleGuardEQ[bool])(nil)

func NewSimpleGuardEQ[T bool | cmp.Ordered](reference ReferenceGetter[bool]) (*SimpleGuardEQ[T], error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardEQ[T]{
		reference: reference,
	}, nil
}

type SimpleGuardEQ[T bool | cmp.Ordered] struct {
	reference ReferenceGetter[bool]
}

// Evaluate implements Guard.
func (s *SimpleGuardEQ[T]) Evaluate(actnCtx string, value bool) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref != value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueNotMatch)
	}

	return true, nil
}
