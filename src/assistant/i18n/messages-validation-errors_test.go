package i18n_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/src/assistant/i18n"
	"github.com/snivilised/cobrass/src/internal/helpers"
	xi18n "github.com/snivilised/extendio/i18n"
)

type validationEntry struct {
	Name   string
	Args   []any
	Fn     any
	Verify func(err error) bool
}

var _ = Describe("MessagesValidationErrors", Ordered, func() {

	var (
		repo     string
		l10nPath string

		from xi18n.LoadFrom
	)

	BeforeAll(func() {
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "Test/data/l10n")

		from = xi18n.LoadFrom{
			Path: l10nPath,
			Sources: xi18n.TranslationFiles{
				i18n.CobrassSourceID: xi18n.TranslationSource{Name: "test"},
			},
		}

		if err := xi18n.Use(func(o *xi18n.UseOptions) {
			o.From = from
		}); err != nil {
			Fail(err.Error())
		}
	})

	// these tests may not be required, because they may be able to be generated
	DescribeTable("Native Errors",
		func(entry validationEntry) {
			err := helpers.CallE(entry.Fn, entry.Args)
			GinkgoWriter.Printf("VALIDATION-ERROR-RESULT: %v", err)
			fmt.Printf("⚠️ VALIDATION-ERROR-RESULT(%v): '%v'\n", entry.Name, err)
			Expect(err).Error().NotTo(BeNil())
		},
		func(entry validationEntry) string {
			return fmt.Sprintf("🧪 --> 🔥 native error function: '%v'", entry.Name)
		},

		Entry(nil, validationEntry{
			Name: "NewWithinOptValidationError",
			Fn:   i18n.NewWithinOptValidationError,
			Args: []any{"foo-flag", 1, 10, 20},
			Verify: func(err error) bool {
				if e, ok := err.(i18n.WithinOptValidationBehaviourQuery); ok {
					return e.IsOutOfRange()
				}
				return false
			},
		}),

		Entry(nil, validationEntry{
			Name: "NewNotWithinOptValidationError",
			Fn:   i18n.NewNotWithinOptValidationError,
			Args: []any{"foo-flag", 5, 10, 20},
			Verify: func(err error) bool {
				return true
			},
		}),
	)

	Context("NewNotContainsOptValidationError", func() {
		It("should: create error", func() {
			err := i18n.NewNotContainsOptValidationError("foo-flag", int(1), []int{2, 4, 6, 8})
			GinkgoWriter.Printf("💥💥💥 ===> ERROR: '%v'", err)
		})
	})
})
