package hashmap

import (
	"fmt"
	"sync"

	cl "github.com/aprksy/bricks/base/collection"
)

var (
	_ cl.Collection[int, string] = (*SimpleHashmap[int, string])(nil)
	_ Hashmap[int, string]       = (*SimpleHashmap[int, string])(nil)
)

func NewSimpleHashmap[K comparable, E comparable]() *SimpleHashmap[K, E] {
	return &SimpleHashmap[K, E]{
		storage: map[K]E{},
	}
}

type SimpleHashmap[K comparable, E comparable] struct {
	mutex   sync.Mutex
	storage map[K]E
}

// Element implements Hashmap.
func (s *SimpleHashmap[K, E]) Element(id K) (*E, error) {
	e, exists := s.storage[id]
	if !exists {
		return nil, fmt.Errorf(cl.ErrElementNotFound)
	}
	return &e, nil
}

// HasElementById implements Hashmap.
func (s *SimpleHashmap[K, E]) HasElementById(id K) bool {
	_, exists := s.storage[id]
	return exists
}

// RemoveById implements Hashmap.
func (s *SimpleHashmap[K, E]) RemoveById(id K) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.HasElementById(id) {
		return fmt.Errorf(cl.ErrElementNotFound)
	}

	delete(s.storage, id)
	return nil
}

// Add implements collection.Collection.
func (s *SimpleHashmap[K, E]) Add(id K, e E) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.HasElement(e) {
		return fmt.Errorf(cl.ErrElementExists)
	}

	s.storage[id] = e
	return nil
}

// Clear implements collection.Collection.
func (s *SimpleHashmap[K, E]) Clear() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.storage = map[K]E{}
	return nil
}

// Elements implements collection.Collection.
func (s *SimpleHashmap[K, E]) Elements() []E {
	result := []E{}
	for _, e := range s.storage {
		result = append(result, e)
	}
	return result
}

// HasElement implements collection.Collection.
func (s *SimpleHashmap[K, E]) HasElement(e E) bool {
	for _, element := range s.storage {
		if e == element {
			return true
		}
	}
	return false
}

// Remove implements collection.Collection.
func (s *SimpleHashmap[K, E]) Remove(e E) error {
	for id, element := range s.storage {
		if e == element {
			delete(s.storage, id)
			return nil
		}
	}
	return fmt.Errorf(cl.ErrElementNotFound)
}

// Size implements collection.Collection.
func (s *SimpleHashmap[K, E]) Size() int {
	return len(s.storage)
}
