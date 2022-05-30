package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenNewKeyDoesNotExistInsertSafePositional(t *testing.T) {

	generic := make(GenericParameterSet)
	generic["colour"] = "blue"
	generic["shape"] = "circle"

	InsertSafePositional(generic, "size", "large")
	assert.True(t, true, "insertion should not cause panic")
}

// comsider https://onsi.github.io/ginkgo/ for more sophisticate unit testing in go
//

func TestGivenNewKeyDoesExistsInsertSafePositional(t *testing.T) {
	generic := make(GenericParameterSet)
	generic["colour"] = "blue"
	generic["shape"] = "circle"

	func() {
		defer func() {
			if r := recover(); r != nil {
				assert.True(t, true, "insertion should cause panic")
			}
		}()
		InsertSafePositional(generic, "shape", "square")
		assert.Fail(t, "insertion should cause panic")
	}()
}
