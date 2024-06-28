package observer

import (
	"fmt"
	"sync"

	"github.com/aprksy/bricks/base/identity"
)

var _ Observer[int, int] = (*SimpleObserver[int, int])(nil)

func NewSimpleObserver[I, T comparable](oid I, onReceive func(key string, value T)) *SimpleObserver[I, T] {
	id := identity.NewSimpleIdentity[I](oid, "simple-observer", nil)
	return &SimpleObserver[I, T]{
		SimpleIdentity: *id,
		subscriptions:  map[I]Subject[I, T]{},
		subsByKey:      map[string]I{},
		dataByKey:      map[string]T{},
		keysBySub:      map[I]string{},
		OnReceive:      onReceive,
	}
}

type SimpleObserver[I, T comparable] struct {
	identity.SimpleIdentity[I]
	mutex         sync.Mutex
	subscriptions map[I]Subject[I, T]
	keysBySub     map[I]string
	subsByKey     map[string]I
	dataByKey     map[string]T
	OnReceive     func(key string, value T)
}

// Extract implements Observer.
func (s *SimpleObserver[I, T]) Extract(key string) (*T, error) {
	value, exists := s.dataByKey[key]
	if !exists {
		return nil, fmt.Errorf(ErrKeyNotFound)
	}
	return &value, nil
}

// Ready implements Observer.
func (s *SimpleObserver[I, T]) Ready(subsId I) {
	s.subscriptions[subsId].Notify()
}

// Receive implements Observer.
func (s *SimpleObserver[I, T]) Receive(subsId I, value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := s.keysBySub[subsId]
	s.dataByKey[key] = value
	if s.OnReceive != nil {
		s.OnReceive(key, value)
	}
}

// Subscribe implements Observer.
func (s *SimpleObserver[I, T]) Subscribe(subject Subject[I, T], key string) (*I, error) {
	err := fmt.Errorf(ErrSubjectNil)
	if subject == nil {
		return nil, err
	}

	subsid, err := subject.Add(s)
	if err != nil {
		return nil, err
	}

	defer s.Ready(*subsid)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriptions[*subsid] = subject
	s.subsByKey[key] = *subsid
	s.keysBySub[*subsid] = key

	return subsid, nil
}

// SubscribeByKey implements Observer.
func (s *SimpleObserver[I, T]) SubscribeByKey(key string) (*I, Subject[I, T], error) {
	panic("unimplemented")
}

// Unsubscribe implements Observer.
func (s *SimpleObserver[I, T]) Unsubscribe(subsId I) error {
	subject, exists := s.subscriptions[subsId]
	if !exists {
		return fmt.Errorf(ErrSubscriptionNotFound)
	}

	if err := subject.Remove(subsId); err != nil {
		return err
	}

	key := s.keysBySub[subsId]
	delete(s.keysBySub, subsId)
	delete(s.subsByKey, key)
	delete(s.dataByKey, key)
	delete(s.subscriptions, subsId)
	return nil
}
