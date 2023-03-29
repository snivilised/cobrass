package assistant_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/internal/helpers"
)

type TcEntry struct {
	Message     string
	Binder      func()
	CommandLine string
	Assert      func()
}

var _ = Describe("ParamSet (manual)", func() {

	When("Binding a flag (manual)", func() {
		var rootCommand *cobra.Command
		var widgetCommand *cobra.Command
		var paramSet *assistant.ParamSet[WidgetParameterSet]
		var outputFormatEnumInfo *assistant.EnumInfo[OutputFormatEnum]
		var outputFormatEnum assistant.EnumValue[OutputFormatEnum]
		var container *assistant.CobraContainer

		const cname = "widget"
		const psname = cname + "-ps"

		BeforeEach(func() {
			rootCommand = &cobra.Command{
				Use:   "poke",
				Short: "A brief description of your application",
				Long:  "A long description of the root poke command",
			}

			widgetCommand = &cobra.Command{
				Version: "1.0.1",
				Use:     "widget",
				Short:   "Create widget",
				Long:    "Index file system at root: '/'",
				Args:    cobra.ExactArgs(1),

				PreRun: func(command *cobra.Command, args []string) {
					GinkgoWriter.Printf("**** 🍉 PRE-RUN\n")
				},
				RunE: func(command *cobra.Command, args []string) error {
					GinkgoWriter.Printf("===> 🍓 EXECUTE (Directory: '%v')\n", args[0])

					paramSet.Native.Directory = args[0]
					return nil
				},
				PostRun: func(command *cobra.Command, args []string) {
					GinkgoWriter.Printf("**** 🍒 POST-RUN\n")
				},
			}
			container = assistant.NewCobraContainer(rootCommand)
			container.MustRegisterRootedCommand(widgetCommand)

			paramSet = assistant.NewParamSet[WidgetParameterSet](widgetCommand)
			outputFormatEnumInfo = assistant.NewEnumInfo(AcceptableOutputFormats)
			outputFormatEnum = outputFormatEnumInfo.NewValue()
		})

		// These are binder based tests that have a characteristic that can't be accommodated easily
		// in the auto generated tests and hence easier just to right the test explicitly.
		//
		DescribeTable("binder",
			func(entry TcEntry) {
				entry.Binder()

				GinkgoWriter.Printf("🍧🍧🍧 ABOUT TO RUN ...\n")
				_, _ = helpers.ExecuteCommand(
					rootCommand, "widget", "/usr/fuse/home/music", entry.CommandLine,
				)
				GinkgoWriter.Printf("🍧🍧🍧 AFTER RUN\n")

				entry.Assert()
			},

			func(entry TcEntry) string {
				return fmt.Sprintf("🧪 --> 🍒 given: flag is '%v'", entry.Message)
			},

			// special scenario, not auto generated
			//
			Entry(nil, TcEntry{
				Message: "bool type flag is NOT present",
				Binder: func() {
					paramSet.BindBool(
						assistant.NewFlagInfo("concise ensures that output is compressed", "c", false),
						&paramSet.Native.Concise,
					)
				},
				CommandLine: "",
				Assert:      func() { Expect(paramSet.Native.Concise).To(BeFalse()) },
			}),
		)

		Context("NewParamSet", func() {
			When("given: non struct native param set type", func() {
				It("🧪 should: panic", func() {
					defer func() {
						_ = recover()
					}()
					assistant.NewParamSet[InvalidParameterSet](widgetCommand)

					Fail("❌ expected panic due to attempt to create a param set with a non struct")
				})
			})
		})

		Context("ParamSet.Validators", func() {
			It("🧪 should: get validators", func() {
				Expect(paramSet.Validators()).NotTo(BeNil())
			})
		})

		Context("ParamSet.CrossValidate", func() {
			When("given: a passing param set", func() {
				It("🧪 should: return no error", func() {
					paramSet.BindEnum(
						assistant.NewFlagInfo("format", "f", "xml"),
						&outputFormatEnum.Source,
					)

					paramSet.BindString(
						assistant.NewFlagInfo("pattern", "p", "cakewalk"),
						&paramSet.Native.Pattern,
					)
					container.MustRegisterParamSet(psname, paramSet)

					commandLine := "--format xml --pattern cakewalk"
					_, _ = helpers.ExecuteCommand(
						rootCommand, "widget", "/usr/fuse/home/music", commandLine,
					)
					paramSet.Native.Format = outputFormatEnum.Value()

					result := paramSet.CrossValidate(func(ps *WidgetParameterSet) error {
						condition := (ps.Format == XMLFormatEn) && strings.Contains(ps.Pattern, "walk")
						return lo.Ternary(condition, nil,
							fmt.Errorf("invalid combination, pattern: '%v'", ps.Pattern))
					})
					Expect(result).Error().To(BeNil())
				})
			})

			When("given: an invalid param set", func() {
				It("🧪 should: return error", func() {
					paramSet.BindEnum(
						assistant.NewFlagInfo("format", "f", "xml"),
						&outputFormatEnum.Source,
					)

					paramSet.BindString(
						assistant.NewFlagInfo("pattern", "p", "cakewalk"),
						&paramSet.Native.Pattern,
					)
					container.MustRegisterParamSet(psname, paramSet)

					commandLine := "--format xml --pattern cakewalk"
					_, _ = helpers.ExecuteCommand(
						rootCommand, "widget", "/usr/fuse/home/music", commandLine,
					)
					paramSet.Native.Format = outputFormatEnum.Value()

					result := paramSet.CrossValidate(func(ps *WidgetParameterSet) error {
						condition := (ps.Format == XMLFormatEn) && strings.Contains(ps.Pattern, "foobar")
						return lo.Ternary(condition, nil,
							fmt.Errorf("invalid combination, pattern: '%v'", ps.Pattern))
					})
					Expect(result).Error().ToNot(BeNil())
				})
			})
		})

		Context("Register ParamSet", func() {

			It("🧪 should: be able get registered param set", func() {
				container.MustRegisterParamSet(psname, paramSet)

				_, _ = helpers.ExecuteCommand(
					rootCommand, cname, "/usr/fuse/home/cache",
				)

				if native, ok := container.Native(psname).(*WidgetParameterSet); ok {
					Expect(native.Directory).To(Equal("/usr/fuse/home/cache"))
				} else {
					Fail("❌ param set should be registered")
				}
			})

			When("parameter set exists", func() {
				It("🧪 should: panic", func() {
					defer func() {
						_ = recover()
					}()

					container.MustRegisterParamSet(psname, paramSet)
					container.MustRegisterParamSet(psname, paramSet)
					Fail("❌ expected panic due to attempt to register parameter that already exists")
				})
			})

			When("parameter set registered with non pointer type", func() {
				It("🧪 should: panic", func() {
					defer func() {
						_ = recover()
					}()
					const foo = "bar"
					container.MustRegisterParamSet(psname, foo)
					Fail("❌ expected panic due to attempt to register non pointer type")
				})
			})

			When("parameter set registered with non struct", func() {
				It("🧪 should: panic", func() {
					defer func() {
						_ = recover()
					}()
					foo := 42
					container.MustRegisterParamSet(psname, &foo)
					Fail("❌ expected panic due to attempt to register pointer to non struct")
				})
			})

			Context("ParamSet", func() {
				When("param set is registered", func() {
					It("🧪 should: return the parameter wrapper", func() {

						container.MustRegisterParamSet(psname, paramSet)
						resultPS := container.MustGetParamSet(psname).(*assistant.ParamSet[WidgetParameterSet])

						Expect(resultPS).ToNot(BeNil())
					})
				})

				When("param set is NOT registered", func() {
					It("🧪 should: panic", func() {
						defer func() {
							_ = recover()
						}()
						_ = container.MustGetParamSet("foo").(*assistant.ParamSet[WidgetParameterSet])

						Fail("❌ expected panic due to parameter set not found")
					})
				})
			})
		})

		Context("NewFlagInfoOnFlagSet", func() {
			It("🧪 should: bind flag to alternative flag set", func() {
				paramSet.BindString(
					assistant.NewFlagInfoOnFlagSet("pattern", "p", "default-pattern",
						widgetCommand.PersistentFlags()), &paramSet.Native.Pattern,
				)
				commandLine := "--pattern=*music.infex*"
				_, _ = helpers.ExecuteCommand(
					rootCommand, "widget", "/usr/fuse/home/music", commandLine,
				)
				Expect(paramSet.Native.Pattern).To(Equal("*music.infex*"))
			})
		})
	})
})
