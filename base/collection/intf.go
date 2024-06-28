package collection

import (
	id "github.com/aprksy/bricks/base/identity"
)

type Collection[K id.IDType, E id.Identity[K]] interface {
	Elements() []E
	Size() int
	HasElement(e E) bool
	Add(e E) error
	Remove(e E) error
	Clear() error
}
