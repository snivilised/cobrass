package adapters_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/cobrass/src/adapters"
	"github.com/spf13/cobra"
)

var _ = Describe("ParamSetBinderHelpers", func() {
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *adapters.ParamSet[WidgetParameterSet]
	// var outputFormatEnumInfo *adapters.EnumInfo[OutputFormatEnum]

	Context("Comparables", func() {
		BeforeEach(func() {
			rootCommand = &cobra.Command{
				Use:   "flick",
				Short: "A brief description of your application",
				Long:  "A long description of the root flick command",
			}

			widgetCommand = &cobra.Command{
				Version: "1.0.1",
				Use:     "widget",
				Short:   "Create widget",
				Long:    "Index file system at root: '/'",
				Args:    cobra.ExactArgs(1),
				RunE: func(command *cobra.Command, args []string) error {
					paramSet.Native.Directory = args[0]
					return nil
				},
			}
			rootCommand.AddCommand(widgetCommand)
			paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)

		})

		DescribeTable("[BindValidatedInt]GreatherThan",
			func(given, should string, threshold, value int, expectNil bool) {
				validator := paramSet.BindValidatedIntGreaterThan(
					adapters.NewFlagInfo("offset", "o", 0),
					&paramSet.Native.Offset, threshold,
				)
				paramSet.Native.Offset = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, threshold, value int, expectNil bool) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", 10, 9, false),
			Entry(nil, "value is equal to threshold", "return error", 10, 10, false),
			Entry(nil, "value is equal to threshold", "NOT return error", 10, 11, true),
		)

		DescribeTable("BindValidatedContainsString",
			func(given, should string, value string, expectNil bool, collection []string, dummy string) {
				validator := paramSet.BindValidatedContainsString(
					adapters.NewFlagInfo("offset", "o", ""),
					&paramSet.Native.Pattern, collection,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, collection []string, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", "a", true, []string{"a", "b", "c"}, "null"),

			Entry(nil, "collection does not contain member", "return error", "x", false, []string{"a", "b", "c"}, "null"),
		)

		// ----> auto generated(Build-BinderHelperTests/gen-help-t)
		DescribeTable("BindValidatedStringWithin",
			func(given, should string, value string, expectNil bool, low, high string) {
				validator := paramSet.BindValidatedStringWithin(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, low, high,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, low, high string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", "b", false, "c", "e"),
			Entry(nil, "value is equal to low end of range", "return error", "c", true, "c", "e"),
			Entry(nil, "value is inside range", "return error", "d", true, "c", "e"),
			Entry(nil, "value is equal to high end of range", "return error", "e", true, "c", "e"),
			Entry(nil, "value is above range", "NOT return error", "f", false, "c", "e"),
		)

		DescribeTable("BindValidatedContainsString",
			func(given, should string, value string, expectNil bool, collection []string, dummy string) {
				validator := paramSet.BindValidatedContainsString(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, collection,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, collection []string, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", "a", true, []string{"a", "b", "c"}, "null"),
			Entry(nil, "collection does not contain member", "return error", "x", false, []string{"a", "b", "c"}, "null"),
		)

		DescribeTable("BindValidatedStringIsMatch",
			func(given, should string, value string, expectNil bool, pattern, dummy string) {
				validator := paramSet.BindValidatedStringIsMatch(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, pattern,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value matches pattern", "return error", "18-10-1997", true, "\\d{2}-\\d{2}-\\d{4}", "null"),
			Entry(nil, "value does not match pattern", "return error", "foo-bar", false, "\\d{2}-\\d{2}-\\d{4}", "null"),
		)

		DescribeTable("BindValidatedStringIsNotMatch",
			func(given, should string, value string, expectNil bool, pattern, dummy string) {
				validator := paramSet.BindValidatedStringIsNotMatch(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, pattern,
				)
				paramSet.Native.Pattern = value

				if !expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value matches pattern", "return error", "18-10-1997", true, "\\d{2}-\\d{2}-\\d{4}", "null"),
			Entry(nil, "value does not match pattern", "return error", "foo-bar", false, "\\d{2}-\\d{2}-\\d{4}", "null"),
		)

		DescribeTable("BindValidatedStringGreaterThan",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringGreaterThan(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, threshold, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", false, "c", ""),
			Entry(nil, "value is equal to threshold", "return error", "c", false, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", true, "c", ""),
		)

		DescribeTable("BindValidatedStringAtLeast",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringAtLeast(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, threshold, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", false, "c", ""),
			Entry(nil, "value is equal to threshold", "return error", "c", true, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", true, "c", ""),
		)

		DescribeTable("BindValidatedStringLessThan",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringLessThan(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, threshold, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", true, "c", ""),
			Entry(nil, "value is equal to threshold", "return error", "c", false, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", false, "c", ""),
		)

		DescribeTable("BindValidatedStringAtMost",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringAtMost(
					adapters.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value

				if expectNil {
					Expect(validator.Validate()).Error().To(BeNil())
				} else {
					Expect(validator.Validate()).Error().ToNot(BeNil())
				}

			},
			func(given, should string, value string, expectNil bool, threshold, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", true, "c", ""),
			Entry(nil, "value is equal to threshold", "return error", "c", true, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", false, "c", ""),
		)

		// <---- auto generated
	})
})
