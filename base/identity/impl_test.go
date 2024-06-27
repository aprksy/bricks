package identity_test

import (
	"fmt"
	"testing"

	id "github.com/aprksy/bricks/base/identity"
	"github.com/stretchr/testify/assert"
)

func TestNewIdentity(t *testing.T) {
	type TestCase[T comparable] struct {
		name     string
		id       T
		typeName string
		onInfo   func(t string, id T) string
	}

	onInfo1 := func(t string, id int) string { return fmt.Sprintf("%s (%d)", t, id) }
	onInfo2 := func(t string, id int) string { return fmt.Sprintf("%s/%d", t, id) }
	testCases := []TestCase[int]{
		{"onInfo = nil", 1, "some-type-int", nil},
		{"onInfo = not nil", 2, "some-type-int", func(t string, id int) string { return fmt.Sprintf("%s/%d", t, id) }},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := id.NewSimpleIdentity(tc.id, tc.typeName, tc.onInfo)
			assert.NotNil(t, instance, "instance should not nil")
			assert.Equal(t, instance.Id(), tc.id, "Id() should equal 1")
			assert.Equal(t, instance.TypeName(), tc.typeName, "TypeName() should equal "+fmt.Sprintf("'%s'", tc.typeName))
			if i == 0 {
				assert.Equal(t, instance.InstanceInfo(), onInfo1(tc.typeName, tc.id), "InstanceInfo() should equal "+
					onInfo1(tc.typeName, tc.id))
			} else {
				assert.Equal(t, instance.InstanceInfo(), onInfo2(tc.typeName, tc.id), "InstanceInfo() should equal "+
					onInfo2(tc.typeName, tc.id))
			}
		})
	}
}
