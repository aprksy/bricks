package hashmap

import (
	"fmt"
	"sync"

	cl "github.com/aprksy/bricks/base/collection"
	id "github.com/aprksy/bricks/base/identity"
)

var (
	_ cl.Collection[int, *id.SimpleIdentity[int]] = (*SimpleHashmap[int, *id.SimpleIdentity[int]])(nil)
	_ Hashmap[int, *id.SimpleIdentity[int]]       = (*SimpleHashmap[int, *id.SimpleIdentity[int]])(nil)
)

type SimpleHashmap[K comparable, E id.Identity[K]] struct {
	*id.SimpleIdentity[K]
	mutex   sync.Mutex
	storage map[K]E
}

// Element implements Hashmap.
func (s *SimpleHashmap[K, E]) Element(id K) (*E, error) {
	e, exists := s.storage[id]
	if !exists {
		return nil, fmt.Errorf("element not found")
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
		return fmt.Errorf("element not found")
	}

	delete(s.storage, id)
	return nil
}

// Add implements collection.Collection.
func (s *SimpleHashmap[K, E]) Add(e E) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.HasElement(e) {
		return fmt.Errorf("element exists")
	}

	s.storage[e.Id()] = e
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
	return s.HasElementById(e.Id())
}

// Remove implements collection.Collection.
func (s *SimpleHashmap[K, E]) Remove(e E) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.RemoveById(e.Id())
}

// Size implements collection.Collection.
func (s *SimpleHashmap[K, E]) Size() int {
	return len(s.storage)
}
