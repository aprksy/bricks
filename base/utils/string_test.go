package utils_test

import (
	"fmt"
	"testing"

	"github.com/aprksy/bricks/base/utils"
	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	testCases := []struct {
		name   string
		length int
		result string
		err    error
	}{
		{name: "negative length", length: -1, result: "", err: fmt.Errorf("invalid length")},
		{name: "zero length", length: 0, result: "", err: fmt.Errorf("invalid length")},
		{name: "one length", length: 1, err: nil},
		{name: "2 length", length: 2, err: nil},
		{name: "5 length", length: 5, err: nil},
		{name: "17 length", length: 17, err: nil},
		{name: "64 length", length: 64, err: nil},
		{name: "1000 length", length: 1000, err: nil},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := utils.RandStr(tc.length)

			switch i {
			case 0, 1:
				assert.Empty(t, result, "result should be empty")
				assert.EqualErrorf(t, err, "invalid length", "err should equal 'invalid length'")
			default:
				assert.Nil(t, err, "err should be nil")
				assert.Equal(t, tc.length, len(result), fmt.Sprintf("result length should equal %d", tc.length))
				assert.Regexp(t, "[a-z,0-9]+", result, "should contains random alphanumerics")
			}
		})
	}
}
