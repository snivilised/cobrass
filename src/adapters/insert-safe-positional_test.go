package adapters_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/src/adapters"
)

var _ = Describe("InsertSafePositional", func() {
	Context("given: new flag does not exist", func() {
		It("🧪 should: insert ok", func() {
			generic := make(adapters.GenericParamSet)
			generic["colour"] = "blue"
			generic["shape"] = "circle"

			const value = "large"
			const flag = "size"

			adapters.InsertSafePositional(generic, flag, value)

			actual, ok := generic[flag]

			message := fmt.Sprintf("❌ inserted flag '%v' of value '%v' but should be present but isn't",
				flag, actual)
			Expect(ok).To(BeTrue(), message)

			message = fmt.Sprintf("❌ inserted param '%v' is '%v' but should be: '%v'",
				flag, actual, value)
			Expect(actual).To(Equal(value), message)
		})
	})

	Context("given: inserted flag alrady exists", func() {
		It("🧪 should: ", func() {
			generic := make(adapters.GenericParamSet)
			generic["colour"] = "blue"
			generic["shape"] = "circle"

			func() {
				defer func() {
					if r := recover(); r != nil {
						Expect(true).To(BeTrue(), "✔️ expected panic handled")
					}
				}()
				adapters.InsertSafePositional(generic, "shape", "square")
				Fail("❌ insertion should cause panic")
			}()
		})
	})
})
