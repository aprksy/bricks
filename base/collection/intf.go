package collection

type Collection[K comparable, E comparable] interface {
	Elements() []E
	Size() int
	HasElement(e E) bool
	Remove(e E) error
	Clear() error
}
