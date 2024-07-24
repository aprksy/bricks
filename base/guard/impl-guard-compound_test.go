package guard_test

import (
	"testing"

	g "github.com/aprksy/bricks/base/guard"
	a "github.com/stretchr/testify/assert"
)

func Test_NewSimpleCompoundGuard(t *testing.T) {
	instance := g.NewSimpleCompoundGuard[int]("key1")
	a.NotNil(t, instance, "result should not be nil")
	a.Equalf(t, "key1", instance.Id(), "Id() should equal 'key1'")
	a.True(t, instance.Evaluate(123), "should always true")

	evalResult, err := instance.EvaluateWithErr(123)
	a.True(t, evalResult, "should always true")
	a.Nil(t, err, "should always nil")

	constraint, err := instance.GetConstraint()
	a.Nil(t, constraint, "should always be nil")
	a.NotNil(t, err, "should not be nil")
}

func Test_GuardOperation(t *testing.T) {
	guard := g.NewSimpleGuardBase[int]("key1", nil)
	instance := g.NewSimpleCompoundGuard[int]("key1")

	instance.SetGuard(&guard)
	result := instance.GetGuard("key1")
	a.NotNil(t, result, "should not be nil")

	result, err := instance.GetGuardWithErr("key1")
	a.NotNil(t, result, "should not be nil")
	a.Nil(t, err, "err should be nil")

	instance.ResetGuard("key1")
	instance.ClearGuard()

	result = instance.GetGuard("key1")
	a.Nil(t, result, "should be nil")

	result, err = instance.GetGuardWithErr("key1")
	a.Nil(t, result, "should be nil")
	a.NotNil(t, err, "err not should be nil")
}
