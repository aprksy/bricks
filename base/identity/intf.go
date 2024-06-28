package identity

type IDType interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | string
}

type Identity[T IDType] interface {
	Id() T
	TypeName() string
	InstanceInfo() string
}
