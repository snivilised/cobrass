package assistant_test

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/internal/lab"
	"github.com/snivilised/cobrass/src/internal/third/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var _ = Describe("Enum", func() {
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *assistant.ParamSet[WidgetParameterSet]
	Context("given: enum type", func() {
		var outputFormatEnumInfo *assistant.EnumInfo[OutputFormatEnum]
		var outputFormatEnumSlice assistant.EnumSlice[OutputFormatEnum]

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
			paramSet = assistant.NewParamSet[WidgetParameterSet](widgetCommand)

			outputFormatEnumInfo = assistant.NewEnumInfo(AcceptableOutputFormats)
			outputFormatEnumSlice = outputFormatEnumInfo.NewSlice()
		})

		It("🧪 should: create enum info", func() {
			Expect(outputFormatEnumInfo.En("x")).To(Equal(XMLFormatEn))
			Expect(outputFormatEnumInfo.En("xml")).To(Equal(XMLFormatEn))

			Expect(outputFormatEnumInfo.En("j")).To(Equal(JSONFormatEn))
			Expect(outputFormatEnumInfo.En("json")).To(Equal(JSONFormatEn))

			Expect(outputFormatEnumInfo.NameOf(XMLFormatEn)).To(Equal("xml"))
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

			Context("EnumInfo.AcceptablePrimes", func() {
				When("primaryOnly is true", func() {
					It("🧪 should: return contents of enum info", func() {
						// NB: We have to use this convoluted way of checking
						// the result of AcceptablePrimes, because the result
						// is not being sorted due to an obscure issue:
						//
						// type func(a E, b E) bool of func(a E, b E) bool {…}
						// does not match inferred type:
						// func(a E, b E) int for func(a E, b E) int
						//
						expected := []string{"scribble", "text", "json", "xml"}
						result := outputFormatEnumInfo.AcceptablePrimes()
						elements := strings.Split(result, "/")
						for _, e := range elements {
							if e != "" {
								Expect(slices.Contains(expected, e)).To(BeTrue())
							}
						}

						GinkgoWriter.Println("===> display contents of OutputFormatEnumInfo:")
						GinkgoWriter.Println(result)
					})
				})

				When("primaryOnly is false", func() {
					It("🧪 should: return contents of enum info", func() {
						result := outputFormatEnumInfo.Acceptable()
						Expect(result).To(Equal("//j/json/scr/scribble/scribbler/text/tx/x/xml//"))
						GinkgoWriter.Println("===> display contents of OutputFormatEnumInfo:")
						GinkgoWriter.Println(result)
					})
				})
			})

			When("given: duplicated enum values in acceptables", func() {
				It("🧪 should: panic", func() {
					defer func() {
						_ = recover()
					}()

					invalidAcceptables := assistant.AcceptableEnumValues[OutputFormatEnum]{
						XMLFormatEn:      []string{"xml", "x"},
						JSONFormatEn:     []string{"json", "j"},
						TextFormatEn:     []string{"text", "tx", "x"},
						ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
					}

					assistant.NewEnumInfo(invalidAcceptables)
					Fail("❌ expected panic due to duplicate enum value 'x'")
				})
			})
		})

		Context("EnumSlice", func() {
			DescribeTable("Values",
				func(_, _ string, source []string, expect []OutputFormatEnum) {
					outputFormatEnumSlice.Source = source

					Expect(reflect.DeepEqual(
						outputFormatEnumSlice.Values(), expect,
					)).To(BeTrue())
				},
				func(given, should string, _ []string, _ []OutputFormatEnum) string {
					return fmt.Sprintf("🧪 --> 🍈 given: '%v', should: '%v'",
						given, should)
				},
				Entry(
					nil, "all valid unique values", "return int enum slice",
					[]string{"json", "text", "xml"}, []OutputFormatEnum{JSONFormatEn, TextFormatEn, XMLFormatEn},
				),
				Entry(
					nil, "all valid same values", "return int enum slice",
					[]string{"json", "json", "json"}, []OutputFormatEnum{JSONFormatEn, JSONFormatEn, JSONFormatEn},
				),
				Entry(
					nil, "at least 1 entry invalid", "return int enum slice",
					[]string{"blah", "json", "json"}, []OutputFormatEnum{OutputFormatEnum(0), JSONFormatEn, JSONFormatEn},
				),
			)

			DescribeTable("AllAreValid",
				func(_, _ string, source []string, expect bool) {
					outputFormatEnumSlice.Source = source

					Expect(outputFormatEnumSlice.AllAreValid()).To(Equal(expect))
				},
				func(given, should string, _ []string, _ bool) string {
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
				func(_, _ string, source []string, expect bool) {
					outputFormatEnumSlice.Source = source

					Expect(outputFormatEnumSlice.AllAreValidOrEmpty()).To(Equal(expect))
				},
				func(given, should string, _ []string, _ bool) string {
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
						assistant.NewFlagInfo("format", "f", "text"),
						&outputFormatEnum.Source,
					)

					_, _ = lab.ExecuteCommand(
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
					Expect(paramSet.Native.Format).To(Equal(XMLFormatEn))
					Expect(outputFormatEnum.String()).To(Equal("xml"))
					Expect(outputFormatEnum.IsValid()).To(BeTrue())
				})
			})

			When("given: Source set to an invalid value", func() {
				It("🧪 should: panic", func() {
					defer func() {
						_ = recover()
					}()

					outputFormatEnum := outputFormatEnumInfo.NewValue()
					outputFormatEnum.Source = "foo-bar"
					_ = outputFormatEnum.String()
					Fail("❌ expected panic due to invalid enum value 'foo-bar'")
				})
			})
		})

		Context("IsValid Enum Validator example via EnumInfo", func() {
			It("🧪 should: invoke custom enum validator", func() {
				outputFormatEnum := outputFormatEnumInfo.NewValue()

				wrapper := paramSet.BindValidatedEnum(
					assistant.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source,
					func(value string, _ *pflag.Flag) error {
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
					assistant.NewFlagInfo("format", "f", "xml"),
					&outputFormatEnum.Source,
					func(value string, _ *pflag.Flag) error {
						return lo.Ternary(outputFormatEnum.IsValid(), nil,
							fmt.Errorf("Enum value: '%v' is not valid", value))
					},
				)
				outputFormatEnum.Source = "blah"
				Expect(wrapper.Validate()).Error().NotTo(BeNil())
			})
		})

		Context("Case Sensitivity", func() {
			Context("EnumInfo.IsValid", func() {
				Context("given: string differs in case from acceptables", func() {
					It("🧪 should:  validate ok", func() {
						Expect(outputFormatEnumInfo.IsValid("XML")).To(BeTrue())
					})
				})
			})

			Context("EnumInfo.En", func() {
				Context("given: string differs in case from acceptables", func() {
					It("🧪 should:  validate ok", func() {
						Expect(outputFormatEnumInfo.En("XML")).To(Equal(XMLFormatEn))
					})
				})
			})

			Context("EnumInfo.IsValidOrEmpty", func() {
				Context("given: string differs in case from acceptables", func() {
					It("🧪 should:  validate ok", func() {
						Expect(outputFormatEnumInfo.IsValidOrEmpty("XML")).To(BeTrue())
					})
				})
			})

			Context("EnumValue", func() {
				Context("given: string differs in case from acceptables", func() {
					It("🧪 should:  validate ok", func() {
						outputFormatEnum := outputFormatEnumInfo.NewWith("Scribble")
						Expect(outputFormatEnum.IsValid()).To(BeTrue())
					})
				})
			})
		})
	})
})
