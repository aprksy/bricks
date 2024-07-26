package guard_test

import (
	"cmp"

	g "github.com/aprksy/bricks/base/guard"
	"github.com/stretchr/testify/mock"
)

var _ g.Guard[int] = (*MockGuard[int])(nil)

type MockGuard[T bool | cmp.Ordered] struct {
	mock.Mock
}

// Evaluate implements guard.Guard.
func (m *MockGuard[T]) Evaluate(value T) bool {
	args := m.Called(value)
	return args.Bool(0)
}

// EvaluateWithErr implements guard.Guard.
func (m *MockGuard[T]) EvaluateWithErr(value T) (bool, error) {
	args := m.Called(value)
	return args.Bool(0), nil
}

// GetConstraint implements guard.Guard.
func (m *MockGuard[T]) GetConstraint() (map[string]T, error) {
	m.Called()
	return map[string]T{}, nil
}

// Id implements guard.Guard.
func (m *MockGuard[T]) Id() string {
	m.Called()
	return "some-id"
}
