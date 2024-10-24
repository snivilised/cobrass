package locale_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok

	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/cobrass/src/internal/lab"
	"github.com/snivilised/li18ngo"
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

		from li18ngo.LoadFrom
	)

	BeforeAll(func() {
		repo = lab.Repo("../..")
		l10nPath = lab.Path(repo, "Test/data/l10n")

		from = li18ngo.LoadFrom{
			Path: l10nPath,
			Sources: li18ngo.TranslationFiles{
				locale.CobrassSourceID: li18ngo.TranslationSource{Name: "test"},
			},
		}

		if err := li18ngo.Use(func(o *li18ngo.UseOptions) {
			o.From = from
		}); err != nil {
			Fail(err.Error())
		}
	})

	// these tests may not be required, because they may be able to be generated
	DescribeTable("Native Errors",
		func(entry validationEntry) {
			err := lab.CallE(entry.Fn, entry.Args)
			GinkgoWriter.Printf("VALIDATION-ERROR-RESULT: %v", err)
			fmt.Printf("⚠️ VALIDATION-ERROR-RESULT(%v): '%v'\n", entry.Name, err)
			Expect(err).Error().NotTo(BeNil())
		},
		func(entry validationEntry) string {
			return fmt.Sprintf("🧪 --> 🔥 native error function: '%v'", entry.Name)
		},

		Entry(nil, validationEntry{
			Name: "NewWithinOptValidationError",
			Fn:   locale.NewWithinOptValidationError,
			Args: []any{"foo-flag", 1, 10, 20},
			Verify: func(err error) bool {
				if e, ok := err.(locale.WithinOptValidationBehaviourQuery); ok {
					return e.IsOutOfRange()
				}
				return false
			},
		}),

		Entry(nil, validationEntry{
			Name: "NewNotWithinOptValidationError",
			Fn:   locale.NewNotWithinOptValidationError,
			Args: []any{"foo-flag", 5, 10, 20},
			Verify: func(_ error) bool {
				return true
			},
		}),
	)

	Context("NewNotContainsOptValidationError", func() {
		It("should: create error", func() {
			err := locale.NewNotContainsOptValidationError("foo-flag", int(1), []int{2, 4, 6, 8})
			GinkgoWriter.Printf("💥💥💥 ===> ERROR: '%v'", err)
		})
	})
})
