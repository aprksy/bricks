package observer

type Subscription[T comparable] struct {
	Subject Subject[T]
	LocalId string
}
