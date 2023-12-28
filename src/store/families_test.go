package store_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	"github.com/snivilised/cobrass/src/internal/helpers"
	"github.com/snivilised/cobrass/src/store"
	xi18n "github.com/snivilised/extendio/i18n"
)

func reason(binder string, err error) string {
	return fmt.Sprintf("🔥 expected '%v' error to be nil, but was '%v'\n",
		binder, err,
	)
}

const (
	shouldMessage = "🧪 should: bind all parameters without error"
)

// --files-gb(G)
// --files-rx(X)
// --folders-gb(Z)
// --folders-rx(y)

type familyTE struct {
	familyType  string
	persistent  bool
	commandLine []string
}

var _ = Describe("Families", Ordered, func() {
	var (
		repo     string
		l10nPath string

		from        xi18n.LoadFrom
		rootCommand *cobra.Command
		execute     func(args []string)
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

		execute = func(args []string) {
			_, err := helpers.ExecuteCommand(
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
			RunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}
	})

	DescribeTable("filter family",
		func(entry *familyTE) {
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
			}

			execute(entry.commandLine)
		},
		func(entry *familyTE) string {
			return shouldMessage
		},
		Entry(
			nil,
			&familyTE{
				familyType:  "files",
				persistent:  true,
				commandLine: []string{"--files-rx", "^foo"},
			},
		),
		Entry(
			nil,
			&familyTE{
				familyType:  "files",
				commandLine: []string{"-X", "^foo"},
			},
		),
		//
		Entry(
			nil,
			&familyTE{
				familyType:  "folders",
				commandLine: []string{"--folders-gb", "bar*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				familyType:  "folders",
				persistent:  true,
				commandLine: []string{"-Z", "bar*"},
			},
		),
		//
		Entry(
			nil,
			&familyTE{
				familyType:  "poly",
				commandLine: []string{"--files-rx", "^foo", "--folders-gb", "bar*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				familyType:  "poly",
				commandLine: []string{"-X", "^foo", "-Z", "bar*"},
			},
		),
		Entry(
			nil,
			&familyTE{
				familyType:  "poly",
				persistent:  true,
				commandLine: []string{"--files-gb", "foo*", "--folders-rx", "^bar"},
			},
		),
		Entry(
			nil,
			&familyTE{
				familyType:  "poly",
				persistent:  true,
				commandLine: []string{"-G", "foo*", "-Y", "^bar"},
			},
		),
		//

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
			return shouldMessage
		},
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"--cpu"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
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
			return shouldMessage
		},
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"--profile", "foo"},
			},
		),
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"-P", "foo"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"--scheme", "foo"},
			},
		),
		Entry(
			nil,
			&familyTE{
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
			return shouldMessage
		},
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"--dry-run"},
				persistent:  true,
			},
		),
		Entry(
			nil,
			&familyTE{
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
			return shouldMessage
		},
		Entry(
			nil,
			&familyTE{
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
			_, err := helpers.ExecuteCommand(
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
			return shouldMessage
		},
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"--depth", "3"},
			},
		),
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"--skim"},
			},
		),
		Entry(
			nil,
			&familyTE{
				commandLine: []string{"-K"},
				persistent:  true,
			},
		),
	)
})
