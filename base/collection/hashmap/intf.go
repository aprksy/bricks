package hashmap

type Hashmap[K comparable, E comparable] interface {
	HasElementById(id K) bool
	Element(id K) (*E, error)
	Add(id K, e E) error
	RemoveById(id K) error
}
