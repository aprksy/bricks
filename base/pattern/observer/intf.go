package observer

type Subject[T comparable] interface {
	Get() T
	Inject(value T)
	Add(obs Observer[T]) (subsId string, err error)
	Remove(subsId string) error
	Notify()
}

type Observer[T comparable] interface {
	Subscribe(subj Subject[T], localId string) (subsId *string, err error)
	Unsubscribe(subsId string) error
	Ready(subsId string)
	Receive(subsId string, value T)
	Extract(key string) (*T, error)
}
