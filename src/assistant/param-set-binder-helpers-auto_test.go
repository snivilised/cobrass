package assistant_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/assistant"
)

type validatorDecorator struct {
	Decorated assistant.OptionValidator
}

func (v *validatorDecorator) Validate() error {
	flag := v.Decorated.GetFlag()
	flag.Changed = true

	return v.Decorated.Validate()
}

var _ = Describe("ParamSetBinderHelpers", func() {
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *assistant.ParamSet[WidgetParameterSet]
	var outputFormatEnumInfo *assistant.EnumInfo[OutputFormatEnum]
	var outputFormatEnum assistant.EnumValue[OutputFormatEnum]

	Context("Comparables", func() {
		BeforeEach(func() {
			outputFormatEnumInfo = assistant.NewEnumInfo(AcceptableOutputFormats)
			outputFormatEnum = outputFormatEnumInfo.NewValue()

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
			paramSet = assistant.NewParamSet[WidgetParameterSet](widgetCommand)
		})

		// ----> auto generated(Build-BinderHelperTests/gen-help-t)

		DescribeTable("BindValidatedDurationWithin",
			func(given, should string, value time.Duration, expectNil bool, low, high time.Duration) {
				validator := paramSet.BindValidatedDurationWithin(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Latency = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, low, high time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", duration("2s"), false, duration("3s"), duration("5s")),
			Entry(nil, "value is equal to low end of range", "return error", duration("3s"), true, duration("3s"), duration("5s")),
			Entry(nil, "value is inside range", "return error", duration("4s"), true, duration("3s"), duration("5s")),
			Entry(nil, "value is equal to high end of range", "return error", duration("5s"), true, duration("3s"), duration("5s")),
			Entry(nil, "value is above range", "NOT return error", duration("6s"), false, duration("3s"), duration("5s")),
		)

		DescribeTable("BindValidatedDurationNotWithin",
			func(given, should string, value time.Duration, expectNil bool, low, high time.Duration) {
				validator := paramSet.BindValidatedDurationNotWithin(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Latency = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, low, high time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", duration("2s"), false, duration("3s"), duration("5s")),
			Entry(nil, "value is equal to low end of range", "return error", duration("3s"), true, duration("3s"), duration("5s")),
			Entry(nil, "value is inside range", "return error", duration("4s"), true, duration("3s"), duration("5s")),
			Entry(nil, "value is equal to high end of range", "return error", duration("5s"), true, duration("3s"), duration("5s")),
			Entry(nil, "value is above range", "NOT return error", duration("6s"), false, duration("3s"), duration("5s")),
		)

		DescribeTable("BindValidatedContainsDuration",
			func(given, should string, value time.Duration, expectNil bool, collection []time.Duration, dummy time.Duration) {
				validator := paramSet.BindValidatedContainsDuration(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, collection,
				)
				paramSet.Native.Latency = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, collection []time.Duration, dummy time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", duration("1s"), true, []time.Duration{duration("1s"), duration("2s"), duration("3s")}, duration("0s")),
			Entry(nil, "collection does not contain member", "return error", duration("99s"), false, []time.Duration{duration("1s"), duration("2s"), duration("3s")}, duration("0s")),
		)

		DescribeTable("BindValidatedNotContainsDuration",
			func(given, should string, value time.Duration, expectNil bool, collection []time.Duration, dummy time.Duration) {
				validator := paramSet.BindValidatedNotContainsDuration(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, collection,
				)
				paramSet.Native.Latency = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, collection []time.Duration, dummy time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", duration("1s"), true, []time.Duration{duration("1s"), duration("2s"), duration("3s")}, duration("0s")),
			Entry(nil, "collection does not contain member", "return error", duration("99s"), false, []time.Duration{duration("1s"), duration("2s"), duration("3s")}, duration("0s")),
		)

		DescribeTable("BindValidatedDurationGreaterThan",
			func(given, should string, value time.Duration, expectNil bool, threshold, dummy time.Duration) {
				validator := paramSet.BindValidatedDurationGreaterThan(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, threshold,
				)
				paramSet.Native.Latency = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, pattern, dummy time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", duration("2s"), false, duration("3s"), duration("0s")),
			Entry(nil, "value is equal threshold", "return error", duration("3s"), false, duration("3s"), duration("0s")),
			Entry(nil, "value is above threshold", "NOT return error", duration("4s"), true, duration("3s"), duration("0s")),
		)

		DescribeTable("BindValidatedDurationAtLeast",
			func(given, should string, value time.Duration, expectNil bool, threshold, dummy time.Duration) {
				validator := paramSet.BindValidatedDurationAtLeast(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, threshold,
				)
				paramSet.Native.Latency = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, pattern, dummy time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", duration("2s"), false, duration("3s"), duration("0s")),
			Entry(nil, "value is equal threshold", "return error", duration("3s"), true, duration("3s"), duration("0s")),
			Entry(nil, "value is above threshold", "NOT return error", duration("4s"), true, duration("3s"), duration("0s")),
		)

		DescribeTable("BindValidatedDurationLessThan",
			func(given, should string, value time.Duration, expectNil bool, threshold, dummy time.Duration) {
				validator := paramSet.BindValidatedDurationLessThan(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, threshold,
				)
				paramSet.Native.Latency = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, pattern, dummy time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", duration("2s"), true, duration("3s"), duration("0s")),
			Entry(nil, "value is equal threshold", "return error", duration("3s"), false, duration("3s"), duration("0s")),
			Entry(nil, "value is above threshold", "NOT return error", duration("4s"), false, duration("3s"), duration("0s")),
		)

		DescribeTable("BindValidatedDurationAtMost",
			func(given, should string, value time.Duration, expectNil bool, threshold, dummy time.Duration) {
				validator := paramSet.BindValidatedDurationAtMost(
					assistant.NewFlagInfo("latency", "l", duration("0ms")),
					&paramSet.Native.Latency, threshold,
				)
				paramSet.Native.Latency = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value time.Duration, expectNil bool, pattern, dummy time.Duration) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", duration("2s"), true, duration("3s"), duration("0s")),
			Entry(nil, "value is equal threshold", "return error", duration("3s"), true, duration("3s"), duration("0s")),
			Entry(nil, "value is above threshold", "NOT return error", duration("4s"), false, duration("3s"), duration("0s")),
		)

		DescribeTable("BindValidatedContainsEnum",
			func(given, should string, value string, expectNil bool, collection []string, dummy string) {
				validator := paramSet.BindValidatedContainsEnum(
					assistant.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source, collection,
				)
				outputFormatEnum.Source = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, collection []string, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", "xml", true, []string{"json", "text", "xml"}, "null"),
			Entry(nil, "collection does not contain member", "return error", "scr", false, []string{"json", "text", "xml"}, "null"),
		)

		DescribeTable("BindValidatedNotContainsEnum",
			func(given, should string, value string, expectNil bool, collection []string, dummy string) {
				validator := paramSet.BindValidatedNotContainsEnum(
					assistant.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source, collection,
				)
				outputFormatEnum.Source = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, collection []string, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", "xml", true, []string{"json", "text", "xml"}, "null"),
			Entry(nil, "collection does not contain member", "return error", "scr", false, []string{"json", "text", "xml"}, "null"),
		)

		DescribeTable("BindValidatedFloat32Within",
			func(given, should string, value float32, expectNil bool, low, high float32) {
				validator := paramSet.BindValidatedFloat32Within(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Gradientf32 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, low, high float32) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", float32(2), false, float32(3), float32(5)),
			Entry(nil, "value is equal to low end of range", "return error", float32(3), true, float32(3), float32(5)),
			Entry(nil, "value is inside range", "return error", float32(4), true, float32(3), float32(5)),
			Entry(nil, "value is equal to high end of range", "return error", float32(5), true, float32(3), float32(5)),
			Entry(nil, "value is above range", "NOT return error", float32(6), false, float32(3), float32(5)),
		)

		DescribeTable("BindValidatedFloat32NotWithin",
			func(given, should string, value float32, expectNil bool, low, high float32) {
				validator := paramSet.BindValidatedFloat32NotWithin(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Gradientf32 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, low, high float32) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", float32(2), false, float32(3), float32(5)),
			Entry(nil, "value is equal to low end of range", "return error", float32(3), true, float32(3), float32(5)),
			Entry(nil, "value is inside range", "return error", float32(4), true, float32(3), float32(5)),
			Entry(nil, "value is equal to high end of range", "return error", float32(5), true, float32(3), float32(5)),
			Entry(nil, "value is above range", "NOT return error", float32(6), false, float32(3), float32(5)),
		)

		DescribeTable("BindValidatedContainsFloat32",
			func(given, should string, value float32, expectNil bool, collection []float32, dummy float32) {
				validator := paramSet.BindValidatedContainsFloat32(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, collection,
				)
				paramSet.Native.Gradientf32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, collection []float32, dummy float32) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", float32(1), true, []float32{1, 2, 3}, float32(0)),
			Entry(nil, "collection does not contain member", "return error", float32(99), false, []float32{1, 2, 3}, float32(0)),
		)

		DescribeTable("BindValidatedNotContainsFloat32",
			func(given, should string, value float32, expectNil bool, collection []float32, dummy float32) {
				validator := paramSet.BindValidatedNotContainsFloat32(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, collection,
				)
				paramSet.Native.Gradientf32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, collection []float32, dummy float32) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", float32(1), true, []float32{1, 2, 3}, float32(0)),
			Entry(nil, "collection does not contain member", "return error", float32(99), false, []float32{1, 2, 3}, float32(0)),
		)

		DescribeTable("BindValidatedFloat32GreaterThan",
			func(given, should string, value float32, expectNil bool, threshold, dummy float32) {
				validator := paramSet.BindValidatedFloat32GreaterThan(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, threshold,
				)
				paramSet.Native.Gradientf32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, pattern, dummy float32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float32(2), false, float32(3), float32(0)),
			Entry(nil, "value is equal threshold", "return error", float32(3), false, float32(3), float32(0)),
			Entry(nil, "value is above threshold", "NOT return error", float32(4), true, float32(3), float32(0)),
		)

		DescribeTable("BindValidatedFloat32AtLeast",
			func(given, should string, value float32, expectNil bool, threshold, dummy float32) {
				validator := paramSet.BindValidatedFloat32AtLeast(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, threshold,
				)
				paramSet.Native.Gradientf32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, pattern, dummy float32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float32(2), false, float32(3), float32(0)),
			Entry(nil, "value is equal threshold", "return error", float32(3), true, float32(3), float32(0)),
			Entry(nil, "value is above threshold", "NOT return error", float32(4), true, float32(3), float32(0)),
		)

		DescribeTable("BindValidatedFloat32LessThan",
			func(given, should string, value float32, expectNil bool, threshold, dummy float32) {
				validator := paramSet.BindValidatedFloat32LessThan(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, threshold,
				)
				paramSet.Native.Gradientf32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, pattern, dummy float32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float32(2), true, float32(3), float32(0)),
			Entry(nil, "value is equal threshold", "return error", float32(3), false, float32(3), float32(0)),
			Entry(nil, "value is above threshold", "NOT return error", float32(4), false, float32(3), float32(0)),
		)

		DescribeTable("BindValidatedFloat32AtMost",
			func(given, should string, value float32, expectNil bool, threshold, dummy float32) {
				validator := paramSet.BindValidatedFloat32AtMost(
					assistant.NewFlagInfo("gradientf32", "t", float32(0)),
					&paramSet.Native.Gradientf32, threshold,
				)
				paramSet.Native.Gradientf32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float32, expectNil bool, pattern, dummy float32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float32(2), true, float32(3), float32(0)),
			Entry(nil, "value is equal threshold", "return error", float32(3), true, float32(3), float32(0)),
			Entry(nil, "value is above threshold", "NOT return error", float32(4), false, float32(3), float32(0)),
		)

		DescribeTable("BindValidatedFloat64Within",
			func(given, should string, value float64, expectNil bool, low, high float64) {
				validator := paramSet.BindValidatedFloat64Within(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Gradientf64 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, low, high float64) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", float64(2), false, float64(3), float64(5)),
			Entry(nil, "value is equal to low end of range", "return error", float64(3), true, float64(3), float64(5)),
			Entry(nil, "value is inside range", "return error", float64(4), true, float64(3), float64(5)),
			Entry(nil, "value is equal to high end of range", "return error", float64(5), true, float64(3), float64(5)),
			Entry(nil, "value is above range", "NOT return error", float64(6), false, float64(3), float64(5)),
		)

		DescribeTable("BindValidatedFloat64NotWithin",
			func(given, should string, value float64, expectNil bool, low, high float64) {
				validator := paramSet.BindValidatedFloat64NotWithin(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Gradientf64 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, low, high float64) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", float64(2), false, float64(3), float64(5)),
			Entry(nil, "value is equal to low end of range", "return error", float64(3), true, float64(3), float64(5)),
			Entry(nil, "value is inside range", "return error", float64(4), true, float64(3), float64(5)),
			Entry(nil, "value is equal to high end of range", "return error", float64(5), true, float64(3), float64(5)),
			Entry(nil, "value is above range", "NOT return error", float64(6), false, float64(3), float64(5)),
		)

		DescribeTable("BindValidatedContainsFloat64",
			func(given, should string, value float64, expectNil bool, collection []float64, dummy float64) {
				validator := paramSet.BindValidatedContainsFloat64(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, collection,
				)
				paramSet.Native.Gradientf64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, collection []float64, dummy float64) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", float64(1), true, []float64{1, 2, 3}, float64(0)),
			Entry(nil, "collection does not contain member", "return error", float64(99), false, []float64{1, 2, 3}, float64(0)),
		)

		DescribeTable("BindValidatedNotContainsFloat64",
			func(given, should string, value float64, expectNil bool, collection []float64, dummy float64) {
				validator := paramSet.BindValidatedNotContainsFloat64(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, collection,
				)
				paramSet.Native.Gradientf64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, collection []float64, dummy float64) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", float64(1), true, []float64{1, 2, 3}, float64(0)),
			Entry(nil, "collection does not contain member", "return error", float64(99), false, []float64{1, 2, 3}, float64(0)),
		)

		DescribeTable("BindValidatedFloat64GreaterThan",
			func(given, should string, value float64, expectNil bool, threshold, dummy float64) {
				validator := paramSet.BindValidatedFloat64GreaterThan(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, threshold,
				)
				paramSet.Native.Gradientf64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, pattern, dummy float64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float64(2), false, float64(3), float64(0)),
			Entry(nil, "value is equal threshold", "return error", float64(3), false, float64(3), float64(0)),
			Entry(nil, "value is above threshold", "NOT return error", float64(4), true, float64(3), float64(0)),
		)

		DescribeTable("BindValidatedFloat64AtLeast",
			func(given, should string, value float64, expectNil bool, threshold, dummy float64) {
				validator := paramSet.BindValidatedFloat64AtLeast(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, threshold,
				)
				paramSet.Native.Gradientf64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, pattern, dummy float64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float64(2), false, float64(3), float64(0)),
			Entry(nil, "value is equal threshold", "return error", float64(3), true, float64(3), float64(0)),
			Entry(nil, "value is above threshold", "NOT return error", float64(4), true, float64(3), float64(0)),
		)

		DescribeTable("BindValidatedFloat64LessThan",
			func(given, should string, value float64, expectNil bool, threshold, dummy float64) {
				validator := paramSet.BindValidatedFloat64LessThan(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, threshold,
				)
				paramSet.Native.Gradientf64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, pattern, dummy float64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float64(2), true, float64(3), float64(0)),
			Entry(nil, "value is equal threshold", "return error", float64(3), false, float64(3), float64(0)),
			Entry(nil, "value is above threshold", "NOT return error", float64(4), false, float64(3), float64(0)),
		)

		DescribeTable("BindValidatedFloat64AtMost",
			func(given, should string, value float64, expectNil bool, threshold, dummy float64) {
				validator := paramSet.BindValidatedFloat64AtMost(
					assistant.NewFlagInfo("gradientf64", "t", float64(0)),
					&paramSet.Native.Gradientf64, threshold,
				)
				paramSet.Native.Gradientf64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value float64, expectNil bool, pattern, dummy float64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", float64(2), true, float64(3), float64(0)),
			Entry(nil, "value is equal threshold", "return error", float64(3), true, float64(3), float64(0)),
			Entry(nil, "value is above threshold", "NOT return error", float64(4), false, float64(3), float64(0)),
		)

		DescribeTable("BindValidatedIntWithin",
			func(given, should string, value int, expectNil bool, low, high int) {
				validator := paramSet.BindValidatedIntWithin(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, low, high int) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", 2, false, 3, 5),
			Entry(nil, "value is equal to low end of range", "return error", 3, true, 3, 5),
			Entry(nil, "value is inside range", "return error", 4, true, 3, 5),
			Entry(nil, "value is equal to high end of range", "return error", 5, true, 3, 5),
			Entry(nil, "value is above range", "NOT return error", 6, false, 3, 5),
		)

		DescribeTable("BindValidatedIntNotWithin",
			func(given, should string, value int, expectNil bool, low, high int) {
				validator := paramSet.BindValidatedIntNotWithin(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, low, high int) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", 2, false, 3, 5),
			Entry(nil, "value is equal to low end of range", "return error", 3, true, 3, 5),
			Entry(nil, "value is inside range", "return error", 4, true, 3, 5),
			Entry(nil, "value is equal to high end of range", "return error", 5, true, 3, 5),
			Entry(nil, "value is above range", "NOT return error", 6, false, 3, 5),
		)

		DescribeTable("BindValidatedContainsInt",
			func(given, should string, value int, expectNil bool, collection []int, dummy int) {
				validator := paramSet.BindValidatedContainsInt(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, collection,
				)
				paramSet.Native.Offset = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, collection []int, dummy int) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", 1, true, []int{1, 2, 3}, 0),
			Entry(nil, "collection does not contain member", "return error", 99, false, []int{1, 2, 3}, 0),
		)

		DescribeTable("BindValidatedNotContainsInt",
			func(given, should string, value int, expectNil bool, collection []int, dummy int) {
				validator := paramSet.BindValidatedNotContainsInt(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, collection,
				)
				paramSet.Native.Offset = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, collection []int, dummy int) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", 1, true, []int{1, 2, 3}, 0),
			Entry(nil, "collection does not contain member", "return error", 99, false, []int{1, 2, 3}, 0),
		)

		DescribeTable("BindValidatedIntGreaterThan",
			func(given, should string, value int, expectNil bool, threshold, dummy int) {
				validator := paramSet.BindValidatedIntGreaterThan(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, threshold,
				)
				paramSet.Native.Offset = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, pattern, dummy int) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", 2, false, 3, 0),
			Entry(nil, "value is equal threshold", "return error", 3, false, 3, 0),
			Entry(nil, "value is above threshold", "NOT return error", 4, true, 3, 0),
		)

		DescribeTable("BindValidatedIntAtLeast",
			func(given, should string, value int, expectNil bool, threshold, dummy int) {
				validator := paramSet.BindValidatedIntAtLeast(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, threshold,
				)
				paramSet.Native.Offset = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, pattern, dummy int) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", 2, false, 3, 0),
			Entry(nil, "value is equal threshold", "return error", 3, true, 3, 0),
			Entry(nil, "value is above threshold", "NOT return error", 4, true, 3, 0),
		)

		DescribeTable("BindValidatedIntLessThan",
			func(given, should string, value int, expectNil bool, threshold, dummy int) {
				validator := paramSet.BindValidatedIntLessThan(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, threshold,
				)
				paramSet.Native.Offset = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, pattern, dummy int) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", 2, true, 3, 0),
			Entry(nil, "value is equal threshold", "return error", 3, false, 3, 0),
			Entry(nil, "value is above threshold", "NOT return error", 4, false, 3, 0),
		)

		DescribeTable("BindValidatedIntAtMost",
			func(given, should string, value int, expectNil bool, threshold, dummy int) {
				validator := paramSet.BindValidatedIntAtMost(
					assistant.NewFlagInfo("offset", "o", -1),
					&paramSet.Native.Offset, threshold,
				)
				paramSet.Native.Offset = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int, expectNil bool, pattern, dummy int) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", 2, true, 3, 0),
			Entry(nil, "value is equal threshold", "return error", 3, true, 3, 0),
			Entry(nil, "value is above threshold", "NOT return error", 4, false, 3, 0),
		)

		DescribeTable("BindValidatedInt16Within",
			func(given, should string, value int16, expectNil bool, low, high int16) {
				validator := paramSet.BindValidatedInt16Within(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset16 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, low, high int16) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int16(2), false, int16(3), int16(5)),
			Entry(nil, "value is equal to low end of range", "return error", int16(3), true, int16(3), int16(5)),
			Entry(nil, "value is inside range", "return error", int16(4), true, int16(3), int16(5)),
			Entry(nil, "value is equal to high end of range", "return error", int16(5), true, int16(3), int16(5)),
			Entry(nil, "value is above range", "NOT return error", int16(6), false, int16(3), int16(5)),
		)

		DescribeTable("BindValidatedInt16NotWithin",
			func(given, should string, value int16, expectNil bool, low, high int16) {
				validator := paramSet.BindValidatedInt16NotWithin(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset16 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, low, high int16) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int16(2), false, int16(3), int16(5)),
			Entry(nil, "value is equal to low end of range", "return error", int16(3), true, int16(3), int16(5)),
			Entry(nil, "value is inside range", "return error", int16(4), true, int16(3), int16(5)),
			Entry(nil, "value is equal to high end of range", "return error", int16(5), true, int16(3), int16(5)),
			Entry(nil, "value is above range", "NOT return error", int16(6), false, int16(3), int16(5)),
		)

		DescribeTable("BindValidatedContainsInt16",
			func(given, should string, value int16, expectNil bool, collection []int16, dummy int16) {
				validator := paramSet.BindValidatedContainsInt16(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, collection,
				)
				paramSet.Native.Offset16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, collection []int16, dummy int16) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int16(1), true, []int16{1, 2, 3}, int16(0)),
			Entry(nil, "collection does not contain member", "return error", int16(99), false, []int16{1, 2, 3}, int16(0)),
		)

		DescribeTable("BindValidatedNotContainsInt16",
			func(given, should string, value int16, expectNil bool, collection []int16, dummy int16) {
				validator := paramSet.BindValidatedNotContainsInt16(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, collection,
				)
				paramSet.Native.Offset16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, collection []int16, dummy int16) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int16(1), true, []int16{1, 2, 3}, int16(0)),
			Entry(nil, "collection does not contain member", "return error", int16(99), false, []int16{1, 2, 3}, int16(0)),
		)

		DescribeTable("BindValidatedInt16GreaterThan",
			func(given, should string, value int16, expectNil bool, threshold, dummy int16) {
				validator := paramSet.BindValidatedInt16GreaterThan(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, threshold,
				)
				paramSet.Native.Offset16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, pattern, dummy int16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int16(2), false, int16(3), int16(0)),
			Entry(nil, "value is equal threshold", "return error", int16(3), false, int16(3), int16(0)),
			Entry(nil, "value is above threshold", "NOT return error", int16(4), true, int16(3), int16(0)),
		)

		DescribeTable("BindValidatedInt16AtLeast",
			func(given, should string, value int16, expectNil bool, threshold, dummy int16) {
				validator := paramSet.BindValidatedInt16AtLeast(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, threshold,
				)
				paramSet.Native.Offset16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, pattern, dummy int16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int16(2), false, int16(3), int16(0)),
			Entry(nil, "value is equal threshold", "return error", int16(3), true, int16(3), int16(0)),
			Entry(nil, "value is above threshold", "NOT return error", int16(4), true, int16(3), int16(0)),
		)

		DescribeTable("BindValidatedInt16LessThan",
			func(given, should string, value int16, expectNil bool, threshold, dummy int16) {
				validator := paramSet.BindValidatedInt16LessThan(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, threshold,
				)
				paramSet.Native.Offset16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, pattern, dummy int16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int16(2), true, int16(3), int16(0)),
			Entry(nil, "value is equal threshold", "return error", int16(3), false, int16(3), int16(0)),
			Entry(nil, "value is above threshold", "NOT return error", int16(4), false, int16(3), int16(0)),
		)

		DescribeTable("BindValidatedInt16AtMost",
			func(given, should string, value int16, expectNil bool, threshold, dummy int16) {
				validator := paramSet.BindValidatedInt16AtMost(
					assistant.NewFlagInfo("offset16", "o", int16(-1)),
					&paramSet.Native.Offset16, threshold,
				)
				paramSet.Native.Offset16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int16, expectNil bool, pattern, dummy int16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int16(2), true, int16(3), int16(0)),
			Entry(nil, "value is equal threshold", "return error", int16(3), true, int16(3), int16(0)),
			Entry(nil, "value is above threshold", "NOT return error", int16(4), false, int16(3), int16(0)),
		)

		DescribeTable("BindValidatedInt32Within",
			func(given, should string, value int32, expectNil bool, low, high int32) {
				validator := paramSet.BindValidatedInt32Within(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset32 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, low, high int32) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int32(2), false, int32(3), int32(5)),
			Entry(nil, "value is equal to low end of range", "return error", int32(3), true, int32(3), int32(5)),
			Entry(nil, "value is inside range", "return error", int32(4), true, int32(3), int32(5)),
			Entry(nil, "value is equal to high end of range", "return error", int32(5), true, int32(3), int32(5)),
			Entry(nil, "value is above range", "NOT return error", int32(6), false, int32(3), int32(5)),
		)

		DescribeTable("BindValidatedInt32NotWithin",
			func(given, should string, value int32, expectNil bool, low, high int32) {
				validator := paramSet.BindValidatedInt32NotWithin(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset32 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, low, high int32) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int32(2), false, int32(3), int32(5)),
			Entry(nil, "value is equal to low end of range", "return error", int32(3), true, int32(3), int32(5)),
			Entry(nil, "value is inside range", "return error", int32(4), true, int32(3), int32(5)),
			Entry(nil, "value is equal to high end of range", "return error", int32(5), true, int32(3), int32(5)),
			Entry(nil, "value is above range", "NOT return error", int32(6), false, int32(3), int32(5)),
		)

		DescribeTable("BindValidatedContainsInt32",
			func(given, should string, value int32, expectNil bool, collection []int32, dummy int32) {
				validator := paramSet.BindValidatedContainsInt32(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, collection,
				)
				paramSet.Native.Offset32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, collection []int32, dummy int32) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int32(1), true, []int32{1, 2, 3}, int32(0)),
			Entry(nil, "collection does not contain member", "return error", int32(99), false, []int32{1, 2, 3}, int32(0)),
		)

		DescribeTable("BindValidatedNotContainsInt32",
			func(given, should string, value int32, expectNil bool, collection []int32, dummy int32) {
				validator := paramSet.BindValidatedNotContainsInt32(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, collection,
				)
				paramSet.Native.Offset32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, collection []int32, dummy int32) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int32(1), true, []int32{1, 2, 3}, int32(0)),
			Entry(nil, "collection does not contain member", "return error", int32(99), false, []int32{1, 2, 3}, int32(0)),
		)

		DescribeTable("BindValidatedInt32GreaterThan",
			func(given, should string, value int32, expectNil bool, threshold, dummy int32) {
				validator := paramSet.BindValidatedInt32GreaterThan(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, threshold,
				)
				paramSet.Native.Offset32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, pattern, dummy int32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int32(2), false, int32(3), int32(0)),
			Entry(nil, "value is equal threshold", "return error", int32(3), false, int32(3), int32(0)),
			Entry(nil, "value is above threshold", "NOT return error", int32(4), true, int32(3), int32(0)),
		)

		DescribeTable("BindValidatedInt32AtLeast",
			func(given, should string, value int32, expectNil bool, threshold, dummy int32) {
				validator := paramSet.BindValidatedInt32AtLeast(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, threshold,
				)
				paramSet.Native.Offset32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, pattern, dummy int32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int32(2), false, int32(3), int32(0)),
			Entry(nil, "value is equal threshold", "return error", int32(3), true, int32(3), int32(0)),
			Entry(nil, "value is above threshold", "NOT return error", int32(4), true, int32(3), int32(0)),
		)

		DescribeTable("BindValidatedInt32LessThan",
			func(given, should string, value int32, expectNil bool, threshold, dummy int32) {
				validator := paramSet.BindValidatedInt32LessThan(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, threshold,
				)
				paramSet.Native.Offset32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, pattern, dummy int32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int32(2), true, int32(3), int32(0)),
			Entry(nil, "value is equal threshold", "return error", int32(3), false, int32(3), int32(0)),
			Entry(nil, "value is above threshold", "NOT return error", int32(4), false, int32(3), int32(0)),
		)

		DescribeTable("BindValidatedInt32AtMost",
			func(given, should string, value int32, expectNil bool, threshold, dummy int32) {
				validator := paramSet.BindValidatedInt32AtMost(
					assistant.NewFlagInfo("offset32", "o", int32(-1)),
					&paramSet.Native.Offset32, threshold,
				)
				paramSet.Native.Offset32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int32, expectNil bool, pattern, dummy int32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int32(2), true, int32(3), int32(0)),
			Entry(nil, "value is equal threshold", "return error", int32(3), true, int32(3), int32(0)),
			Entry(nil, "value is above threshold", "NOT return error", int32(4), false, int32(3), int32(0)),
		)

		DescribeTable("BindValidatedInt64Within",
			func(given, should string, value int64, expectNil bool, low, high int64) {
				validator := paramSet.BindValidatedInt64Within(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset64 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, low, high int64) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int64(2), false, int64(3), int64(5)),
			Entry(nil, "value is equal to low end of range", "return error", int64(3), true, int64(3), int64(5)),
			Entry(nil, "value is inside range", "return error", int64(4), true, int64(3), int64(5)),
			Entry(nil, "value is equal to high end of range", "return error", int64(5), true, int64(3), int64(5)),
			Entry(nil, "value is above range", "NOT return error", int64(6), false, int64(3), int64(5)),
		)

		DescribeTable("BindValidatedInt64NotWithin",
			func(given, should string, value int64, expectNil bool, low, high int64) {
				validator := paramSet.BindValidatedInt64NotWithin(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset64 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, low, high int64) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int64(2), false, int64(3), int64(5)),
			Entry(nil, "value is equal to low end of range", "return error", int64(3), true, int64(3), int64(5)),
			Entry(nil, "value is inside range", "return error", int64(4), true, int64(3), int64(5)),
			Entry(nil, "value is equal to high end of range", "return error", int64(5), true, int64(3), int64(5)),
			Entry(nil, "value is above range", "NOT return error", int64(6), false, int64(3), int64(5)),
		)

		DescribeTable("BindValidatedContainsInt64",
			func(given, should string, value int64, expectNil bool, collection []int64, dummy int64) {
				validator := paramSet.BindValidatedContainsInt64(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, collection,
				)
				paramSet.Native.Offset64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, collection []int64, dummy int64) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int64(1), true, []int64{1, 2, 3}, int64(0)),
			Entry(nil, "collection does not contain member", "return error", int64(99), false, []int64{1, 2, 3}, int64(0)),
		)

		DescribeTable("BindValidatedNotContainsInt64",
			func(given, should string, value int64, expectNil bool, collection []int64, dummy int64) {
				validator := paramSet.BindValidatedNotContainsInt64(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, collection,
				)
				paramSet.Native.Offset64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, collection []int64, dummy int64) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int64(1), true, []int64{1, 2, 3}, int64(0)),
			Entry(nil, "collection does not contain member", "return error", int64(99), false, []int64{1, 2, 3}, int64(0)),
		)

		DescribeTable("BindValidatedInt64GreaterThan",
			func(given, should string, value int64, expectNil bool, threshold, dummy int64) {
				validator := paramSet.BindValidatedInt64GreaterThan(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, threshold,
				)
				paramSet.Native.Offset64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, pattern, dummy int64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int64(2), false, int64(3), int64(0)),
			Entry(nil, "value is equal threshold", "return error", int64(3), false, int64(3), int64(0)),
			Entry(nil, "value is above threshold", "NOT return error", int64(4), true, int64(3), int64(0)),
		)

		DescribeTable("BindValidatedInt64AtLeast",
			func(given, should string, value int64, expectNil bool, threshold, dummy int64) {
				validator := paramSet.BindValidatedInt64AtLeast(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, threshold,
				)
				paramSet.Native.Offset64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, pattern, dummy int64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int64(2), false, int64(3), int64(0)),
			Entry(nil, "value is equal threshold", "return error", int64(3), true, int64(3), int64(0)),
			Entry(nil, "value is above threshold", "NOT return error", int64(4), true, int64(3), int64(0)),
		)

		DescribeTable("BindValidatedInt64LessThan",
			func(given, should string, value int64, expectNil bool, threshold, dummy int64) {
				validator := paramSet.BindValidatedInt64LessThan(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, threshold,
				)
				paramSet.Native.Offset64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, pattern, dummy int64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int64(2), true, int64(3), int64(0)),
			Entry(nil, "value is equal threshold", "return error", int64(3), false, int64(3), int64(0)),
			Entry(nil, "value is above threshold", "NOT return error", int64(4), false, int64(3), int64(0)),
		)

		DescribeTable("BindValidatedInt64AtMost",
			func(given, should string, value int64, expectNil bool, threshold, dummy int64) {
				validator := paramSet.BindValidatedInt64AtMost(
					assistant.NewFlagInfo("offset64", "o", int64(-1)),
					&paramSet.Native.Offset64, threshold,
				)
				paramSet.Native.Offset64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int64, expectNil bool, pattern, dummy int64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int64(2), true, int64(3), int64(0)),
			Entry(nil, "value is equal threshold", "return error", int64(3), true, int64(3), int64(0)),
			Entry(nil, "value is above threshold", "NOT return error", int64(4), false, int64(3), int64(0)),
		)

		DescribeTable("BindValidatedInt8Within",
			func(given, should string, value int8, expectNil bool, low, high int8) {
				validator := paramSet.BindValidatedInt8Within(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset8 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, low, high int8) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int8(2), false, int8(3), int8(5)),
			Entry(nil, "value is equal to low end of range", "return error", int8(3), true, int8(3), int8(5)),
			Entry(nil, "value is inside range", "return error", int8(4), true, int8(3), int8(5)),
			Entry(nil, "value is equal to high end of range", "return error", int8(5), true, int8(3), int8(5)),
			Entry(nil, "value is above range", "NOT return error", int8(6), false, int8(3), int8(5)),
		)

		DescribeTable("BindValidatedInt8NotWithin",
			func(given, should string, value int8, expectNil bool, low, high int8) {
				validator := paramSet.BindValidatedInt8NotWithin(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Offset8 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, low, high int8) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", int8(2), false, int8(3), int8(5)),
			Entry(nil, "value is equal to low end of range", "return error", int8(3), true, int8(3), int8(5)),
			Entry(nil, "value is inside range", "return error", int8(4), true, int8(3), int8(5)),
			Entry(nil, "value is equal to high end of range", "return error", int8(5), true, int8(3), int8(5)),
			Entry(nil, "value is above range", "NOT return error", int8(6), false, int8(3), int8(5)),
		)

		DescribeTable("BindValidatedContainsInt8",
			func(given, should string, value int8, expectNil bool, collection []int8, dummy int8) {
				validator := paramSet.BindValidatedContainsInt8(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, collection,
				)
				paramSet.Native.Offset8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, collection []int8, dummy int8) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int8(1), true, []int8{1, 2, 3}, int8(0)),
			Entry(nil, "collection does not contain member", "return error", int8(99), false, []int8{1, 2, 3}, int8(0)),
		)

		DescribeTable("BindValidatedNotContainsInt8",
			func(given, should string, value int8, expectNil bool, collection []int8, dummy int8) {
				validator := paramSet.BindValidatedNotContainsInt8(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, collection,
				)
				paramSet.Native.Offset8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, collection []int8, dummy int8) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", int8(1), true, []int8{1, 2, 3}, int8(0)),
			Entry(nil, "collection does not contain member", "return error", int8(99), false, []int8{1, 2, 3}, int8(0)),
		)

		DescribeTable("BindValidatedInt8GreaterThan",
			func(given, should string, value int8, expectNil bool, threshold, dummy int8) {
				validator := paramSet.BindValidatedInt8GreaterThan(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, threshold,
				)
				paramSet.Native.Offset8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, pattern, dummy int8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int8(2), false, int8(3), int8(0)),
			Entry(nil, "value is equal threshold", "return error", int8(3), false, int8(3), int8(0)),
			Entry(nil, "value is above threshold", "NOT return error", int8(4), true, int8(3), int8(0)),
		)

		DescribeTable("BindValidatedInt8AtLeast",
			func(given, should string, value int8, expectNil bool, threshold, dummy int8) {
				validator := paramSet.BindValidatedInt8AtLeast(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, threshold,
				)
				paramSet.Native.Offset8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, pattern, dummy int8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int8(2), false, int8(3), int8(0)),
			Entry(nil, "value is equal threshold", "return error", int8(3), true, int8(3), int8(0)),
			Entry(nil, "value is above threshold", "NOT return error", int8(4), true, int8(3), int8(0)),
		)

		DescribeTable("BindValidatedInt8LessThan",
			func(given, should string, value int8, expectNil bool, threshold, dummy int8) {
				validator := paramSet.BindValidatedInt8LessThan(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, threshold,
				)
				paramSet.Native.Offset8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, pattern, dummy int8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int8(2), true, int8(3), int8(0)),
			Entry(nil, "value is equal threshold", "return error", int8(3), false, int8(3), int8(0)),
			Entry(nil, "value is above threshold", "NOT return error", int8(4), false, int8(3), int8(0)),
		)

		DescribeTable("BindValidatedInt8AtMost",
			func(given, should string, value int8, expectNil bool, threshold, dummy int8) {
				validator := paramSet.BindValidatedInt8AtMost(
					assistant.NewFlagInfo("offset8", "o", int8(-1)),
					&paramSet.Native.Offset8, threshold,
				)
				paramSet.Native.Offset8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value int8, expectNil bool, pattern, dummy int8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", int8(2), true, int8(3), int8(0)),
			Entry(nil, "value is equal threshold", "return error", int8(3), true, int8(3), int8(0)),
			Entry(nil, "value is above threshold", "NOT return error", int8(4), false, int8(3), int8(0)),
		)

		DescribeTable("BindValidatedStringWithin",
			func(given, should string, value string, expectNil bool, low, high string) {
				validator := paramSet.BindValidatedStringWithin(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Pattern = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, low, high string) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", "b", false, "c", "e"),
			Entry(nil, "value is equal to low end of range", "return error", "c", true, "c", "e"),
			Entry(nil, "value is inside range", "return error", "d", true, "c", "e"),
			Entry(nil, "value is equal to high end of range", "return error", "e", true, "c", "e"),
			Entry(nil, "value is above range", "NOT return error", "f", false, "c", "e"),
		)

		DescribeTable("BindValidatedStringNotWithin",
			func(given, should string, value string, expectNil bool, low, high string) {
				validator := paramSet.BindValidatedStringNotWithin(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Pattern = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, low, high string) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
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
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, collection,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, collection []string, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", "a", true, []string{"a", "b", "c"}, "null"),
			Entry(nil, "collection does not contain member", "return error", "x", false, []string{"a", "b", "c"}, "null"),
		)

		DescribeTable("BindValidatedNotContainsString",
			func(given, should string, value string, expectNil bool, collection []string, dummy string) {
				validator := paramSet.BindValidatedNotContainsString(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, collection,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, collection []string, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", "a", true, []string{"a", "b", "c"}, "null"),
			Entry(nil, "collection does not contain member", "return error", "x", false, []string{"a", "b", "c"}, "null"),
		)

		DescribeTable("BindValidatedStringIsMatch",
			func(given, should string, value string, expectNil bool, pattern, dummy string) {
				validator := paramSet.BindValidatedStringIsMatch(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, pattern,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍇 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value matches pattern", "return error", "18-10-1997", true, "\\d{2}-\\d{2}-\\d{4}", "null"),
			Entry(nil, "value does not match pattern", "return error", "foo-bar", false, "\\d{2}-\\d{2}-\\d{4}", "null"),
		)

		DescribeTable("BindValidatedStringIsNotMatch",
			func(given, should string, value string, expectNil bool, pattern, dummy string) {
				validator := paramSet.BindValidatedStringIsNotMatch(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, pattern,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍇 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value matches pattern", "return error", "18-10-1997", true, "\\d{2}-\\d{2}-\\d{4}", "null"),
			Entry(nil, "value does not match pattern", "return error", "foo-bar", false, "\\d{2}-\\d{2}-\\d{4}", "null"),
		)

		DescribeTable("BindValidatedStringGreaterThan",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringGreaterThan(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", false, "c", ""),
			Entry(nil, "value is equal threshold", "return error", "c", false, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", true, "c", ""),
		)

		DescribeTable("BindValidatedStringAtLeast",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringAtLeast(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", false, "c", ""),
			Entry(nil, "value is equal threshold", "return error", "c", true, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", true, "c", ""),
		)

		DescribeTable("BindValidatedStringLessThan",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringLessThan(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", true, "c", ""),
			Entry(nil, "value is equal threshold", "return error", "c", false, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", false, "c", ""),
		)

		DescribeTable("BindValidatedStringAtMost",
			func(given, should string, value string, expectNil bool, threshold, dummy string) {
				validator := paramSet.BindValidatedStringAtMost(
					assistant.NewFlagInfo("pattern", "p", "default-pattern"),
					&paramSet.Native.Pattern, threshold,
				)
				paramSet.Native.Pattern = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value string, expectNil bool, pattern, dummy string) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", "b", true, "c", ""),
			Entry(nil, "value is equal threshold", "return error", "c", true, "c", ""),
			Entry(nil, "value is above threshold", "NOT return error", "d", false, "c", ""),
		)

		DescribeTable("BindValidatedUint16Within",
			func(given, should string, value uint16, expectNil bool, low, high uint16) {
				validator := paramSet.BindValidatedUint16Within(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count16 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, low, high uint16) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint16(2), false, uint16(3), uint16(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint16(3), true, uint16(3), uint16(5)),
			Entry(nil, "value is inside range", "return error", uint16(4), true, uint16(3), uint16(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint16(5), true, uint16(3), uint16(5)),
			Entry(nil, "value is above range", "NOT return error", uint16(6), false, uint16(3), uint16(5)),
		)

		DescribeTable("BindValidatedUint16NotWithin",
			func(given, should string, value uint16, expectNil bool, low, high uint16) {
				validator := paramSet.BindValidatedUint16NotWithin(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count16 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, low, high uint16) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint16(2), false, uint16(3), uint16(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint16(3), true, uint16(3), uint16(5)),
			Entry(nil, "value is inside range", "return error", uint16(4), true, uint16(3), uint16(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint16(5), true, uint16(3), uint16(5)),
			Entry(nil, "value is above range", "NOT return error", uint16(6), false, uint16(3), uint16(5)),
		)

		DescribeTable("BindValidatedContainsUint16",
			func(given, should string, value uint16, expectNil bool, collection []uint16, dummy uint16) {
				validator := paramSet.BindValidatedContainsUint16(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, collection,
				)
				paramSet.Native.Count16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, collection []uint16, dummy uint16) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint16(1), true, []uint16{1, 2, 3}, uint16(0)),
			Entry(nil, "collection does not contain member", "return error", uint16(99), false, []uint16{1, 2, 3}, uint16(0)),
		)

		DescribeTable("BindValidatedNotContainsUint16",
			func(given, should string, value uint16, expectNil bool, collection []uint16, dummy uint16) {
				validator := paramSet.BindValidatedNotContainsUint16(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, collection,
				)
				paramSet.Native.Count16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, collection []uint16, dummy uint16) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint16(1), true, []uint16{1, 2, 3}, uint16(0)),
			Entry(nil, "collection does not contain member", "return error", uint16(99), false, []uint16{1, 2, 3}, uint16(0)),
		)

		DescribeTable("BindValidatedUint16GreaterThan",
			func(given, should string, value uint16, expectNil bool, threshold, dummy uint16) {
				validator := paramSet.BindValidatedUint16GreaterThan(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, threshold,
				)
				paramSet.Native.Count16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, pattern, dummy uint16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint16(2), false, uint16(3), uint16(0)),
			Entry(nil, "value is equal threshold", "return error", uint16(3), false, uint16(3), uint16(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint16(4), true, uint16(3), uint16(0)),
		)

		DescribeTable("BindValidatedUint16AtLeast",
			func(given, should string, value uint16, expectNil bool, threshold, dummy uint16) {
				validator := paramSet.BindValidatedUint16AtLeast(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, threshold,
				)
				paramSet.Native.Count16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, pattern, dummy uint16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint16(2), false, uint16(3), uint16(0)),
			Entry(nil, "value is equal threshold", "return error", uint16(3), true, uint16(3), uint16(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint16(4), true, uint16(3), uint16(0)),
		)

		DescribeTable("BindValidatedUint16LessThan",
			func(given, should string, value uint16, expectNil bool, threshold, dummy uint16) {
				validator := paramSet.BindValidatedUint16LessThan(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, threshold,
				)
				paramSet.Native.Count16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, pattern, dummy uint16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint16(2), true, uint16(3), uint16(0)),
			Entry(nil, "value is equal threshold", "return error", uint16(3), false, uint16(3), uint16(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint16(4), false, uint16(3), uint16(0)),
		)

		DescribeTable("BindValidatedUint16AtMost",
			func(given, should string, value uint16, expectNil bool, threshold, dummy uint16) {
				validator := paramSet.BindValidatedUint16AtMost(
					assistant.NewFlagInfo("count16", "c", uint16(0)),
					&paramSet.Native.Count16, threshold,
				)
				paramSet.Native.Count16 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint16, expectNil bool, pattern, dummy uint16) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint16(2), true, uint16(3), uint16(0)),
			Entry(nil, "value is equal threshold", "return error", uint16(3), true, uint16(3), uint16(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint16(4), false, uint16(3), uint16(0)),
		)

		DescribeTable("BindValidatedUint32Within",
			func(given, should string, value uint32, expectNil bool, low, high uint32) {
				validator := paramSet.BindValidatedUint32Within(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count32 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, low, high uint32) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint32(2), false, uint32(3), uint32(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint32(3), true, uint32(3), uint32(5)),
			Entry(nil, "value is inside range", "return error", uint32(4), true, uint32(3), uint32(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint32(5), true, uint32(3), uint32(5)),
			Entry(nil, "value is above range", "NOT return error", uint32(6), false, uint32(3), uint32(5)),
		)

		DescribeTable("BindValidatedUint32NotWithin",
			func(given, should string, value uint32, expectNil bool, low, high uint32) {
				validator := paramSet.BindValidatedUint32NotWithin(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count32 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, low, high uint32) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint32(2), false, uint32(3), uint32(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint32(3), true, uint32(3), uint32(5)),
			Entry(nil, "value is inside range", "return error", uint32(4), true, uint32(3), uint32(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint32(5), true, uint32(3), uint32(5)),
			Entry(nil, "value is above range", "NOT return error", uint32(6), false, uint32(3), uint32(5)),
		)

		DescribeTable("BindValidatedContainsUint32",
			func(given, should string, value uint32, expectNil bool, collection []uint32, dummy uint32) {
				validator := paramSet.BindValidatedContainsUint32(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, collection,
				)
				paramSet.Native.Count32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, collection []uint32, dummy uint32) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint32(1), true, []uint32{1, 2, 3}, uint32(0)),
			Entry(nil, "collection does not contain member", "return error", uint32(99), false, []uint32{1, 2, 3}, uint32(0)),
		)

		DescribeTable("BindValidatedNotContainsUint32",
			func(given, should string, value uint32, expectNil bool, collection []uint32, dummy uint32) {
				validator := paramSet.BindValidatedNotContainsUint32(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, collection,
				)
				paramSet.Native.Count32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, collection []uint32, dummy uint32) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint32(1), true, []uint32{1, 2, 3}, uint32(0)),
			Entry(nil, "collection does not contain member", "return error", uint32(99), false, []uint32{1, 2, 3}, uint32(0)),
		)

		DescribeTable("BindValidatedUint32GreaterThan",
			func(given, should string, value uint32, expectNil bool, threshold, dummy uint32) {
				validator := paramSet.BindValidatedUint32GreaterThan(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, threshold,
				)
				paramSet.Native.Count32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, pattern, dummy uint32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint32(2), false, uint32(3), uint32(0)),
			Entry(nil, "value is equal threshold", "return error", uint32(3), false, uint32(3), uint32(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint32(4), true, uint32(3), uint32(0)),
		)

		DescribeTable("BindValidatedUint32AtLeast",
			func(given, should string, value uint32, expectNil bool, threshold, dummy uint32) {
				validator := paramSet.BindValidatedUint32AtLeast(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, threshold,
				)
				paramSet.Native.Count32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, pattern, dummy uint32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint32(2), false, uint32(3), uint32(0)),
			Entry(nil, "value is equal threshold", "return error", uint32(3), true, uint32(3), uint32(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint32(4), true, uint32(3), uint32(0)),
		)

		DescribeTable("BindValidatedUint32LessThan",
			func(given, should string, value uint32, expectNil bool, threshold, dummy uint32) {
				validator := paramSet.BindValidatedUint32LessThan(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, threshold,
				)
				paramSet.Native.Count32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, pattern, dummy uint32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint32(2), true, uint32(3), uint32(0)),
			Entry(nil, "value is equal threshold", "return error", uint32(3), false, uint32(3), uint32(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint32(4), false, uint32(3), uint32(0)),
		)

		DescribeTable("BindValidatedUint32AtMost",
			func(given, should string, value uint32, expectNil bool, threshold, dummy uint32) {
				validator := paramSet.BindValidatedUint32AtMost(
					assistant.NewFlagInfo("count32", "c", uint32(0)),
					&paramSet.Native.Count32, threshold,
				)
				paramSet.Native.Count32 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint32, expectNil bool, pattern, dummy uint32) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint32(2), true, uint32(3), uint32(0)),
			Entry(nil, "value is equal threshold", "return error", uint32(3), true, uint32(3), uint32(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint32(4), false, uint32(3), uint32(0)),
		)

		DescribeTable("BindValidatedUint64Within",
			func(given, should string, value uint64, expectNil bool, low, high uint64) {
				validator := paramSet.BindValidatedUint64Within(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count64 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, low, high uint64) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint64(2), false, uint64(3), uint64(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint64(3), true, uint64(3), uint64(5)),
			Entry(nil, "value is inside range", "return error", uint64(4), true, uint64(3), uint64(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint64(5), true, uint64(3), uint64(5)),
			Entry(nil, "value is above range", "NOT return error", uint64(6), false, uint64(3), uint64(5)),
		)

		DescribeTable("BindValidatedUint64NotWithin",
			func(given, should string, value uint64, expectNil bool, low, high uint64) {
				validator := paramSet.BindValidatedUint64NotWithin(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count64 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, low, high uint64) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint64(2), false, uint64(3), uint64(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint64(3), true, uint64(3), uint64(5)),
			Entry(nil, "value is inside range", "return error", uint64(4), true, uint64(3), uint64(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint64(5), true, uint64(3), uint64(5)),
			Entry(nil, "value is above range", "NOT return error", uint64(6), false, uint64(3), uint64(5)),
		)

		DescribeTable("BindValidatedContainsUint64",
			func(given, should string, value uint64, expectNil bool, collection []uint64, dummy uint64) {
				validator := paramSet.BindValidatedContainsUint64(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, collection,
				)
				paramSet.Native.Count64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, collection []uint64, dummy uint64) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint64(1), true, []uint64{1, 2, 3}, uint64(0)),
			Entry(nil, "collection does not contain member", "return error", uint64(99), false, []uint64{1, 2, 3}, uint64(0)),
		)

		DescribeTable("BindValidatedNotContainsUint64",
			func(given, should string, value uint64, expectNil bool, collection []uint64, dummy uint64) {
				validator := paramSet.BindValidatedNotContainsUint64(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, collection,
				)
				paramSet.Native.Count64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, collection []uint64, dummy uint64) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint64(1), true, []uint64{1, 2, 3}, uint64(0)),
			Entry(nil, "collection does not contain member", "return error", uint64(99), false, []uint64{1, 2, 3}, uint64(0)),
		)

		DescribeTable("BindValidatedUint64GreaterThan",
			func(given, should string, value uint64, expectNil bool, threshold, dummy uint64) {
				validator := paramSet.BindValidatedUint64GreaterThan(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, threshold,
				)
				paramSet.Native.Count64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, pattern, dummy uint64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint64(2), false, uint64(3), uint64(0)),
			Entry(nil, "value is equal threshold", "return error", uint64(3), false, uint64(3), uint64(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint64(4), true, uint64(3), uint64(0)),
		)

		DescribeTable("BindValidatedUint64AtLeast",
			func(given, should string, value uint64, expectNil bool, threshold, dummy uint64) {
				validator := paramSet.BindValidatedUint64AtLeast(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, threshold,
				)
				paramSet.Native.Count64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, pattern, dummy uint64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint64(2), false, uint64(3), uint64(0)),
			Entry(nil, "value is equal threshold", "return error", uint64(3), true, uint64(3), uint64(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint64(4), true, uint64(3), uint64(0)),
		)

		DescribeTable("BindValidatedUint64LessThan",
			func(given, should string, value uint64, expectNil bool, threshold, dummy uint64) {
				validator := paramSet.BindValidatedUint64LessThan(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, threshold,
				)
				paramSet.Native.Count64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, pattern, dummy uint64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint64(2), true, uint64(3), uint64(0)),
			Entry(nil, "value is equal threshold", "return error", uint64(3), false, uint64(3), uint64(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint64(4), false, uint64(3), uint64(0)),
		)

		DescribeTable("BindValidatedUint64AtMost",
			func(given, should string, value uint64, expectNil bool, threshold, dummy uint64) {
				validator := paramSet.BindValidatedUint64AtMost(
					assistant.NewFlagInfo("count64", "c", uint64(0)),
					&paramSet.Native.Count64, threshold,
				)
				paramSet.Native.Count64 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint64, expectNil bool, pattern, dummy uint64) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint64(2), true, uint64(3), uint64(0)),
			Entry(nil, "value is equal threshold", "return error", uint64(3), true, uint64(3), uint64(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint64(4), false, uint64(3), uint64(0)),
		)

		DescribeTable("BindValidatedUint8Within",
			func(given, should string, value uint8, expectNil bool, low, high uint8) {
				validator := paramSet.BindValidatedUint8Within(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count8 = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, low, high uint8) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint8(2), false, uint8(3), uint8(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint8(3), true, uint8(3), uint8(5)),
			Entry(nil, "value is inside range", "return error", uint8(4), true, uint8(3), uint8(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint8(5), true, uint8(3), uint8(5)),
			Entry(nil, "value is above range", "NOT return error", uint8(6), false, uint8(3), uint8(5)),
		)

		DescribeTable("BindValidatedUint8NotWithin",
			func(given, should string, value uint8, expectNil bool, low, high uint8) {
				validator := paramSet.BindValidatedUint8NotWithin(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count8 = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, low, high uint8) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint8(2), false, uint8(3), uint8(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint8(3), true, uint8(3), uint8(5)),
			Entry(nil, "value is inside range", "return error", uint8(4), true, uint8(3), uint8(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint8(5), true, uint8(3), uint8(5)),
			Entry(nil, "value is above range", "NOT return error", uint8(6), false, uint8(3), uint8(5)),
		)

		DescribeTable("BindValidatedContainsUint8",
			func(given, should string, value uint8, expectNil bool, collection []uint8, dummy uint8) {
				validator := paramSet.BindValidatedContainsUint8(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, collection,
				)
				paramSet.Native.Count8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, collection []uint8, dummy uint8) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint8(1), true, []uint8{1, 2, 3}, uint8(0)),
			Entry(nil, "collection does not contain member", "return error", uint8(99), false, []uint8{1, 2, 3}, uint8(0)),
		)

		DescribeTable("BindValidatedNotContainsUint8",
			func(given, should string, value uint8, expectNil bool, collection []uint8, dummy uint8) {
				validator := paramSet.BindValidatedNotContainsUint8(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, collection,
				)
				paramSet.Native.Count8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, collection []uint8, dummy uint8) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint8(1), true, []uint8{1, 2, 3}, uint8(0)),
			Entry(nil, "collection does not contain member", "return error", uint8(99), false, []uint8{1, 2, 3}, uint8(0)),
		)

		DescribeTable("BindValidatedUint8GreaterThan",
			func(given, should string, value uint8, expectNil bool, threshold, dummy uint8) {
				validator := paramSet.BindValidatedUint8GreaterThan(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, threshold,
				)
				paramSet.Native.Count8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, pattern, dummy uint8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint8(2), false, uint8(3), uint8(0)),
			Entry(nil, "value is equal threshold", "return error", uint8(3), false, uint8(3), uint8(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint8(4), true, uint8(3), uint8(0)),
		)

		DescribeTable("BindValidatedUint8AtLeast",
			func(given, should string, value uint8, expectNil bool, threshold, dummy uint8) {
				validator := paramSet.BindValidatedUint8AtLeast(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, threshold,
				)
				paramSet.Native.Count8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, pattern, dummy uint8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint8(2), false, uint8(3), uint8(0)),
			Entry(nil, "value is equal threshold", "return error", uint8(3), true, uint8(3), uint8(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint8(4), true, uint8(3), uint8(0)),
		)

		DescribeTable("BindValidatedUint8LessThan",
			func(given, should string, value uint8, expectNil bool, threshold, dummy uint8) {
				validator := paramSet.BindValidatedUint8LessThan(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, threshold,
				)
				paramSet.Native.Count8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, pattern, dummy uint8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint8(2), true, uint8(3), uint8(0)),
			Entry(nil, "value is equal threshold", "return error", uint8(3), false, uint8(3), uint8(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint8(4), false, uint8(3), uint8(0)),
		)

		DescribeTable("BindValidatedUint8AtMost",
			func(given, should string, value uint8, expectNil bool, threshold, dummy uint8) {
				validator := paramSet.BindValidatedUint8AtMost(
					assistant.NewFlagInfo("count8", "c", uint8(0)),
					&paramSet.Native.Count8, threshold,
				)
				paramSet.Native.Count8 = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint8, expectNil bool, pattern, dummy uint8) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint8(2), true, uint8(3), uint8(0)),
			Entry(nil, "value is equal threshold", "return error", uint8(3), true, uint8(3), uint8(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint8(4), false, uint8(3), uint8(0)),
		)

		DescribeTable("BindValidatedUintWithin",
			func(given, should string, value uint, expectNil bool, low, high uint) {
				validator := paramSet.BindValidatedUintWithin(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count = value

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, low, high uint) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint(2), false, uint(3), uint(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint(3), true, uint(3), uint(5)),
			Entry(nil, "value is inside range", "return error", uint(4), true, uint(3), uint(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint(5), true, uint(3), uint(5)),
			Entry(nil, "value is above range", "NOT return error", uint(6), false, uint(3), uint(5)),
		)

		DescribeTable("BindValidatedUintNotWithin",
			func(given, should string, value uint, expectNil bool, low, high uint) {
				validator := paramSet.BindValidatedUintNotWithin(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, low, high,
				)
				decorator := validatorDecorator{
					Decorated: validator,
				}

				paramSet.Native.Count = value

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, low, high uint) string {
				return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below range", "return error", uint(2), false, uint(3), uint(5)),
			Entry(nil, "value is equal to low end of range", "return error", uint(3), true, uint(3), uint(5)),
			Entry(nil, "value is inside range", "return error", uint(4), true, uint(3), uint(5)),
			Entry(nil, "value is equal to high end of range", "return error", uint(5), true, uint(3), uint(5)),
			Entry(nil, "value is above range", "NOT return error", uint(6), false, uint(3), uint(5)),
		)

		DescribeTable("BindValidatedContainsUint",
			func(given, should string, value uint, expectNil bool, collection []uint, dummy uint) {
				validator := paramSet.BindValidatedContainsUint(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, collection,
				)
				paramSet.Native.Count = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, collection []uint, dummy uint) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint(1), true, []uint{1, 2, 3}, uint(0)),
			Entry(nil, "collection does not contain member", "return error", uint(99), false, []uint{1, 2, 3}, uint(0)),
		)

		DescribeTable("BindValidatedNotContainsUint",
			func(given, should string, value uint, expectNil bool, collection []uint, dummy uint) {
				validator := paramSet.BindValidatedNotContainsUint(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, collection,
				)
				paramSet.Native.Count = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if !expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, collection []uint, dummy uint) string {
				return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "collection contains member", "return error", uint(1), true, []uint{1, 2, 3}, uint(0)),
			Entry(nil, "collection does not contain member", "return error", uint(99), false, []uint{1, 2, 3}, uint(0)),
		)

		DescribeTable("BindValidatedUintGreaterThan",
			func(given, should string, value uint, expectNil bool, threshold, dummy uint) {
				validator := paramSet.BindValidatedUintGreaterThan(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, threshold,
				)
				paramSet.Native.Count = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, pattern, dummy uint) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint(2), false, uint(3), uint(0)),
			Entry(nil, "value is equal threshold", "return error", uint(3), false, uint(3), uint(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint(4), true, uint(3), uint(0)),
		)

		DescribeTable("BindValidatedUintAtLeast",
			func(given, should string, value uint, expectNil bool, threshold, dummy uint) {
				validator := paramSet.BindValidatedUintAtLeast(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, threshold,
				)
				paramSet.Native.Count = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, pattern, dummy uint) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint(2), false, uint(3), uint(0)),
			Entry(nil, "value is equal threshold", "return error", uint(3), true, uint(3), uint(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint(4), true, uint(3), uint(0)),
		)

		DescribeTable("BindValidatedUintLessThan",
			func(given, should string, value uint, expectNil bool, threshold, dummy uint) {
				validator := paramSet.BindValidatedUintLessThan(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, threshold,
				)
				paramSet.Native.Count = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, pattern, dummy uint) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint(2), true, uint(3), uint(0)),
			Entry(nil, "value is equal threshold", "return error", uint(3), false, uint(3), uint(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint(4), false, uint(3), uint(0)),
		)

		DescribeTable("BindValidatedUintAtMost",
			func(given, should string, value uint, expectNil bool, threshold, dummy uint) {
				validator := paramSet.BindValidatedUintAtMost(
					assistant.NewFlagInfo("count", "c", uint(0)),
					&paramSet.Native.Count, threshold,
				)
				paramSet.Native.Count = value
				decorator := validatorDecorator{
					Decorated: validator,
				}

				if expectNil {
					Expect(decorator.Validate()).Error().To(BeNil())
				} else {
					Expect(decorator.Validate()).Error().ToNot(BeNil())
				}
			},
			func(given, should string, value uint, expectNil bool, pattern, dummy uint) string {
				return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
					given, should)
			},
			Entry(nil, "value is below threshold", "return error", uint(2), true, uint(3), uint(0)),
			Entry(nil, "value is equal threshold", "return error", uint(3), true, uint(3), uint(0)),
			Entry(nil, "value is above threshold", "NOT return error", uint(4), false, uint(3), uint(0)),
		)

		// <---- auto generated
	})
})
