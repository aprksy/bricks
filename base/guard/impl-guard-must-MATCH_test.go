package guard_test

import (
	"testing"

	g "github.com/aprksy/bricks/base/guard"
	a "github.com/stretchr/testify/assert"
)

func Test_NewSimpleGuardMatch(t *testing.T) {
	type args struct {
		id        string
		reference g.ReferenceGetter[string]
	}
	tests := []struct {
		name string
		args args
	}{
		{"correct", args{"guard1", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := g.NewSimpleGuardMatch(tt.args.id, tt.args.reference)
			a.NotNil(t, instance, "instance should not be nil")
		})
	}
}

func Test_SimpleGuardMatch_Evaluate(t *testing.T) {
	type args struct {
		value string
	}

	strRef := g.NewSimpleReference[string]()
	strRef.Set("key1", "^[a-z,0-9,\\-,\\.,_]*@[a-z,0-9,-,\\.,_]*.[a-z]$") // simple email address pattern
	strRef.Set("key2", "a(b")
	emptyRef := g.NewSimpleReference[string]()
	strGuard := g.NewSimpleGuardMatch("key1", strRef)
	emptyGuard1 := g.NewSimpleGuardMatch("key1", nil)
	emptyGuard2 := g.NewSimpleGuardMatch("key1", emptyRef)
	errGuard1 := g.NewSimpleGuardMatch("key2", strRef)

	testCases := []struct {
		name     string
		args     args
		expected bool
		guard    *g.SimpleGuardMatch
		comment  string
	}{
		{"correct", args{"user@example.com"}, true, &emptyGuard1, "result should be true"},
		{"correct", args{"user@example.com"}, true, &emptyGuard2, "result should be true"},
		{"correct", args{"user@example.com"}, true, &strGuard, "result should be true"},
		{"incorrect", args{"12345"}, false, &strGuard, "result should be false"},
		{"incorrect", args{"12345"}, false, &errGuard1, "result should be false"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.guard.Evaluate(tc.args.value)
			a.Equal(t, tc.expected, result, tc.comment)
		})
	}
}

func Test_SimpleGuardMatch_EvaluateWithErr(t *testing.T) {
	type args struct {
		value string
	}

	strRef := g.NewSimpleReference[string]()
	strRef.Set("key1", "^[a-z,0-9,\\-,\\.,_]*@[a-z,0-9,-,\\.,_]*.[a-z]$")
	strRef.Set("key2", "a(b")
	emptyRef := g.NewSimpleReference[string]()
	strGuard := g.NewSimpleGuardMatch("key1", strRef)
	emptyGuard1 := g.NewSimpleGuardMatch("key1", nil)
	emptyGuard2 := g.NewSimpleGuardMatch("key1", emptyRef)
	errGuard1 := g.NewSimpleGuardMatch("key2", strRef)

	testCases := []struct {
		name     string
		args     args
		expected bool
		gotError bool
		guard    *g.SimpleGuardMatch
		comment  string
	}{
		{"correct", args{"user@example.com"}, true, false, &emptyGuard1, "result should be true"},
		{"correct", args{"user@example.com"}, true, false, &emptyGuard2, "result should be true"},
		{"correct", args{"user@example.com"}, true, false, &strGuard, "result should be true"},
		{"incorrect", args{"12345"}, false, true, &strGuard, "result should be false"},
		{"incorrect", args{"12345"}, false, true, &errGuard1, "result should be false"},
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

func Test_SimpleGuardMatch_GetConstrastr(t *testing.T) {
	strRef := g.NewSimpleReference[string]()
	strRef.Set("key1", "^[a-z,0-9,\\-,\\.,_]*@[a-z,0-9,-,\\.,_]*.[a-z]$")
	emptyRef := g.NewSimpleReference[string]()
	strGuard := g.NewSimpleGuardMatch("key1", strRef)
	emptyGuard1 := g.NewSimpleGuardMatch("key1", nil)
	emptyGuard2 := g.NewSimpleGuardMatch("key1", emptyRef)

	testCases := []struct {
		name     string
		expected map[string]string
		gotError bool
		guard    *g.SimpleGuardMatch
	}{
		{"case 01", map[string]string{"key1": "^[a-z,0-9,\\-,\\.,_]*@[a-z,0-9,-,\\.,_]*.[a-z]$"}, false, &strGuard},
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
