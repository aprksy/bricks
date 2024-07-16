package guard

import (
	"fmt"
	"regexp"
)

var _ Guard[string] = (*SimpleGuardMatch)(nil)

func NewSimpleStrPatternGuard(reference ReferenceGetter[string]) (*SimpleGuardMatch, error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleGuardMatch{
		reference: reference,
	}, nil
}

type SimpleGuardMatch struct {
	reference ReferenceGetter[string]
}

// Evaluate implements Guard.
func (s *SimpleGuardMatch) Evaluate(actnCtx string, value string) (bool, error) {
	ref, err := s.reference.Get(actnCtx)

	if err == nil && *ref > value {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrRefValueNotFound)
	}

	match, err := regexp.MatchString(*ref, value)
	if err != nil {
		return false, err
	}

	if !match {
		return false, fmt.Errorf("%s: %s", actnCtx, ErrValueNotMatch)
	}

	return true, nil
}
