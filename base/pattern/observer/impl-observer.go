package observer

import (
	"fmt"
	"sync"
)

var _ Observer[uint, int] = (*SimpleObserver[uint, int])(nil)

func NewSimpleObserverWithSubjectManager[I comparable, T comparable](oid I, onReceive func(key string, value T), sm *SubjectManager[I]) *SimpleObserver[I, T] {
	observer := NewSimpleObserver[I](oid, onReceive)
	observer.subjectMgr = sm
	return observer
}

func NewSimpleObserver[I comparable, T comparable](oid I, onReceive func(key string, value T)) *SimpleObserver[I, T] {
	return &SimpleObserver[I, T]{
		subscriptions: map[string]Subject[I, T]{},
		subsByKey:     map[string]string{},
		dataByKey:     map[string]T{},
		keysBySub:     map[string]string{},
		OnReceive:     onReceive,
	}
}

type SimpleObserver[I comparable, T comparable] struct {
	subjectMgr    *SubjectManager[I]
	mutex         sync.Mutex
	subscriptions map[string]Subject[I, T]
	keysBySub     map[string]string
	subsByKey     map[string]string
	dataByKey     map[string]T
	OnReceive     func(key string, value T)
}

// Extract implements Observer.
func (s *SimpleObserver[I, T]) Extract(key string) (*T, error) {
	chValue, exists := s.dataByKey[key]
	if !exists {
		return nil, fmt.Errorf(ErrKeyNotFound)
	}
	value := chValue
	return &value, nil
}

// Ready implements Observer.
func (s *SimpleObserver[I, T]) Ready(subsId string) {
	s.subscriptions[subsId].Notify()
}

// Receive implements Observer.
func (s *SimpleObserver[I, T]) Receive(subsId string, value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := s.keysBySub[subsId]
	s.dataByKey[key] = value
	if s.OnReceive != nil {
		s.OnReceive(key, value)
	}
}

// Subscribe implements Observer.
func (s *SimpleObserver[I, T]) Subscribe(subject Subject[I, T], key string) (*string, error) {
	err := fmt.Errorf(ErrSubjectNil)
	if subject == nil {
		return nil, err
	}

	subsid, _ := subject.Add(s)

	defer s.Ready(*subsid)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriptions[*subsid] = subject
	s.subsByKey[key] = *subsid
	s.keysBySub[*subsid] = key

	return subsid, nil
}

// SubscribeByKey implements Observer.
func (s *SimpleObserver[I, T]) SubscribeByKey(key string) (*string, Subject[I, T], error) {
	if s.subjectMgr == nil {
		return nil, nil, fmt.Errorf(ErrSubjectMgrNil)
	}

	subsid, subject, err := Subscribe[I, T](s.subjectMgr, key, s)
	if err != nil {
		return nil, nil, err
	}

	defer s.Ready(*subsid)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriptions[*subsid] = subject
	s.subsByKey[key] = *subsid
	s.keysBySub[*subsid] = key

	return subsid, subject, nil
}

// Unsubscribe implements Observer.
func (s *SimpleObserver[I, T]) Unsubscribe(subsId string) error {
	subject, exists := s.subscriptions[subsId]
	if !exists {
		return fmt.Errorf(ErrSubscriptionNotFound)
	}

	subject.Remove(subsId)

	key := s.keysBySub[subsId]
	delete(s.keysBySub, subsId)
	delete(s.subsByKey, key)
	delete(s.dataByKey, key)
	delete(s.subscriptions, subsId)
	return nil
}
