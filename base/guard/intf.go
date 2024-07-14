package guard

type Reference[T comparable] interface {
	Set(ket string, value T)
	Get(key string) (T, error)
}

type Guard[T comparable] interface {
	Evaluate(actnCtx string, value T) (bool, error)
}

type Guardable[T comparable] interface {
	Allow(actnCtx string, value T) (bool, error)
}
