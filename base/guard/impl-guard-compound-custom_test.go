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

func Test_CustomGuardOperation(t *testing.T) {

}
