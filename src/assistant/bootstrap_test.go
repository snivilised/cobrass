package assistant_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/translate"
	"golang.org/x/text/language"
)

var _ = Describe("Bootstrap", func() {

	Context("widget command", func() {
		It("🧪 should: invoke without error", func() {

			directory, _ := filepath.Abs("../../src/assistant/internal/l10n/out")
			bootstrap := assistant.Bootstrap{}
			bootstrap.Execute(func(detector assistant.LocaleDetector) {
				translate.Initialise(func(o *translate.LanguageInitOptions) {
					o.TranslationFilename = "active.en-US.json"
					o.Path = directory
					o.Detected = language.AmericanEnglish
				})

			})
			Expect(true)
		})
	})
})
