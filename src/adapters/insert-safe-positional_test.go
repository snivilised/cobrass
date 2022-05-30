package adapters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenNewKeyDoesNotExistInsertSafePositional(t *testing.T) {

	generic := make(GenericParameterSet)
	generic["colour"] = "blue"
	generic["shape"] = "circle"

	const expect = "large"
	const param = "size"
	InsertSafePositional(generic, param, expect)
	assert.True(t, true, "insertion should not cause panic")

	if value, ok := generic[param]; !ok {
		assert.Fail(t, fmt.Sprintf("inserted param '%v' is '%v' but should be: '%v'",
			param, value, expect))
	} else {
		assert.Equal(t, value, expect, fmt.Sprintf("%v", value))
	}
}

// comsider https://onsi.github.io/ginkgo/ for more sophisticate unit testing in go
// provides bdd style of unit testing
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
