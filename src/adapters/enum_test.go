package adapters_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
	"github.com/snivilised/cobrass/src/testhelpers"
)

var _ = Describe("Enum", func() {
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *adapters.ParamSet[WidgetParameterSet]
	Context("given: enum type", func() {
		var outputFormatEnumInfo *adapters.EnumInfo[OutputFormatEnum]
		var outputFormatEnumSlice adapters.EnumSlice[OutputFormatEnum]

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
			}
			rootCommand.AddCommand(widgetCommand)
			paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)

			outputFormatEnumInfo = adapters.NewEnumInfo(adapters.AcceptableEnumValues[OutputFormatEnum]{
				XmlFormatEn:      []string{"xml", "x"},
				JsonFormatEn:     []string{"json", "j"},
				TextFormatEn:     []string{"text", "tx"},
				ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
			})
			outputFormatEnumSlice = outputFormatEnumInfo.NewSlice()
		})

		It("🧪 should: create enum info", func() {
			Expect(outputFormatEnumInfo.En("x")).To(Equal(XmlFormatEn))
			Expect(outputFormatEnumInfo.En("xml")).To(Equal(XmlFormatEn))

			Expect(outputFormatEnumInfo.En("j")).To(Equal(JsonFormatEn))
			Expect(outputFormatEnumInfo.En("json")).To(Equal(JsonFormatEn))

			Expect(outputFormatEnumInfo.NameOf(XmlFormatEn)).To(Equal("xml"))
		})

		Context("EnumInfo", func() {
			Context("IsValid", func() {
				When("given: valid value", func() {
					It("🧪 should: return true", func() {
						Expect(outputFormatEnumInfo.IsValid("text")).To(BeTrue())
					})
				})

				When("given: invalid value", func() {
					It("🧪 should: return false", func() {
						Expect(outputFormatEnumInfo.IsValid("foo")).To(BeFalse())
					})
				})
			})

			Context("IsValidOrEmpty", func() {
				When("given: empty value", func() {
					It("🧪 should: return true", func() {
						outputFormatEnum := outputFormatEnumInfo.NewValue()
						Expect(outputFormatEnum.IsValidOrEmpty()).To(BeTrue())
					})
				})
			})

			Context("EnumInfo.String", func() {
				It("🧪 should: return contents of enum info", func() {
					result := outputFormatEnumInfo.String()
					GinkgoWriter.Println("===> contents of OutputFormatEnumInfo:")
					GinkgoWriter.Println(result)
				})
			})
		})

		Context("EnumSlice", func() {
			DescribeTable("Values",
				func(given, should string, source []string, expect []OutputFormatEnum) {
					outputFormatEnumSlice.Source = source

					Expect(reflect.DeepEqual(
						outputFormatEnumSlice.Values(), expect,
					)).To(BeTrue())
				},
				func(given, should string, source []string, expect []OutputFormatEnum) string {
					return fmt.Sprintf("🧪 --> 🍈 given: '%v', should: '%v'",
						given, should)
				},
				Entry(
					nil, "all valid unique values", "return int enum slice",
					[]string{"json", "text", "xml"}, []OutputFormatEnum{JsonFormatEn, TextFormatEn, XmlFormatEn},
				),
				Entry(
					nil, "all valid same values", "return int enum slice",
					[]string{"json", "json", "json"}, []OutputFormatEnum{JsonFormatEn, JsonFormatEn, JsonFormatEn},
				),
				Entry(
					nil, "at least 1 entry invalid", "return int enum slice",
					[]string{"blah", "json", "json"}, []OutputFormatEnum{OutputFormatEnum(0), JsonFormatEn, JsonFormatEn},
				),
			)

			DescribeTable("AllAreValid",
				func(given, should string, source []string, expect bool) {
					outputFormatEnumSlice.Source = source

					Expect(outputFormatEnumSlice.AllAreValid()).To(Equal(expect))
				},
				func(given, should string, source []string, expect bool) string {
					return fmt.Sprintf("🧪 --> 🍈 given: '%v', should: '%v'",
						given, should)
				},
				Entry(
					nil, "all valid unique values", "return true",
					[]string{"json", "text", "xml"}, true,
				),
				Entry(
					nil, "all valid same values", "true",
					[]string{"json", "json", "json"}, true,
				),
				Entry(
					nil, "at least 1 entry invalid", "return false",
					[]string{"blah", "json", "json"}, false,
				),
			)

			DescribeTable("AllAreValidOrEmpty",
				func(given, should string, source []string, expect bool) {
					outputFormatEnumSlice.Source = source

					Expect(outputFormatEnumSlice.AllAreValidOrEmpty()).To(Equal(expect))
				},
				func(given, should string, source []string, expect bool) string {
					return fmt.Sprintf("🧪 --> 🍈 given: '%v', should: '%v'",
						given, should)
				},
				Entry(
					nil, "all valid unique values", "return true",
					[]string{"json", "text", "xml"}, true,
				),
				Entry(
					nil, "all valid same values", "true",
					[]string{"json", "json", "json"}, true,
				),
				Entry(
					nil, "at least 1 entry invalid", "return false",
					[]string{"blah", "json", "json"}, false,
				),
				Entry(
					nil, "at least 1 entry is empty", "return false",
					[]string{"", "json", "json"}, true,
				),
			)
		})

		Context("EnumValue", func() {
			Context("given: int based enum type", func() {
				It("🧪 should: populate member of native parameter set", func() {
					outputFormatEnum := outputFormatEnumInfo.NewValue()

					paramSet.BindEnum(
						adapters.NewFlagInfo("format", "f", "text"),
						&outputFormatEnum.Source,
					)

					testhelpers.ExecuteCommand(
						rootCommand, "widget", "/usr/fuse/home/music", "--format=xml",
					)

					// This is the line of code that ideally needs to be executed automatically
					// somehow, after the command line has been parsed, possibly command.PreRun
					// rebind/rebinder/bindnative/ for enum fields? Without this automation, the
					// client needs to do this manually, but a single line of code is hardly
					// burdensome, given that any automated scheme would probably not be
					// as efficient, probably requiring more than a single line of code anyway.
					// The documentation of BindEnum does in fact instruct the reader to do so.
					//
					paramSet.Native.Format = outputFormatEnum.Value()
					Expect(paramSet.Native.Format).To(Equal(XmlFormatEn))
					Expect(outputFormatEnum.String()).To(Equal("xml"))
					Expect(outputFormatEnum.IsValid()).To(BeTrue())
				})
			})
		})

		Context("IsValid Enum Validator example via EnumInfo", func() {
			It("🧪 should: invoke custom enum validator", func() {
				outputFormatEnum := outputFormatEnumInfo.NewValue()

				wrapper := paramSet.BindValidatedEnum(
					adapters.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source,
					func(value string) error {
						return lo.Ternary(outputFormatEnumInfo.IsValid(value), nil,
							fmt.Errorf("Enum value: '%v' is not valid", value))
					},
				)
				outputFormatEnum.Source = "xml"
				Expect(wrapper.Validate()).To(BeNil())
			})
		})

		Context("IsValid Enum Validator example via EnumValue", func() {
			It("🧪 should: invoke custom enum validator", func() {
				outputFormatEnum := outputFormatEnumInfo.NewValue()

				wrapper := paramSet.BindValidatedEnum(
					adapters.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source,
					func(value string) error {
						return lo.Ternary(outputFormatEnum.IsValid(), nil,
							fmt.Errorf("Enum value: '%v' is not valid", value))
					},
				)
				outputFormatEnum.Source = "blah"
				Expect(wrapper.Validate()).Error().NotTo(BeNil())
			})
		})
	})
})
