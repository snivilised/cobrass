package assistant_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/cobrass/src/clif"
	"github.com/snivilised/cobrass/src/internal/helpers"
	"github.com/snivilised/cobrass/src/store"
	xi18n "github.com/snivilised/extendio/i18n"
)

type baseTE struct {
	given  string
	should string
}

type helperTS struct {
	baseTE
	args     []string
	knownBy  cobrass.KnownByCollection
	expected clif.ThirdPartyCommandLine
}

var _ = Describe("GetThirdPartyCL", Ordered, func() {
	var (
		repo     string
		l10nPath string

		from        xi18n.LoadFrom
		rootCommand *cobra.Command

		paramSet *assistant.ParamSet[store.ProfileParameterSet]
		knownBy  cobrass.KnownByCollection
	)

	BeforeAll(func() {
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "Test/data/l10n")

		from = xi18n.LoadFrom{
			Path: l10nPath,
			Sources: xi18n.TranslationFiles{
				locale.CobrassSourceID: xi18n.TranslationSource{Name: "test"},
			},
		}

		if err := xi18n.Use(func(o *xi18n.UseOptions) {
			o.From = from
		}); err != nil {
			Fail(err.Error())
		}

		knownBy = cobrass.KnownByCollection{
			"profile": "P",
			"scheme":  "S",
		}
	})

	DescribeTable("GetThirdPartyCL",
		func(entry *helperTS) {
			entry.knownBy = knownBy
			args := entry.args

			rootCommand = &cobra.Command{
				Use:   "scorch",
				Short: "scotch",
				Long:  "scorch is a fake test command which contains filtering capabilities",
				RunE: func(cmd *cobra.Command, _ []string) error {
					flagSet := cmd.Flags()
					actual := assistant.GetThirdPartyCL(flagSet, entry.knownBy)

					Expect(actual).To(ContainElements(entry.expected))

					return nil
				},
			}
			paramSet = assistant.NewParamSet[store.ProfileParameterSet](rootCommand)
			paramSet.Native.BindAll(paramSet, rootCommand.PersistentFlags())

			_, err := helpers.ExecuteCommand(
				rootCommand, args...,
			)

			Expect(err).Error().To(BeNil())
		},
		func(entry *helperTS) string {
			return fmt.Sprintf("given: '%v', 🧪 should: '%v'", entry.given, entry.should)
		},

		Entry(nil, &helperTS{
			baseTE: baseTE{
				given:  "flags in long form",
				should: "return a command line containing long form flags",
			},
			args:     []string{"/usr/fuse/home/music", "--profile", "blur", "--scheme", "ectoplasm"},
			expected: clif.ThirdPartyCommandLine{"--profile", "blur", "--scheme", "ectoplasm"},
		}),

		Entry(nil, &helperTS{
			baseTE: baseTE{
				given:  "flags in short form",
				should: "return a command line containing long form flags",
			},
			args:     []string{"/usr/fuse/home/music", "-P", "blur", "-S", "ectoplasm"},
			expected: clif.ThirdPartyCommandLine{"--profile", "blur", "--scheme", "ectoplasm"},
		}),
	)
})
