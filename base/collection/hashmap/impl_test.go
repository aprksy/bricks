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
		oid  uint
	}{
		{name: "create hashmap 1", oid: 1},
		{name: "create hashmap 2", oid: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := hashmap.NewSimpleHashmap[uint, *identity.SimpleIdentity[uint]]()

			assert.NotNil(t, instance, "instance should not nil")
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

	instance := hashmap.NewSimpleHashmap[int, int]()

	assert.Panicsf(t, func() { instance.Add(1) }, "not implemented", "call to instance.Add() should cause panic")

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := instance.AddWithId(tc.oid, tc.oid)
			switch i {
			case 0, 1:
				size, _ := instance.Size()
				assert.Nil(t, err, "err should be nil")
				assert.Equal(t, tc.count, size, fmt.Sprintf("instance.Size() should equal %d", tc.count))
			default:
				size, _ := instance.Size()
				assert.NotNil(t, err, "err should not be nil")
				assert.EqualError(t, err, collection.ErrElementExists, fmt.Sprintf("err should equal '%s'", tc.err.Error()))
				assert.Equal(t, tc.count, size, fmt.Sprintf("instance.Size() should equal %d", tc.count))
			}
		})
	}
}

func TestElement(t *testing.T) {
	testCases := []struct {
		name string
		oid  int
		err  error
	}{
		{name: "get element 1", oid: 1, err: nil},
		{name: "get element 2", oid: 2, err: nil},
		{name: "get element 3", oid: 3, err: fmt.Errorf(collection.ErrElementNotFound)},
	}

	instance := hashmap.NewSimpleHashmap[int, int]()
	for i, tc := range testCases {
		if i == 2 {
			break
		}
		instance.AddWithId(tc.oid, tc.oid)
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptrE, err := instance.Element(tc.oid)
			switch i {
			case 0, 1:
				assert.Nil(t, err, "err should be nil")
				assert.NotNil(t, ptrE, "element should not be nil")
			default:
				assert.Nil(t, ptrE, "element should be nil")
				assert.NotNil(t, err, "err should not be nil")
				assert.EqualError(t, err, collection.ErrElementNotFound, fmt.Sprintf("err should equal '%s'", tc.err.Error()))
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name string
		oid  int
		err  error
	}{
		{name: "rm element 1", oid: 1, err: nil},
		{name: "rm element 2", oid: 2, err: nil},
		{name: "rm element 3", oid: 3, err: fmt.Errorf(collection.ErrElementNotFound)},
	}

	instance := hashmap.NewSimpleHashmap[int, int]()
	elements := []int{}
	for i, tc := range testCases {
		elements = append(elements, tc.oid)
		if i == 2 {
			break
		}
		instance.AddWithId(tc.oid, tc.oid)
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := instance.Remove(elements[i])
			switch i {
			case 0, 1:
				assert.Nil(t, err, "err should be nil")
			default:
				assert.NotNil(t, err, "err should not be nil")
				assert.EqualError(t, err, collection.ErrElementNotFound, fmt.Sprintf("err should equal '%s'", tc.err.Error()))
			}
		})
	}

	elements = []int{}
	for i, tc := range testCases {
		elements = append(elements, tc.oid)
		if i == 2 {
			break
		}
		instance.AddWithId(tc.oid, tc.oid)
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := instance.RemoveById(tc.oid)
			switch i {
			case 0, 1:
				assert.Nil(t, err, "err should be nil")
			default:
				assert.NotNil(t, err, "err should not be nil")
				assert.EqualError(t, err, collection.ErrElementNotFound, fmt.Sprintf("err should equal '%s'", tc.err.Error()))
			}
		})
	}
}

func TestClear(t *testing.T) {
	testCases := []struct {
		name string
		oid  int
		err  error
	}{
		{name: "get element 1", oid: 1, err: nil},
		{name: "get element 2", oid: 2, err: nil},
	}

	instance := hashmap.NewSimpleHashmap[int, int]()
	for _, tc := range testCases {
		instance.AddWithId(tc.oid, tc.oid)
	}

	t.Run("clear", func(t *testing.T) {
		err := instance.Clear()
		size, _ := instance.Size()
		elements, _ := instance.Elements()
		assert.Nil(t, err, "err should be nil")
		assert.Zero(t, size, "Size() should be 0")
		assert.Empty(t, elements, "Elements() should be empty")
	})
}

func TestElements(t *testing.T) {
	testCases := []struct {
		name string
		oid  int
		err  error
	}{
		{name: "element 1", oid: 1, err: nil},
		{name: "element 2", oid: 2, err: nil},
		{name: "element 3", oid: 3, err: nil},
		{name: "element 4", oid: 3, err: fmt.Errorf(collection.ErrElementNotFound)},
	}

	instance := hashmap.NewSimpleHashmap[int, int]()
	elements := []int{}
	for i, tc := range testCases {
		elements = append(elements, tc.oid)
		if i == 3 {
			break
		}
		instance.AddWithId(tc.oid, tc.oid)
	}

	t.Run("elements", func(t *testing.T) {
		elems, _ := instance.Elements()
		size, _ := instance.Size()
		assert.NotNil(t, elems, "Elements() should not be nil")
		assert.Equal(t, size, len(elems), "Size() should equal len(elems)")

		instance.Remove(elements[0])
		elems, _ = instance.Elements()
		size, _ = instance.Size()
		assert.Equal(t, size, len(elems), "Size() should equal len(elems)")

		el, err := instance.Element(1)
		assert.EqualErrorf(t, err, collection.ErrElementNotFound, fmt.Sprintf("err should equal %s", collection.ErrElementNotFound))
		assert.Nil(t, el, "element should be nil")
	})
}
