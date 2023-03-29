package translate_test

import (
	"path/filepath"
	"strings"
	"text/template"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/internal/l10n"
	"github.com/snivilised/cobrass/src/assistant/translate"
)

func ShowDaysRemaining(p *message.Printer, days int) string {
	return p.Sprintf("You have %d days remaining", days)
}

func ShowInternationalisation(p *message.Printer, to string) string {
	return p.Sprintf("greetings '%v', welcome to internationalisation", to)
}

var _ = Describe("i18n", func() {

	var languages *translate.LanguageInfo
	var printer *message.Printer

	BeforeEach(func() {
		directory, _ := filepath.Abs("../../assistant/internal/l10n/out")
		bootstrap := assistant.Bootstrap{}
		bootstrap.Execute(func(detector assistant.LocaleDetector) {
			translate.Initialise(func(o *translate.LanguageInitOptions) {
				o.Detected = language.BritishEnglish
				o.Path = directory
				o.TranslationFilename = "active.en-US.json"
			})
		})
	})

	Context("UseTag", func() {
		When("given: tag is supported", func() {
			It("🧪 should: not return error", func() {
				us := language.AmericanEnglish
				Expect(translate.UseTag(us)).Error().To(BeNil())
				Expect(translate.GetLanguageInfo().Current).To(Equal(us))
			})

			It("🧪 should: localise in requested non default language", func() {
				_ = translate.UseTag(language.AmericanEnglish)
				data := l10n.LessThanOptValidationTemplData{
					RelationalOV: l10n.RelationalOV{Flag: "flag", Value: 1, Threshold: 2},
				}

				_, tag, _ := translate.GetLocaliser().LocalizeWithTag(&i18n.LocalizeConfig{
					DefaultMessage: data.Message(),
					TemplateData:   data,
				})
				Expect(tag.String()).To(Equal(language.AmericanEnglish.String()))
			})
		})

		When("given: tag is NOT supported", func() {
			It("🧪 should: return error", func() {
				es := language.MustParse("es")
				Expect(translate.UseTag(es)).Error().ToNot(BeNil())
			})
		})
	})

	Context("different ways of creating equivalent language tags", func() {
		It("🧪 should: be equal", func() {
			// Just comprehension tests
			//
			Expect(language.MustParse("en-GB")).To(Equal(language.BritishEnglish))
			Expect(language.MustParse("en-US")).To(Equal(language.AmericanEnglish))
		})
	})

	Context("manually defined translation", Ordered, func() {
		BeforeAll(func() {
			_ = message.SetString(language.AmericanEnglish,
				"greetings '%v', welcome to internationalisation",
				"greetings '%v', welcome to internationalization",
			)
		})

		When("using default language", func() {
			It("🧪 should: show un-translated text", func() {
				printer = message.NewPrinter(language.BritishEnglish)

				msg := ShowInternationalisation(printer, "earthlings")
				const expected = "greetings 'earthlings', welcome to internationalisation"
				Expect(msg).To(Equal(expected))
			})
		})

		When("using alternative language", func() {
			It("🧪 should: show translated text", func() {
				printer = message.NewPrinter(language.AmericanEnglish)

				_ = translate.UseTag(language.AmericanEnglish)
				msg := ShowInternationalisation(printer, "earthlings")
				const expected = "greetings 'earthlings', welcome to internationalization"
				Expect(msg).To(Equal(expected))
			})
		})
	})

	Describe("pluralisation example", Ordered, func() {
		BeforeAll(func() {
			languages = translate.GetLanguageInfo()

			// "one"/ "other" are known as selectors
			// "You have %d day remaining" is the message
			// plural.Selectf takes pairs of selectors/messages -> a case
			// but
			// what is the 1 here and what does %[1]d mean?
			// and, >100 / <10 don't seem to work here
			//
			_ = message.Set(languages.Default, "You have %d days remaining", plural.Selectf(1, "%d",
				"one", "You have %d day remaining",
				"other", "You have %[1]d days remaining",
				">100", "You have %d days remaining, over the maximum",
				"<10", "You have %d days remaining, less than the minimum",
			))
		})

		It("🧪 should: show pluralised text", func() {
			printer = message.NewPrinter(language.BritishEnglish)

			GinkgoWriter.Println(ShowDaysRemaining(printer, 1))
			GinkgoWriter.Println(ShowDaysRemaining(printer, 22))

			_ = translate.UseTag(language.AmericanEnglish)
			GinkgoWriter.Println(ShowDaysRemaining(printer, 333))
		})
	})

	Describe("template example", Ordered, func() {

		BeforeAll(func() {
			// This combines pluralisation and localisation requirements
			// But note here we are using format specifiers, even though
			// dependent code is using templates, should still work?
			//
			// This does NOT work, in the test table:
			// "with impractical translation strategy (don't do this)"
			//
			// because the printer is being passed
			// a fully created literal string created from a template and
			// is not aware of the variable. It will only work when it
			// is presented with a format string with place holders.
			//
			_ = message.SetString(language.AmericanEnglish,
				"There are %d colours in this rainbow", // <-- key
				"There are %d colors in this rainbow",  // <-- message
			)

			// This works in test table:
			// "with impractical translation strategy (don't do this)"
			// but is not practical. Just included here for illustration
			//
			// Clearly, we can't do this. The printer has to be aware of the variable
			// substitution. We need to improve the way we we do this.
			//
			_ = message.SetString(language.AmericanEnglish,
				"There are 7 colours in this rainbow",
				"There are 7 colors in this rainbow",
			)
		})

		// https://appdividend.com/2019/11/27/golang-template-example-package-template-in-golang/
		//
		// So to templatise strings we need the following components for each individual string/message
		// * a struct defining the variable entities. By our own convention, we'll define the struct
		// with a TemplData suffix
		// * a string builder instance
		//
		// Steps to creating a templated string
		//
		// - create an instance of the template data
		// - create a template string containing embedded variables that map to member(s) of our Templ
		// - create a parsed template with a kebab-case name, from a template body string (the body
		// contains the body of the text with place holders that match the template data members)
		// - execute the parsed template passing in the string builder
		// - check the execute result
		// - access the builder content, and use at will
		// - nb, this work-fow does not include translation, but the result of the template execution
		// can be printed in language context using a printer
		//
		// All of the above should be wrapped up into am author object which contains the
		// standard print methods, perhaps id it can be made to make sense
		//

		// Actually, using a TemplData like this is a little indulgent. There is little use
		// in defining a context specific struct like DaysRemainingTemplData which simply
		// contains a single number, but is left in here for the purposes of illustration.
		//

		type DaysRemainingTemplData struct {
			Days int
		}

		It("🧪 should: get templated string", func() {
			builder := strings.Builder{}
			days := DaysRemainingTemplData{66}

			templ, err := template.New("you-have-n-days-remaining").Parse("You have {{ .Days}} days remaining")
			if err != nil {
				panic(err)
			}

			err = templ.Execute(&builder, days)
			if err != nil {
				panic(err)
			}
			expected := "You have 66 days remaining"
			Expect(builder.String()).To(Equal(expected))
		})
	})

	Context("go-i18n", func() {
		When("using map of any", func() {
			It("🧪 should: translate", func() {
				violationMsg := &i18n.Message{
					ID:    "ov-failed-out-of-range",
					Other: "({{.Flag}}): option validation failed, '{{.Value}}', out of range: [{{.Lo}}]..[{{.Hi}}]",
				}

				// using map of any, is more concise, but not type safe. Any coding errors
				// won't make themselves apparent until runtime.
				//
				localised := translate.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: violationMsg,
					TemplateData:   map[string]any{"Flag": "Strike", "Value": 999, "Lo": 1, "Hi": 99},
				})
				expected := "(Strike): option validation failed, '999', out of range: [1]..[99]"
				Expect(localised).To(Equal(expected))
			})
		})

		When("using template", func() {
			It("🧪 should: translate", func() {
				violationMsg := &i18n.Message{
					ID:    "ov-failed-out-of-range",
					Other: "({{.Flag}}): option validation failed, '{{.Value}}', out of range: [{{.Lo}}]..[{{.Hi}}]",
				}

				// using a template is not as concise as a map of any, but it is type-safe
				// so is the much preferred solution.
				//
				localised := translate.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: violationMsg,
					TemplateData: l10n.WithinOptValidationTemplData{
						OutOfRangeOV: l10n.OutOfRangeOV{Flag: "Strike", Value: 999, Lo: 1, Hi: 99},
					},
				})
				expected := "(Strike): option validation failed, '999', out of range: [1]..[99]"
				Expect(localised).To(Equal(expected))
			})
		})
	})
})
