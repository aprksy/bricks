package observer

import (
	"fmt"
	"sync"

	id "github.com/aprksy/bricks/base/identity"
	"github.com/aprksy/bricks/base/utils"
)

var _ Subject[uint, int] = (*SimpleSubject[uint, int])(nil)

func NewSimpleSubject[I id.IDType, T comparable](oid I, key string, value T) *SimpleSubject[I, T] {
	return &SimpleSubject[I, T]{
		Identity:    id.NewSimpleIdentity(oid, "simple-subject", nil),
		key:         key,
		value:       value,
		subscribers: map[string]Observer[I, T]{},
	}
}

type SimpleSubject[I id.IDType, T comparable] struct {
	id.Identity[I]
	mutex       sync.Mutex
	key         string
	value       T
	subscribers map[string]Observer[I, T]
}

// Add implements Subject.
func (s *SimpleSubject[I, T]) Add(obs Observer[I, T]) (*string, error) {
	subsid, _ := utils.RandStr(6)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscribers[subsid] = obs
	return &subsid, nil
}

// Extract implements Subject.
func (s *SimpleSubject[I, T]) Extract() T {
	return s.value
}

// Inject implements Subject.
func (s *SimpleSubject[I, T]) Inject(value T) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.value = value
	s.Notify()
	return nil
}

// Notify implements Subject.
func (s *SimpleSubject[I, T]) Notify() {
	for subsid, obs := range s.subscribers {
		go obs.Receive(subsid, s.value)
	}
}

// Remove implements Subject.
func (s *SimpleSubject[I, T]) Remove(subsId string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.subscribers[subsId]; !exists {
		return fmt.Errorf(ErrSubscriptionNotFound)
	}

	delete(s.subscribers, subsId)
	return nil
}

// Supportedkey implements Subject.
func (s *SimpleSubject[I, T]) Supportedkey() string {
	return s.key
}
