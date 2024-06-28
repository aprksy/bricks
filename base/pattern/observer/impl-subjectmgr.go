package observer

import "fmt"

var _ SubjectManager[int] = (*SimpleSubjectManager[int])(nil)

func NewSubjectManager[I comparable]() *SimpleSubjectManager[I] {
	return &SimpleSubjectManager[I]{
		subjects: map[string]Subject[I, any]{},
	}
}

type SimpleSubjectManager[I comparable] struct {
	subjects map[string]Subject[I, any]
}

// AddSubjects implements SubjectManager.
func (s *SimpleSubjectManager[I]) AddSubjects(subjects ...Subject[I, any]) error {
	for _, subject := range subjects {
		_, exists := s.subjects[subject.Supportedkey()]
		if !exists {
			return fmt.Errorf(ErrKeyExists)
		}
	}

	for _, subject := range subjects {
		s.subjects[subject.Supportedkey()] = subject
	}

	return nil
}

// Inject implements SubjectManager.
func (s *SimpleSubjectManager[I]) Inject(key string, value any) error {
	subject, exists := s.subjects[key]
	if !exists {
		return fmt.Errorf(ErrKeyNotFound)
	}

	return subject.Inject(value)
}

// Subscribe implements SubjectManager.
func (s *SimpleSubjectManager[I]) Subscribe(key string, observer Observer[I, any]) (*I, Subject[I, any], error) {
	subject, exists := s.subjects[key]
	if !exists {
		return nil, nil, fmt.Errorf(ErrKeyNotFound)
	}

	subsid, err := subject.Add(observer)
	return subsid, subject, err
}
