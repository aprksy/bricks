package identity

import "fmt"

var _ Identity[int] = (*SimpleIdentity[int])(nil)

func NewSimpleIdentity[T comparable](id T, typeName string) *SimpleIdentity[T] {
	return &SimpleIdentity[T]{
		id:       id,
		typeName: typeName,
	}
}

type SimpleIdentity[T comparable] struct {
	id       T
	typeName string
}

// Id implements Identity.
func (s *SimpleIdentity[T]) Id() T {
	return s.id
}

// TypeName implements Identity.
func (s *SimpleIdentity[T]) TypeName() string {
	return s.typeName
}

// InstanceInfo implements Identity.
func (s *SimpleIdentity[T]) InstanceInfo() string {
	return fmt.Sprintf("%s (%v)", s.typeName, s.id)
}
