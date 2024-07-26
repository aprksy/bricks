package guard_test

import (
	"testing"

	"github.com/aprksy/bricks/base/guard"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleReference(t *testing.T) {
	intRef := guard.NewSimpleReference[int]()
	assert.NotNil(t, intRef, "intRef should not be nil")
}

func TestSetGet(t *testing.T) {
	testCases := []struct {
		name     string
		positive bool
		setkey   string
		getkey   string
		value    int
	}{
		{name: "correct", positive: true, setkey: "key1", getkey: "key1", value: 1},
		{name: "correct", positive: true, setkey: "key2", getkey: "key2", value: 2},
		{name: "incorrect", positive: false, setkey: "key3", getkey: "key8", value: 3},
		{name: "incorrect", positive: false, setkey: "key4", getkey: "key9", value: 4},
	}

	intRef := guard.NewSimpleReference[int]()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			intRef.Set(tc.setkey, tc.value)
			value, err := intRef.Get(tc.getkey)
			if tc.positive {
				assert.Nil(t, err, "err should be nil")
				assert.Equal(t, tc.value, *value, "*value should equals tc.value")
			} else {
				assert.NotNil(t, err, "err should not be nil")
				assert.Nil(t, value, "value should be nil")
			}
		})
	}
}
