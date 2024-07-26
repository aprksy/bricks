package guard_test

import (
	"testing"

	g "github.com/aprksy/bricks/base/guard"
	a "github.com/stretchr/testify/assert"
)

func Test_NewSimpleCustomCompoundGuard(t *testing.T) {
	instance := g.NewSimpleCustomCompoundGuard[int]("key1")
	a.NotNil(t, instance, "instance should not be nil")
	a.Equalf(t, "key1", instance.Id(), "instance.Id() should equal 'keys1'")
}

func Test_CustomEvaluate(t *testing.T) {
	// create reference
	intRef := g.NewSimpleReference[int]()
	intRef.Set("SOME-CONTEXT.VALUE", 100)

	// create equality guard
	eqGuard := g.NewSimpleGuardEQ("SOME-CONTEXT.VALUE", intRef)

	// create custom guard from equality guard
	instance := g.NewSimpleCustomCompoundGuard[int]("SOME-CONTEXT")
	instance.SetGuard(&eqGuard)

	result := instance.Evaluate(0)
	a.True(t, result, "result should be true")

	result = instance.Evaluate(100)
	a.True(t, result, "result should be true")

	result = instance.Evaluate(101)
	a.True(t, result, "result should be true")

	instance.SetOnEvaluate(func(value int) bool {
		return instance.GetGuard("SOME-CONTEXT.VALUE").Evaluate(value)
	})

	instance.SetOnGetConstraint(func() (map[string]int, error) {
		return instance.GetGuard("SOME-CONTEXT").GetConstraint()
	})

	result = instance.Evaluate(0)
	a.False(t, result, "result should be false")

	result = instance.Evaluate(100)
	a.True(t, result, "result should be true")

	result = instance.Evaluate(101)
	a.False(t, result, "result should be false")
}

func Test_CustomEvaluateWithErr(t *testing.T) {
	// create reference
	intRef := g.NewSimpleReference[int]()
	intRef.Set("SOME-CONTEXT.VALUE", 100)

	// create equality guard
	eqGuard := g.NewSimpleGuardEQ("SOME-CONTEXT.VALUE", intRef)

	// create custom guard from equality guard
	instance := g.NewSimpleCustomCompoundGuard[int]("SOME-CONTEXT")
	instance.SetGuard(&eqGuard)

	result, err := instance.EvaluateWithErr(0)
	a.True(t, result, "result should be true")
	a.Nil(t, err, "err should be nil")

	result, err = instance.EvaluateWithErr(100)
	a.True(t, result, "result should be true")
	a.Nil(t, err, "err should be nil")

	result, err = instance.EvaluateWithErr(101)
	a.True(t, result, "result should be true")
	a.Nil(t, err, "err should be nil")

	instance.SetOnEvaluateWithErr(func(value int) (bool, error) {
		guard, err := instance.GetGuardWithErr("SOME-CONTEXT.VALUE")
		if err != nil {
			return true, nil
		}
		return guard.EvaluateWithErr(value)
	})

	instance.SetOnGetConstraint(func() (map[string]int, error) {
		return instance.GetGuard("SOME-CONTEXT").GetConstraint()
	})

	result, err = instance.EvaluateWithErr(0)
	a.False(t, result, "result should be false")
	a.NotNil(t, err, "err should not be nil")

	result, err = instance.EvaluateWithErr(100)
	a.True(t, result, "result should be true")
	a.Nil(t, err, "err should be nil")

	result, err = instance.EvaluateWithErr(101)
	a.False(t, result, "result should be false")
	a.NotNil(t, err, "err should not be nil")
}

func Test_CustomGetConstraint(t *testing.T) {
	// create reference
	intRef := g.NewSimpleReference[int]()
	intRef.Set("SOME-CONTEXT.VALUE", 100)

	// create equality guard
	eqGuard := g.NewSimpleGuardEQ("SOME-CONTEXT.VALUE", intRef)

	// create custom guard from equality guard
	instance := g.NewSimpleCustomCompoundGuard[int]("SOME-CONTEXT")
	instance.SetGuard(&eqGuard)

	result, err := instance.GetConstraint()
	a.Nil(t, result, "result should be nil")
	a.NotNil(t, err, "err should not be nil")

	instance.SetOnGetConstraint(func() (map[string]int, error) {
		return instance.GetGuard("SOME-CONTEXT.VALUE").GetConstraint()
	})

	result, err = instance.GetConstraint()
	a.NotNil(t, result, "result should not be nil")
	a.Nil(t, err, "err should be nil")
}
