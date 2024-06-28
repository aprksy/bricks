package hashmap_test

import (
	"fmt"
	"testing"

	"github.com/aprksy/bricks/base/collection"
	"github.com/aprksy/bricks/base/collection/hashmap"
	"github.com/aprksy/bricks/base/identity"
	"github.com/stretchr/testify/assert"
)

func TestNewHashmap(t *testing.T) {
	testCases := []struct {
		name string
		oid  int
	}{
		{name: "create hashmap 1", oid: 1},
		{name: "create hashmap 2", oid: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := hashmap.NewSimpleHashmap[int, *identity.SimpleIdentity[int]](tc.oid)

			assert.NotNil(t, instance, "instance should not nil")
			assert.NotNil(t, instance.SimpleIdentity, "instance's SimpleIdentity should not nil")
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name  string
		oid   int
		err   error
		count int
	}{
		{name: "add element 1", oid: 1, err: nil, count: 1},
		{name: "add element 2", oid: 2, err: nil, count: 2},
		{name: "add element 2", oid: 2, err: fmt.Errorf(collection.ErrElementExists), count: 2},
	}

	instance := hashmap.NewSimpleHashmap[int, *identity.SimpleIdentity[int]](1)
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := instance.Add(identity.NewSimpleIdentity(tc.oid, "element-type", nil))
			switch i {
			case 0, 1:
				assert.Nil(t, err, "err should be nil")
				assert.Equal(t, instance.Size(), tc.count, fmt.Sprintf("instance.Size() should equal %d", tc.count))
			default:
				assert.NotNil(t, err, "err should not be nil")
				assert.EqualError(t, err, collection.ErrElementExists, fmt.Sprintf("err should equal '%s'", tc.err.Error()))
				assert.Equal(t, instance.Size(), tc.count, fmt.Sprintf("instance.Size() should equal %d", tc.count))
			}
		})
	}
}
