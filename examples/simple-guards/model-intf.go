package main

// the blueprint of the User model
type UserIntf interface {
	GetAge() int
	GetEmail() string
	GetAuth() bool
	GetBlist() bool
	RequestService() error
}
