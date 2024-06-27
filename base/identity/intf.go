package identity

type Identity[T comparable] interface {
	Id() T
	TypeName() string
	InstanceInfo() string
}
