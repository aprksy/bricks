package main

import (
	"fmt"

	"github.com/aprksy/bricks/base/guard"
)

type UserIntf interface {
	GetAge() int
	GetEmail() string
	GetAuth() bool
	GetBlist() bool
	RequestService() error
}

type User struct {
	Age           int
	Email         string
	Authenticated bool
	BlackListed   bool
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetAuth() bool {
	return u.Authenticated
}

func (u User) GetBlist() bool {
	return u.BlackListed
}

func (u User) RequestService() error {
	fmt.Printf("user %s requested a service\n", u.Email)
	return nil
}

type UserGuardDecorator struct {
	minGuard   guard.Guardable[int]
	maxGuard   guard.Guardable[int]
	emailGuard guard.Guardable[string]
	authGuard  guard.Guardable[bool]
	blGuard    guard.Guardable[bool]
	instance   UserIntf
}

func (u UserGuardDecorator) RequestService() error {
	if _, err := u.minGuard.Allow("MIN", u.instance.GetAge()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.maxGuard.Allow("MAX", u.instance.GetAge()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.authGuard.Allow("AUTHENTICATED", u.instance.GetAuth()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.blGuard.Allow("BLACKLISTED", u.instance.GetBlist()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.emailGuard.Allow("EMAIL", u.instance.GetEmail()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return u.instance.RequestService()
}

func main() {
	// reference for int values
	intRef := guard.NewSimpleReference[int]()
	// add min: 10, max:100
	intRef.Set("MIN", 20)
	intRef.Set("MAX", 30)

	// reference for bool values
	boolRef := guard.NewSimpleReference[bool]()
	// add authenticated: true, blacklisted: false
	boolRef.Set("AUTHENTICATED", true)
	boolRef.Set("BLACKLISTED", false)

	// reference for string values
	strRef := guard.NewSimpleReference[string]()
	// add email:^[a-z,0-9,\-,\.,_]*@[a-z,0-9,-,\.,_]*.[a-z]$
	strRef.Set("EMAIL", "^[a-z,0-9,\\-,\\.,_]*@[a-z,0-9,-,\\.,_]*.[a-z]$")

	// setup guards
	intMinGuard, _ := guard.NewSimpleMinGuard[int](intRef)
	intMaxGuard, _ := guard.NewSimpleMaxGuard[int](intRef)
	boolMatchGuard, _ := guard.NewSimpleFlagGuard(boolRef)
	strEmailGuard, _ := guard.NewSimpleStrPatternGuard(strRef)

	user := User{
		Age:           201,
		Email:         "aprksy@gmail.co.id",
		Authenticated: true,
		BlackListed:   false,
	}

	authGuard, _ := guard.NewSimpleGuardable(boolMatchGuard)
	blGuard, _ := guard.NewSimpleGuardable(boolMatchGuard)
	minGuard, _ := guard.NewSimpleGuardable(intMinGuard)
	maxGuard, _ := guard.NewSimpleGuardable(intMaxGuard)
	emailGuard, _ := guard.NewSimpleGuardable(strEmailGuard)

	userGuard := UserGuardDecorator{
		authGuard:  authGuard,
		blGuard:    blGuard,
		minGuard:   minGuard,
		maxGuard:   maxGuard,
		emailGuard: emailGuard,
		instance:   user,
	}

	userGuard.RequestService()
}
