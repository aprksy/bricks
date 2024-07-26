package guard

import (
	"fmt"
	"regexp"
)

var _ Guard[string] = (*SimpleGuardMatch)(nil)

func NewSimpleGuardMatch(id string, reference ReferenceGetter[string]) SimpleGuardMatch {
	return SimpleGuardMatch{
		SimpleGuardBase: NewSimpleGuardBase[string](id, reference),
	}
}

type SimpleGuardMatch struct {
	SimpleGuardBase[string]
}

// Evaluate implements Guard.
func (s *SimpleGuardMatch) Evaluate(value string) bool {
	if s.reference == nil {
		return true
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true
	}

	match, err := regexp.MatchString(*ref, value)
	if err != nil {
		return false
	}

	if !match {
		return false
	}

	return true
}

// EvaluateWithErr implements Guard.
func (s *SimpleGuardMatch) EvaluateWithErr(value string) (bool, error) {
	if s.reference == nil {
		return true, nil
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return true, nil
	}

	match, err := regexp.MatchString(*ref, value)
	if err != nil {
		return false, err
	}

	if !match {
		return false, fmt.Errorf("%s: %s", s.id, ErrValueNotMatch)
	}

	return true, nil
}

// GetConstraint implements Guard.
func (s *SimpleGuardMatch) GetConstraint() (map[string]string, error) {
	if s.reference == nil {
		return nil, fmt.Errorf(ErrRefNotSet)
	}

	ref, err := s.reference.Get(s.id)
	if err != nil {
		return nil, err
	}
	return map[string]string{s.id: *ref}, nil
}
