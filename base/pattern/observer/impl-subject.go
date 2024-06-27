package observer

import (
	"fmt"
	"sync"
)

var _ Subject[int] = (*SimpleSubject[int])(nil)

func NewSimpleSubject[T comparable](initData T) *SimpleSubject[T] {
	return &SimpleSubject[T]{
		data:          initData,
		subscriptions: map[string]Observer[T]{},
	}
}

type SimpleSubject[T comparable] struct {
	mtx           sync.Mutex
	subscriptions map[string]Observer[T]
	data          T
}

// Add implements Subject.
func (s *SimpleSubject[T]) Add(obs Observer[T]) (subsId string, err error) {
	subsId = randStr(8)

	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.subscriptions[subsId] = obs
	return subsId, nil
}

// Get implements Subject.
func (s *SimpleSubject[T]) Get() T {
	return s.data
}

// Inject implements Subject.
func (s *SimpleSubject[T]) Inject(value T) {
	s.data = value
	s.Notify()
}

// Notify implements Subject.
func (s *SimpleSubject[T]) Notify() {
	for subsid, obs := range s.subscriptions {
		go obs.Receive(subsid, s.data)
	}
}

// Remove implements Subject.
func (s *SimpleSubject[T]) Remove(subsId string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, exists := s.subscriptions[subsId]; !exists {
		return fmt.Errorf("subscription not exists")
	}

	delete(s.subscriptions, subsId)
	return nil
}
