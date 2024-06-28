package observer

import (
	"fmt"
	"sync"
)

var _ Observer[int] = (*SimpleObserver[int])(nil)

func NewSimpleObserver[T comparable](onReceive func(localId string, value T)) *SimpleObserver[T] {
	return &SimpleObserver[T]{
		states:        map[string]T{},
		subscriptions: map[string]Subscription[T]{},
		OnReceive:     onReceive,
	}
}

type SimpleObserver[T comparable] struct {
	mutex         sync.Mutex
	states        map[string]T
	subscriptions map[string]Subscription[T]
	OnReceive     func(localId string, value T)
}

// Receive implements Observer.
func (s *SimpleObserver[T]) Receive(subsId string, value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	localId := s.subscriptions[subsId].LocalId
	s.states[localId] = value
	if s.OnReceive != nil {
		s.OnReceive(localId, value)
	}
}

// Subscribe implements Observer.
func (s *SimpleObserver[T]) Subscribe(subj Subject[T], localId string) (subsId *string, err error) {
	err = fmt.Errorf(ErrSubjectNil)
	if subj == nil {
		return nil, err
	}

	if id, err := subj.Add(s); err != nil {
		return nil, err
	} else {
		subsId = &id
	}
	defer s.Ready(*subsId)

	subs := Subscription[T]{
		LocalId: localId,
		Subject: subj,
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriptions[*subsId] = subs

	return subsId, nil
}

// Unsubscribe implements Observer.
func (s *SimpleObserver[T]) Unsubscribe(subsId string) error {
	subs, exists := s.subscriptions[subsId]
	if !exists {
		return fmt.Errorf(ErrSubscriptionNotFound)
	}

	delete(s.states, subs.LocalId)
	delete(s.subscriptions, subsId)
	return nil
}

// Ready implements Observer.
func (s *SimpleObserver[T]) Ready(subsId string) {
	s.subscriptions[subsId].Subject.Notify()
}

// Extract implements Observer.
func (s *SimpleObserver[T]) Extract(key string) (*T, error) {
	value, exists := s.states[key]
	if !exists {
		return nil, fmt.Errorf(ErrKeyNotFound)
	}

	return &value, nil
}
