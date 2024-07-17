package main

import "fmt"

// the actual user used in the program
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
