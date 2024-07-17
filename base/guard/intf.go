package guard

import "cmp"

type ReferenceGetter[T bool | cmp.Ordered] interface {
	Get(key string) (*T, error)
}

type ReferenceSetter[T bool | cmp.Ordered] interface {
	Set(ket string, value T)
}

type Guard[T bool | cmp.Ordered] interface {
	Id() string
	EvaluateWithErr(value T) (bool, error)
	Evaluate(value T) bool
	GetConstraint() (map[string]T, error)
}

type CompoundGuard[T bool | cmp.Ordered] interface {
	Guard[T]
	GetGuardWithErr(key string) (Guard[T], error)
	GetGuard(key string) Guard[T]
	SetGuard(guard Guard[T]) CompoundGuard[T]
	ResetGuard(key string) CompoundGuard[T]
	ClearGuard() CompoundGuard[T]
}

type CustomCompoundGuard[T bool | cmp.Ordered] interface {
	CompoundGuard[T]
	HasOnEvaluate() bool
	HasOnEvaluateWithErr() bool
	SetOnEvaluate(evalFunc func(value T) bool)
	SetOnEvaluateWithErr(evalFunc func(value T) (bool, error))
	SetOnGetConstraint(getConstraint func() (map[string]T, error))
}

type Guardable[T bool | cmp.Ordered] interface {
	Allow(value T) bool
	AllowWithErr(value T) (bool, error)
	GetConstraint() (map[string]T, error)
}
