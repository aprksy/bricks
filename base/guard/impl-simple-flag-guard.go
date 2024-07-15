package guard

import (
	"fmt"
)

var _ Guard[bool] = (*SimpleFlagGuard)(nil)

func NewSimpleFlagGuard(reference ReferenceGetter[bool]) (*SimpleFlagGuard, error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleFlagGuard{
		reference: reference,
	}, nil
}

type SimpleFlagGuard struct {
	reference ReferenceGetter[bool]
}

// Evaluate implements Guard.
func (s *SimpleFlagGuard) Evaluate(actnCtx string, value bool) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref != value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueNotMatch)
	}

	return true, nil
}
