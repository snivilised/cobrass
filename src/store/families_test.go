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
		func(commandLine []string) {
			ps := assistant.NewParamSet[store.FilterParameterSet](rootCommand)
			ps.Native.BindAll(ps)

			execute(commandLine)
		},
		func(args []string) string {
			return shouldMessage
		},
		Entry(
			nil, []string{"--files-rx", "^foo", "--folders-gb", "bar*"},
		),
		Entry(
			nil, []string{"-X", "^foo", "-z", "bar*"},
		),
		Entry(
			nil, []string{"--files-gb", "foo*", "--folders-rx", "^bar"},
		),
		Entry(
			nil, []string{"-G", "foo*", "-y", "^bar"},
		),
	)

	DescribeTable("worker pool family",
		func(commandLine []string) {
			ps := assistant.NewParamSet[store.WorkerPoolParameterSet](rootCommand)
			ps.Native.BindAll(ps)

			execute(commandLine)
		},
		func(args []string) string {
			return shouldMessage
		},
		Entry(
			nil, []string{"--cpu"},
		),
		Entry(
			nil, []string{"-C"},
		),
		Entry(
			nil, []string{"--now", "4"},
		),
		Entry(
			nil, []string{"-N", "4"},
		),
	)

	DescribeTable("profile family",
		func(commandLine []string) {
			ps := assistant.NewParamSet[store.ProfileParameterSet](rootCommand)
			ps.Native.BindAll(ps)

			execute(commandLine)
		},
		func(args []string) string {
			return shouldMessage
		},
		Entry(
			nil, []string{"--profile", "foo"},
		),
		Entry(
			nil, []string{"-P", "foo"},
		),
	)

	DescribeTable("profile family",
		func(commandLine []string) {
			ps := assistant.NewParamSet[store.PreviewParameterSet](rootCommand)
			ps.Native.BindAll(ps)

			execute(commandLine)
		},
		func(args []string) string {
			return shouldMessage
		},
		Entry(
			nil, []string{"--dry-run"},
		),
		Entry(
			nil, []string{"-D"},
		),
	)

	When("usage requested", func() {
		It("should: 🧪 show help text", func() {
			filtersPS := assistant.NewParamSet[store.FilterParameterSet](rootCommand)
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
})
