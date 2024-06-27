package identity

import "fmt"

var _ Identity[int] = (*SimpleIdentity[int])(nil)

func NewSimpleIdentity[T comparable](id T, typeName string, onInfo func(t string, id T) string) *SimpleIdentity[T] {
	return &SimpleIdentity[T]{
		id:       id,
		typeName: typeName,
		onInfo:   onInfo,
	}
}

type SimpleIdentity[T comparable] struct {
	id       T
	typeName string
	onInfo   func(t string, id T) string
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
	if s.onInfo != nil {
		return s.onInfo(s.typeName, s.id)
	}
	return fmt.Sprintf("%s (%v)", s.typeName, s.id)
}
