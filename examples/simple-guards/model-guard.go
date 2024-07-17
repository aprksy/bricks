package main

import (
	"fmt"

	"github.com/aprksy/bricks/base/guard"
)

// define our guard decorator, shall implement same blueprint
type UserGuardDecorator struct {
	ageGuard   guard.Guardable[int]
	emailGuard guard.Guardable[string]
	authGuard  guard.Guardable[bool]
	blGuard    guard.Guardable[bool]
	instance   UserIntf
}

func (u UserGuardDecorator) RequestService() error {
	if _, err := u.ageGuard.AllowWithErr(u.instance.GetAge()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.authGuard.AllowWithErr(u.instance.GetAuth()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.blGuard.AllowWithErr(u.instance.GetBlist()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if _, err := u.emailGuard.AllowWithErr(u.instance.GetEmail()); err != nil {
		fmt.Println(err.Error())
		return err
	}

	// if all passed, invoke the actual code
	return u.instance.RequestService()
}
