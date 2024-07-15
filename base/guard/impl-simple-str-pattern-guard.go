package guard

import (
	"fmt"
	"regexp"
)

var _ Guard[string] = (*SimpleStrPatternGuard)(nil)

func NewSimpleStrPatternGuard(reference ReferenceGetter[string]) (*SimpleStrPatternGuard, error) {
	if reference == nil {
		return nil, fmt.Errorf(ErrRefProviderNil)
	}

	return &SimpleStrPatternGuard{
		reference: reference,
	}, nil
}

type SimpleStrPatternGuard struct {
	reference ReferenceGetter[string]
}

// Evaluate implements Guard.
func (s *SimpleStrPatternGuard) Evaluate(actnCtx string, value string) (bool, error) {
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
