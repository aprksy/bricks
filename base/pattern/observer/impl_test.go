package observer_test

import (
	"testing"

	obs "github.com/aprksy/bricks/base/pattern/observer"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleSubject(t *testing.T) {
	testCases := []struct {
		name  string
		oid   uint
		key   string
		value int
	}{
		{name: "subject 1", oid: 1, key: "key-1", value: 1},
		{name: "subject 2", oid: 2, key: "key-2", value: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := obs.NewSimpleSubject[uint, int](tc.oid, tc.key, tc.value)
			assert.NotNilf(t, instance, "instance should not be nil")
		})
	}
}

func TestNewSimpleObserver(t *testing.T) {
	testCases := []struct {
		name  string
		oid   uint
		key   string
		value int
	}{
		{name: "subject 1", oid: 1, key: "key-1", value: 1},
		{name: "subject 2", oid: 2, key: "key-2", value: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := obs.NewSimpleSubject[uint, int](tc.oid, tc.key, tc.value)
			assert.NotNilf(t, instance, "instance should not be nil")
		})
	}
}

func TestSubscribe(t *testing.T) {

}
