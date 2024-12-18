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
	"github.com/snivilised/cobrass/src/internal/lab"
	"github.com/snivilised/cobrass/src/store"
	"github.com/snivilised/li18ngo"
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

		from        li18ngo.LoadFrom
		rootCommand *cobra.Command

		paramSet *assistant.ParamSet[store.ProfileParameterSet]
		knownBy  cobrass.KnownByCollection
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

			_, err := lab.ExecuteCommand(
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
