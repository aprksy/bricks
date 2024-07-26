package guard_test

import (
	"testing"

	g "github.com/aprksy/bricks/base/guard"
	a "github.com/stretchr/testify/assert"
)

func Test_NewSimpleGuardEQ(t *testing.T) {
	type args struct {
		id        string
		reference g.ReferenceGetter[int]
	}
	tests := []struct {
		name string
		args args
	}{
		{"correct", args{"guard1", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := g.NewSimpleGuardEQ(tt.args.id, tt.args.reference)
			a.NotNil(t, instance, "instance should not be nil")
		})
	}
}

func Test_SimpleGuardEQ_Evaluate(t *testing.T) {
	type args struct {
		value int
	}

	intRef := g.NewSimpleReference[int]()
	intRef.Set("key1", 1)
	emptyRef := g.NewSimpleReference[int]()
	intGuard := g.NewSimpleGuardEQ[int]("key1", intRef)
	emptyGuard1 := g.NewSimpleGuardEQ[int]("key1", nil)
	emptyGuard2 := g.NewSimpleGuardEQ[int]("key1", emptyRef)

	testCases := []struct {
		name     string
		args     args
		expected bool
		guard    *g.SimpleGuardEQ[int]
		comment  string
	}{
		{"correct", args{1}, true, &emptyGuard1, "result should be true"},
		{"correct", args{1}, true, &emptyGuard2, "result should be true"},
		{"correct", args{1}, true, &intGuard, "result should be true"},
		{"incorrect", args{2}, false, &intGuard, "result should be false"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.guard.Evaluate(tc.args.value)
			a.Equal(t, tc.expected, result, tc.comment)
		})
	}
}

func Test_SimpleGuardEQ_EvaluateWithErr(t *testing.T) {
	type args struct {
		value int
	}

	intRef := g.NewSimpleReference[int]()
	intRef.Set("key1", 1)
	emptyRef := g.NewSimpleReference[int]()
	intGuard := g.NewSimpleGuardEQ[int]("key1", intRef)
	emptyGuard1 := g.NewSimpleGuardEQ[int]("key1", nil)
	emptyGuard2 := g.NewSimpleGuardEQ[int]("key1", emptyRef)

	testCases := []struct {
		name     string
		args     args
		expected bool
		gotError bool
		guard    *g.SimpleGuardEQ[int]
		comment  string
	}{
		{"correct", args{1}, true, false, &emptyGuard1, "result should be true"},
		{"correct", args{1}, true, false, &emptyGuard2, "result should be true"},
		{"correct", args{1}, true, false, &intGuard, "result should be true"},
		{"incorrect", args{2}, false, true, &intGuard, "result should be false"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.guard.EvaluateWithErr(tc.args.value)
			a.Equal(t, tc.expected, result, tc.comment)
			if !tc.gotError {
				a.Nil(t, err, "err should be nil")
			} else {
				a.NotNil(t, err, "err should not be nil")
			}
		})
	}
}

func Test_SimpleGuardEQ_GetConstraint(t *testing.T) {
	intRef := g.NewSimpleReference[int]()
	intRef.Set("key1", 1)
	emptyRef := g.NewSimpleReference[int]()
	intGuard := g.NewSimpleGuardEQ[int]("key1", intRef)
	emptyGuard1 := g.NewSimpleGuardEQ[int]("key1", nil)
	emptyGuard2 := g.NewSimpleGuardEQ[int]("key1", emptyRef)

	testCases := []struct {
		name     string
		expected map[string]int
		gotError bool
		guard    *g.SimpleGuardEQ[int]
	}{
		{"case 01", map[string]int{"key1": 1}, false, &intGuard},
		{"case 02", nil, true, &emptyGuard1},
		{"case 03", nil, true, &emptyGuard2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.guard.GetConstraint()
			a.Equal(t, tc.expected, result, "result should equal the expected")
			if !tc.gotError {
				a.Nil(t, err, "err should be nil")
			} else {
				a.NotNil(t, err, "err should not be nil")
			}
		})
	}
}
