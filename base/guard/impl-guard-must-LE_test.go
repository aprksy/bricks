package guard_test

import (
	"testing"

	g "github.com/aprksy/bricks/base/guard"
	a "github.com/stretchr/testify/assert"
)

func Test_NewSimpleGuardLE(t *testing.T) {
	type args struct {
		id        string
		reference g.ReferenceGetter[int]
	}
	tests := []struct {
		name string
		args args
	}{
		{"case 01", args{"guard1", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := g.NewSimpleGuardLE(tt.args.id, tt.args.reference)
			a.NotNil(t, instance, "instance should not be nil")
		})
	}
}

func Test_SimpleGuardLE_Evaluate(t *testing.T) {
	type args struct {
		value int
	}

	intRef := g.NewSimpleReference[int]()
	intRef.Set("key1", 1)
	emptyRef := g.NewSimpleReference[int]()
	intGuard := g.NewSimpleGuardLE[int]("key1", intRef)
	emptyGuard1 := g.NewSimpleGuardLE[int]("key1", nil)
	emptyGuard2 := g.NewSimpleGuardLE[int]("key1", emptyRef)

	testCases := []struct {
		name     string
		args     args
		expected bool
		guard    *g.SimpleGuardLE[int]
	}{
		{"case 01", args{1}, true, &emptyGuard1},
		{"case 02", args{1}, true, &emptyGuard2},
		{"case 03", args{1}, true, &intGuard},
		{"case 04", args{2}, false, &intGuard},
		{"case 04", args{0}, true, &intGuard},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.guard.Evaluate(tc.args.value)
			a.Equalf(t, tc.expected, result, "result should equal %t", tc.expected)
		})
	}
}

func Test_SimpleGuardLE_EvaluateWithErr(t *testing.T) {
	type args struct {
		value int
	}

	intRef := g.NewSimpleReference[int]()
	intRef.Set("key1", 1)
	emptyRef := g.NewSimpleReference[int]()
	intGuard := g.NewSimpleGuardLE[int]("key1", intRef)
	emptyGuard1 := g.NewSimpleGuardLE[int]("key1", nil)
	emptyGuard2 := g.NewSimpleGuardLE[int]("key1", emptyRef)

	testCases := []struct {
		name     string
		args     args
		expected bool
		gotError bool
		guard    *g.SimpleGuardLE[int]
	}{
		{"case 01", args{1}, true, false, &emptyGuard1},
		{"case 02", args{1}, true, false, &emptyGuard2},
		{"case 03", args{1}, true, false, &intGuard},
		{"case 04", args{2}, false, true, &intGuard},
		{"case 04", args{0}, true, false, &intGuard},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.guard.EvaluateWithErr(tc.args.value)
			a.Equalf(t, tc.expected, result, "result should equal %t", tc.expected)
			if !tc.gotError {
				a.Nil(t, err, "err should be nil")
			} else {
				a.NotNil(t, err, "err should not be nil")
			}
		})
	}
}

func Test_SimpleGuardLE_GetConstraint(t *testing.T) {
	intRef := g.NewSimpleReference[int]()
	intRef.Set("key1", 1)
	emptyRef := g.NewSimpleReference[int]()
	intGuard := g.NewSimpleGuardLE[int]("key1", intRef)
	emptyGuard1 := g.NewSimpleGuardLE[int]("key1", nil)
	emptyGuard2 := g.NewSimpleGuardLE[int]("key1", emptyRef)

	testCases := []struct {
		name     string
		expected map[string]int
		gotError bool
		guard    *g.SimpleGuardLE[int]
	}{
		{"case 01", map[string]int{"key1": 1}, false, &intGuard},
		{"case 02", nil, true, &emptyGuard1},
		{"case 03", nil, true, &emptyGuard2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.guard.GetConstraint()
			a.Equal(t, tc.expected, result, "result should NEual the expected")
			if !tc.gotError {
				a.Nil(t, err, "err should be nil")
			} else {
				a.NotNil(t, err, "err should not be nil")
			}
		})
	}
}
