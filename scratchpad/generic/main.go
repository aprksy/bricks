package main

import "cmp"

type Dependency[T bool | cmp.Ordered] interface {
	Evaluate(value T) bool
}

var _ Dependency[int] = (*Dependency1[int])(nil)

type Dependency1[T cmp.Ordered] struct {
	value T
}

// Get implements Dependency.
func (d *Dependency1[T]) Evaluate(value T) bool {
	panic("unimplemented")
}

var _ Dependency[bool] = (*Dependency2[bool])(nil)

type Dependency2[T comparable] struct {
	value T
}

// Get implements Dependency.
func (d *Dependency2[T]) Evaluate(value T) bool {
	panic("unimplemented")
}

type Main[T bool | cmp.Ordered] struct {
	dependency Dependency[T]
}
