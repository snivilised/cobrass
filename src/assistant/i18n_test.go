package assistant_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"

	"github.com/snivilised/cobrass/src/assistant"
)

var _ = Describe("i18n", func() {
	Context("UseTag", func() {
		When("given: tag is supported", func() {
			It("🧪 should: not return error", func() {
				us := language.MustParse("en-US")
				Expect(assistant.UseTag(us)).Error().To(BeNil())
				Expect(assistant.GetLanguageInfo().Current).To(Equal(us))
			})
		})

		When("given: tag is supported", func() {
			It("🧪 should: return error", func() {
				es := language.MustParse("es")
				Expect(assistant.UseTag(es)).Error().ToNot(BeNil())
			})
		})
	})
})
