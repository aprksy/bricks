package hashmap

import (
	id "github.com/aprksy/bricks/base/identity"
)

type Hashmap[K id.IDType, E id.Identity[K]] interface {
	HasElementById(id K) bool
	Element(id K) (*E, error)
	RemoveById(id K) error
}
