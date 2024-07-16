package guard

import "cmp"

type ReferenceGetter[T bool | cmp.Ordered] interface {
	Get(key string) (*T, error)
}

type ReferenceSetter[T bool | cmp.Ordered] interface {
	Set(ket string, value T)
}

type Guard[T bool | cmp.Ordered] interface {
	Evaluate(actnCtx string, value T) (bool, error)
}

type CompoundGuard[T bool | cmp.Ordered] interface {
	Guard[T]
	SetGuard(key string, guard Guard[T]) CompoundGuard[T]
	ResetGuard(key string) CompoundGuard[T]
	ClearGuard() CompoundGuard[T]
}

type CustomCompoundGuard[T bool | cmp.Ordered] interface {
	CompoundGuard[T]
	SetOnEvaluate(evalFunc func(actnCtx string, value T) (bool, error))
}

type Guardable[T bool | cmp.Ordered] interface {
	Allow(actnCtx string, value T) (bool, error)
}
