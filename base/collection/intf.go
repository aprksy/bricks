package collection

type Collection[K comparable, E comparable] interface {
	Elements() ([]E, error)
	Size() (int, error)
	HasElement(e E) (bool, error)
	Add(e E) error
	Remove(e E) error
	Clear() error
}

type CollectionWithId[K comparable, E comparable] interface {
	HasElementById(id K) (bool, error)
	Element(id K) (*E, error)
	AddWithId(id K, e E) error
	RemoveById(id K) error
}
