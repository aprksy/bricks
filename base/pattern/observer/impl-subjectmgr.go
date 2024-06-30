package observer

import (
	"fmt"

	id "github.com/aprksy/bricks/base/identity"
)

func NewSubjectManager[I id.IDType]() *SubjectManager[I] {
	return &SubjectManager[I]{
		subjects: map[string]any{},
	}
}

type SubjectManager[I id.IDType] struct {
	subjects map[string]any
}

func AddSubjects[I id.IDType, T comparable](subjmgr *SubjectManager[I], subjs ...Subject[I, T]) error {
	s := subjmgr
	for _, subject := range subjs {
		_, exists := s.subjects[subject.Supportedkey()]
		if exists {
			return fmt.Errorf(ErrKeyExists)
		}
	}

	for _, subject := range subjs {
		s.subjects[subject.Supportedkey()] = subject
	}

	return nil
}

func Inject[I id.IDType, T comparable](subjmgr *SubjectManager[I], key string, value T) error {
	s := subjmgr
	subject, exists := s.subjects[key]
	if !exists {
		return fmt.Errorf(ErrKeyNotFound)
	}

	return subject.(Subject[I, T]).Inject(value)
}

func Subscribe[I id.IDType, T comparable](subjmgr *SubjectManager[I], key string, obs Observer[I, T]) (*string, Subject[I, T], error) {
	s := subjmgr
	subjectRaw, exists := s.subjects[key]
	if !exists {
		return nil, nil, fmt.Errorf(ErrKeyNotFound)
	}

	subject := subjectRaw.(Subject[I, T])
	subsid, err := subject.Add(obs)
	return subsid, subject, err
}
