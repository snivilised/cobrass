package assistant_test

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/snivilised/cobrass/src/assistant"
)

type OptValidationOutOfRangeTemplData struct {
	Flag  string
	Value any
	Lo    any
	Hi    any
}

func ShowDaysRemaining(p *message.Printer, days int) string {
	return p.Sprintf("You have %d days remaining", days)
}

func ShowInternationalisation(p *message.Printer, to string) string {
	return p.Sprintf("greetings '%v', welcome to internationalisation", to)
}

var _ = Describe("i18n", func() {

	var languages assistant.LanguageInfo
	var printer *message.Printer

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
			message.SetString(language.AmericanEnglish,
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

				assistant.UseTag(language.AmericanEnglish)
				msg := ShowInternationalisation(printer, "earthlings")
				const expected = "greetings 'earthlings', welcome to internationalization"
				Expect(msg).To(Equal(expected))
			})
		})
	})

	Describe("pluralisation example", Ordered, func() {
		BeforeAll(func() {
			languages = assistant.GetLanguageInfo()

			// "one"/ "other" are known as selectors
			// "You have %d day remaining" is the message
			// plural.Selectf takes pairs of selectors/messages -> a case
			// but
			// what is the 1 here and what does %[1]d mean?
			// and, >100 / <10 don't seem to work here
			//
			message.Set(languages.Default, "You have %d days remaining", plural.Selectf(1, "%d",
				"one", "You have %d day remaining",
				"other", "You have %[1]d days remaining",
				">100", "You have %d days remaining, over the maximum",
				"<10", "You have %d days remaining, less than the minium",
			))
		})

		It("🧪 should: show pluralised text", func() {
			printer = message.NewPrinter(language.BritishEnglish)

			GinkgoWriter.Println(ShowDaysRemaining(printer, 1))
			GinkgoWriter.Println(ShowDaysRemaining(printer, 22))

			assistant.UseTag(language.AmericanEnglish)
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
			// "with impractcal translation strategy (don't do this)"
			//
			// because the printer is being passed
			// a fully created literal string created from a template and
			// is not aware of the variable. It will only work when it
			// is presented with a format string with place holders.
			//
			message.SetString(language.AmericanEnglish,
				"There are %d colours in this rainbow", // <-- key
				"There are %d colors in this rainbow",  // <-- message
			)

			// This works in test table:
			// "with impractcal translation strategy (don't do this)"
			// but is not practical. Just included here for illustration
			//
			// Clearly, we can't do this. The printer has to be aware of the variable
			// substitution. We need to improve the way we we do this.
			//
			message.SetString(language.AmericanEnglish,
				"There are 7 colours in this rainbow",
				"There are 7 colors in this rainbow",
			)
		})

		// https://appdividend.com/2019/11/27/golang-template-example-package-template-in-golang/
		//
		// So to templatise strings we need the following components for each individual string/message
		// * a struct defining the variable entities. By our own covention, we'll define the struct
		// with a TemplData suffix
		// * a string builder instance
		//
		// Steps to creating a templated string
		//
		// - create an instance of the template data
		// - create a template string containing embedded variables that map to member(s) of our Templ
		// - create a parsed template with a kebak-case name, from a template body string (the body
		// conatins the body of the text with place holders that match the template data members)
		// - execute the parsed template passing in the string builder
		// - check the execute result
		// - access the builder content, and use at will
		// - nb, this workfow does not include translation, but the result of the template execution
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

		// This defintion is designed as a more generic version of DaysRemainingTemplData
		// which means that it can be used as the backing data source for more templates
		// and is probably the strategy that should be used going forward, to avoid having
		// multiple frivolous definitions. The specificity should be in the parsed
		// template and the template data should be as generic as possible to aid reuse
		// and cut down on excessive definitions.
		//
		type SingleIntTemplData struct {
			No int
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

		DescribeTable("with impractcal translation strategy (don't do this)",
			func(variant, expected string, t language.Tag) {
				builder := strings.Builder{}
				colours := SingleIntTemplData{7}

				templ, err := template.New("there-are-n-colours-in-this-rainbow").Parse("There are {{ .No}} colours in this rainbow")
				if err != nil {
					panic(err)
				}

				err = templ.Execute(&builder, colours)
				if err != nil {
					panic(err)
				}
				source := builder.String()

				printer = message.NewPrinter(t)

				// The reason for the impracticality here is that these tests only work because
				// we have set the translation for explicit value of n (=7). The printer is not
				// being presented with '%v' format field and therefore can't work. Clearly,
				// this is not a solution, see "with fixed translation strategy" for
				// solution.
				//
				translated := printer.Sprintf(source)
				Expect(translated).To(Equal(expected))
			},
			func(variant, expected string, t language.Tag) string {
				return fmt.Sprintf("🧪 should: get text in '%v' ", variant)
			},
			Entry(nil, "british", "There are 7 colours in this rainbow", language.BritishEnglish),
			Entry(nil, "american", "There are 7 colors in this rainbow", language.AmericanEnglish),
		)
	})

	Context("go-i18n", func() {
		When("using map of any", func() {
			It("🧪 should: translate", func() {
				violationMsg := &i18n.Message{
					ID:    "ov-failed-out-of-range",
					Other: "({{.Flag}}): option validation failed, '{{.Value}}', out of range: [{{.Lo}}]..[{{.Hi}}]",
				}

				localised := assistant.Localiser.MustLocalize(&i18n.LocalizeConfig{
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

				localised := assistant.Localiser.MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: violationMsg,
					TemplateData:   OptValidationOutOfRangeTemplData{"Strike", 999, 1, 99},
				})
				expected := "(Strike): option validation failed, '999', out of range: [1]..[99]"
				Expect(localised).To(Equal(expected))
			})
		})

		When("using translation", func() {
			It("🧪 should: translate", func() {
				localised := assistant.GetOutOfRangeErrorMessage("Strike", 999, 1, 99)
				GinkgoWriter.Printf("===> localised: '%v'\n", localised)

				expected := "(Strike): option validation failed, '999', out of range: [1]..[99]"
				Expect(localised).To(Equal(expected))
			})
		})
	})
})
