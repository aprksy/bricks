package observer

import "github.com/aprksy/bricks/base/identity"

/*
	- Subject can only be a dispatcher of a single domain value (slice length, circle radius etc.)
	- key is the name of the value it holds (being dispatched to observers)
*/

type Subject[I, T comparable] interface {
	identity.Identity[I]
	Supports(key string) bool
	Extract() T
	Inject(value T) error
	Add(obs Observer[I, T]) (*I, error)
	Remove(subsId I) error
	Notify()
}

type Observer[I, T comparable] interface {
	identity.Identity[I]
	SubscribeByKey(key string) (*I, Subject[I, T], error)
	Subscribe(subject Subject[I, T], key string) (*I, error)
	Unsubscribe(subsId I) error
	Ready(subsId I)
	Receive(subsId I, value T)
	Extract(key string) (*T, error)
}
