package store_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/cobrass/src/internal/lab"
	"github.com/snivilised/cobrass/src/store"
	"github.com/snivilised/li18ngo"
)

func reason(binder string, err error) string {
	return fmt.Sprintf("🔥 expected '%v' error to be nil, but was '%v'\n",
		binder, err,
	)
}

// --files(F)
// --files-gb(G)
// --files-rx(X)
// --folders-gb(Z)
// --folders-rx(Y)

type familyTE struct {
	given       string
	familyType  string
	persistent  bool
	commandLine []string
}

var _ = Describe("Families", Ordered, func() {
	var (
		repo     string
		l10nPath string

		from        li18ngo.LoadFrom
		rootCommand *cobra.Command
		execute     func(args []string)
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

		execute = func(args []string) {
			_, err := lab.ExecuteCommand(
				rootCommand, args...,
			)
			Expect(err).Error().To(BeNil(), reason("BindAll", err))
		}
	})

	BeforeEach(func() {
		rootCommand = &cobra.Command{
			Use:   "scorch",
			Short: "scotch",
			Long:  "scorch is a fake test command which contains filtering capabilities",
			RunE: func(_ *cobra.Command, _ []string) error {
				return nil
			},
		}
	})

	DescribeTable("filter family",
		func(entry *familyTE) {
			defer func() {
				r := recover()

				if r != nil {
					Fail(fmt.Sprintf("!!! 💥 %v\n", r))
				}
			}()

			switch entry.familyType {
			case "poly":
				{
					ps := assistant.NewParamSet[store.PolyFilterParameterSet](rootCommand)
					if entry.persistent {
						ps.Native.BindAll(ps, rootCommand.PersistentFlags())
					} else {
						ps.Native.BindAll(ps)
					}
				}

			case "files":
				{
					ps := assistant.NewParamSet[store.FilesFilterParameterSet](rootCommand)
					if entry.persistent {
						ps.Native.BindAll(ps, rootCommand.PersistentFlags())
					} else {
						ps.Native.BindAll(ps)
					}
				}

			case "folders":
				{
					ps := assistant.NewParamSet[store.FoldersFilterParameterSet](rootCommand)
					if entry.persistent {
						ps.Native.BindAll(ps, rootCommand.PersistentFlags())
					} else {
						ps.Native.BindAll(ps)
					}
				}

			case "alloy":
				{
					ps := assistant.NewParamSet[store.AlloyFilterParameterSet](rootCommand)
					if entry.persistent {
						ps.Native.BindAll(ps, rootCommand.PersistentFlags())
					} else {
						ps.Native.BindAll(ps)
					}
				}
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--files",
				familyType:  "files",
				persistent:  true,
				commandLine: []string{"--files", "foo*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-f",
				familyType:  "files",
				commandLine: []string{"-X", "foo*"},
			},
		),
		//
		Entry(
			nil,
			&familyTE{
				given:       "--files-rx",
				familyType:  "files",
				persistent:  true,
				commandLine: []string{"--files-rx", "^foo"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--X",
				familyType:  "files",
				commandLine: []string{"-X", "^foo"},
			},
		),
		//
		Entry(
			nil,
			&familyTE{
				given:       "--folders-gb",
				familyType:  "folders",
				commandLine: []string{"--folders-gb", "bar*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-Z",
				familyType:  "folders",
				persistent:  true,
				commandLine: []string{"-Z", "bar*"},
			},
		),
		//
		Entry(
			nil,
			&familyTE{
				given:       "--files --folders-gb",
				familyType:  "poly",
				commandLine: []string{"--files", "foo*", "--folders-gb", "bar*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-F -Z",
				familyType:  "poly",
				commandLine: []string{"-F", "foo*", "-Z", "bar*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--files --folders-rx",
				familyType:  "poly",
				persistent:  true,
				commandLine: []string{"--files", "foo*", "--folders-rx", "^bar"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-F -Y",
				familyType:  "poly",
				persistent:  true,
				commandLine: []string{"-F", "foo*", "-Y", "^bar"},
			},
		),
		//
		Entry(
			nil,
			&familyTE{
				given:       "--files",
				familyType:  "alloy",
				persistent:  true,
				commandLine: []string{"--files", "foo*|jpg,txt"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-F",
				familyType:  "alloy",
				commandLine: []string{"-F", "foo*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--folders-gb",
				familyType:  "alloy",
				persistent:  true,
				commandLine: []string{"--folders-gb", "foo*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-Z",
				familyType:  "alloy",
				commandLine: []string{"-Z", "foo*"},
			},
		),
	)

	DescribeTable("worker pool family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.WorkerPoolParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--cpu",
				commandLine: []string{"--cpu"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--now",
				commandLine: []string{"--now", "4"},
				persistent:  true,
			},
		),
	)

	DescribeTable("profile family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.ProfileParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--profile",
				commandLine: []string{"--profile", "foo"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-P",
				commandLine: []string{"-P", "foo"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--scheme",
				commandLine: []string{"--scheme", "foo"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-S",
				commandLine: []string{"-S", "foo"},
				persistent:  true,
			},
		),
	)

	DescribeTable("preview family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.PreviewParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--dry-run",
				commandLine: []string{"--dry-run"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-D",
				commandLine: []string{"-D"},
			},
		),
	)

	DescribeTable("i18n family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.I18nParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--language",
				commandLine: []string{"--language", "en-GB"},
				persistent:  true,
			},
		),
	)

	When("usage requested", func() {
		It("should: 🧪 show help text", func() {
			filtersPS := assistant.NewParamSet[store.PolyFilterParameterSet](rootCommand)
			filtersPS.Native.BindAll(filtersPS)
			//
			poolPS := assistant.NewParamSet[store.WorkerPoolParameterSet](rootCommand)
			poolPS.Native.BindAll(poolPS)
			//
			profilePS := assistant.NewParamSet[store.ProfileParameterSet](rootCommand)
			profilePS.Native.BindAll(profilePS)
			//
			previewPS := assistant.NewParamSet[store.PreviewParameterSet](rootCommand)
			previewPS.Native.BindAll(previewPS)
			//
			commandLine := []string{"scorch", "--help"}
			_, err := lab.ExecuteCommand(
				rootCommand, commandLine...,
			)
			Expect(err).Error().To(BeNil(), reason("BindAll", err))
		})
	})

	DescribeTable("cascade family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.CascadeParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--depth",
				commandLine: []string{"--depth", "3"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--no-recurse",
				commandLine: []string{"--no-recurse"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "-N",
				commandLine: []string{"-N"},
				persistent:  true,
			},
		),
	)

	DescribeTable("sampling family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.SamplingParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--sample",
				commandLine: []string{"--sample"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--no-files",
				commandLine: []string{"--no-files", "4"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--no-folders",
				commandLine: []string{"--no-folders", "5"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--last",
				commandLine: []string{"--last"},
			},
		),
	)

	DescribeTable("textual-interaction family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.TextualInteractionParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--no-tui",
				commandLine: []string{"--no-tui"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--no-tui",
				commandLine: []string{"--no-tui"},
				persistent:  true,
			},
		),
	)

	DescribeTable("cli-interaction family",
		func(entry *familyTE) {
			ps := assistant.NewParamSet[store.CliInteractionParameterSet](rootCommand)
			if entry.persistent {
				ps.Native.BindAll(ps, rootCommand.PersistentFlags())
			} else {
				ps.Native.BindAll(ps)
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return fmt.Sprintf("🧪 given: '%v', should: bind all parameters without error", entry.given)
		},
		Entry(
			nil,
			&familyTE{
				given:       "--tui",
				commandLine: []string{"--tui"},
			},
		),
		Entry(
			nil,
			&familyTE{
				given:       "--tui",
				commandLine: []string{"--tui"},
				persistent:  true,
			},
		),
	)
})
