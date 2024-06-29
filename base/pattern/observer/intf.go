package observer

import id "github.com/aprksy/bricks/base/identity"

/*
	- Subject can only be a dispatcher of a single domain value (slice length, circle radius etc.)
	- key is the name of the value it holds (being dispatched to observers)
*/

type Subject[I id.IDType, T comparable] interface {
	id.Identity[I]
	Supportedkey() string
	Extract() T
	Inject(value T) error
	Add(obs Observer[I, T]) (*string, error)
	Remove(subsId string) error
	Notify()
}

type Observer[I id.IDType, T comparable] interface {
	id.Identity[I]
	SubscribeByKey(key string) (*string, Subject[I, T], error)
	Subscribe(subject Subject[I, T], key string) (*string, error)
	Unsubscribe(subsId string) error
	Ready(subsId string)
	Receive(subsId string, value T)
	Extract(key string) (*T, error)
}
