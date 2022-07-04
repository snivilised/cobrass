package adapters_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
	"github.com/snivilised/cobrass/src/testhelpers"
)

var _ = Describe("Enum", func() {
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *adapters.ParamSet[WidgetParameterSet]
	Context("given: enum type", func() {
		var OutputFormatEnumInfo *adapters.EnumInfo[OutputFormatEnum]

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

			OutputFormatEnumInfo = adapters.NewEnumInfo(adapters.AcceptableEnumValues[OutputFormatEnum]{
				XmlFormatEn:      []string{"xml", "x"},
				JsonFormatEn:     []string{"json", "j"},
				TextFormatEn:     []string{"text", "tx"},
				ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
			})
		})

		It("🧪 should: create enum info", func() {
			Expect(OutputFormatEnumInfo.En("x")).To(Equal(XmlFormatEn))
			Expect(OutputFormatEnumInfo.En("xml")).To(Equal(XmlFormatEn))

			Expect(OutputFormatEnumInfo.En("j")).To(Equal(JsonFormatEn))
			Expect(OutputFormatEnumInfo.En("json")).To(Equal(JsonFormatEn))

			Expect(OutputFormatEnumInfo.NameOf(XmlFormatEn)).To(Equal("xml"))
		})

		Context("IsValid", func() {
			When("given: valid value", func() {
				It("🧪 should: return true", func() {
					Expect(OutputFormatEnumInfo.IsValid("text")).To(BeTrue())
				})
			})

			When("given: invalid value", func() {
				It("🧪 should: return false", func() {
					Expect(OutputFormatEnumInfo.IsValid("foo")).To(BeFalse())
				})
			})
		})

		Context("IsValidOrEmpty", func() {
			When("given: empty value", func() {
				It("🧪 should: return true", func() {
					outputFormatEnum := OutputFormatEnumInfo.NewValue()
					Expect(outputFormatEnum.IsValidOrEmpty()).To(BeTrue())
				})
			})
		})

		Context("given: int based enum type", func() {
			It("🧪 should: populate member of native parameter set", func() {
				outputFormatEnum := OutputFormatEnumInfo.NewValue()

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

		Context("EnumInfo.String", func() {
			It("🧪 should: return contents of enum info", func() {
				result := OutputFormatEnumInfo.String()
				GinkgoWriter.Println("===> contents of OutputFormatEnumInfo:")
				GinkgoWriter.Println(result)
			})
		})
	})
})
