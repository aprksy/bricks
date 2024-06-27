package collection

import (
	id "github.com/aprksy/bricks/base/identity"
)

type Collection[K comparable, E id.Identity[K]] interface {
	Elements() []E
	Size() int
	HasElement(e E) bool
	Add(e E) error
	Remove(e E) error
	Clear() error
}
