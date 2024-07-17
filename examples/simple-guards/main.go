package main

import (
	"fmt"

	"github.com/aprksy/bricks/base/guard"
)

func main() {
	// -------------------------------------------------------------------
	// reference shall be populated from configuration
	// reference for int values
	intRef := guard.NewSimpleReference[int]()
	// add min: 20, max:30, value can only between these values
	intRef.Set("AGE.MIN", 20)
	intRef.Set("AGE.MAX", 30)

	// reference for bool values
	boolRef := guard.NewSimpleReference[bool]()
	// add authenticated: true, blacklisted: false;
	// must authenticated, must not blacklisted
	boolRef.Set("AUTHENTICATED", true)
	boolRef.Set("BLACKLISTED", false)

	// reference for string values
	strRef := guard.NewSimpleReference[string]()
	// add email:^[a-z,0-9,\-,\.,_]*@[a-z,0-9,-,\.,_]*.[a-z]$
	strRef.Set("EMAIL", "^[a-z,0-9,\\-,\\.,_]*@[a-z,0-9,-,\\.,_]*.[a-z]$")
	// -------------------------------------------------------------------

	// setup guards
	ageRangeGuard := guard.NewSimpleCustomCompoundGuard[int]("AGE")
	ageMinGuard := guard.NewSimpleGuardGE[int]("AGE.MIN", intRef)
	ageMaxGuard := guard.NewSimpleGuardLE[int]("AGE.MAX", intRef)
	ageRangeGuard.
		SetGuard(&ageMinGuard).
		SetGuard(&ageMaxGuard)
	ageRangeGuard.SetOnEvaluateWithErr(func(value int) (bool, error) {
		minResult, _ := ageRangeGuard.GetGuard("AGE.MIN").EvaluateWithErr(value)
		maxResult, _ := ageRangeGuard.GetGuard("AGE.MAX").EvaluateWithErr(value)

		var err error
		if !(minResult && maxResult) {
			err = fmt.Errorf("%s: value out of range", ageRangeGuard.Id())
		}
		return minResult && maxResult, err
	})
	ageRangeGuard.SetOnGetConstraint(func() (map[string]int, error) {
		min, _ := ageRangeGuard.GetGuard("AGE.MIN").GetConstraint()
		max, _ := ageRangeGuard.GetGuard("AGE.MAX").GetConstraint()
		result := map[string]int{"AGE.MIN": min["AGE.MIN"], "AGE.MAX": max["AGE.MAX"]}
		return result, nil
	})

	authGuard := guard.NewSimpleGuardEQ[bool]("AUTHENTICATED", boolRef)
	blGuard := guard.NewSimpleGuardEQ[bool]("BLACKLISTED", boolRef)
	strEmailGuard := guard.NewSimpleGuardMatch("EMAIL", strRef)

	user := User{
		Age:           23,
		Email:         "aprksy@example.com",
		Authenticated: true,
		BlackListed:   false,
	}

	authGuardable, _ := guard.NewSimpleGuardable(&authGuard)
	blGuardable, _ := guard.NewSimpleGuardable(&blGuard)
	ageGuardable, _ := guard.NewSimpleGuardable(&ageRangeGuard)
	emailGuardable, _ := guard.NewSimpleGuardable(&strEmailGuard)

	userGuard := UserGuardDecorator{
		authGuard:  authGuardable,
		blGuard:    blGuardable,
		ageGuard:   ageGuardable,
		emailGuard: emailGuardable,
		instance:   user,
	}

	// request service
	userGuard.RequestService()

	// change the min value
	// intRef.Set("AGE.MIN", 27)
	// boolRef.Set("BLACKLISTED", true)
	// boolRef.Set("AUTHENTICATED", false)

	// request service again
	userGuard.RequestService()
	// constraint, _ := ageRangeGuard.GetConstraint()
	// fmt.Println(constraint)
}
