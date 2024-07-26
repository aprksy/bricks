package guard_test

import (
	"testing"

	g "github.com/aprksy/bricks/base/guard"
	a "github.com/stretchr/testify/assert"
)

func Test_NewSimpleGuardable(t *testing.T) {
	instance, err := g.NewSimpleGuardable[int](nil)
	a.Nil(t, instance, "instance should be nil")
	a.NotNil(t, err, "err should not be nil")

	instance, err = g.NewSimpleGuardable[int](&MockGuard[int]{})
	a.NotNil(t, instance, "instance should not be nil")
	a.Nil(t, err, "err should be nil")
}

func Test_Allow(t *testing.T) {
	mock := &MockGuard[int]{}
	instance, _ := g.NewSimpleGuardable[int](mock)

	mock.On("Evaluate", 100).Return(false)
	instance.Allow(100)
	mock.AssertExpectations(t)
}

func Test_AllowWithErr(t *testing.T) {
	mock := &MockGuard[int]{}
	instance, _ := g.NewSimpleGuardable[int](mock)

	mock.On("EvaluateWithErr", 100).Return(false, nil)
	instance.AllowWithErr(100)
	mock.AssertExpectations(t)
}

func Test_GetConstraint(t *testing.T) {
	mock := &MockGuard[int]{}
	instance, _ := g.NewSimpleGuardable[int](mock)

	mock.On("GetConstraint").Return(map[string]int{}, nil)
	instance.GetConstraint()
	mock.AssertExpectations(t)
}
