package guard_test

import (
	"testing"

	"github.com/aprksy/bricks/base/guard"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleGuardBase(t *testing.T) {
	intGuard := guard.NewSimpleGuardBase[int]("guard1", nil)

	assert.NotNil(t, intGuard, "intGuard should not be nil")
	assert.Equal(t, "guard1", intGuard.Id(), "Id() should equal 'guard1'")
	assert.Equal(t, nil, intGuard.Reference(), "Reference() should equal 'nil'")
	assert.Equal(t, true, intGuard.Evaluate(123), "Evaluate() should return true")

	value, err := intGuard.EvaluateWithErr(123)
	assert.Equal(t, true, value, "value should equal true")
	assert.Nil(t, err, "err should be nil")

	refs, err := intGuard.GetConstraint()
	assert.Equal(t, map[string]int{}, refs, "value should equal empty map[string]int")
	assert.Nil(t, err, "err should be nil")
}
