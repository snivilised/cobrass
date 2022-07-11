package adapters_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
	"github.com/snivilised/cobrass/src/testhelpers"
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
		var paramSet *adapters.ParamSet[WidgetParameterSet]

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
			rootCommand.AddCommand(widgetCommand)

			paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)
		})

		// These are binder based tests that have a characteristic that can't be accommodated easily
		// in the auto generated tests and hence easier just to right the test explicitly.
		//
		DescribeTable("binder",
			func(entry TcEntry) {
				entry.Binder()

				GinkgoWriter.Printf("🍧🍧🍧 ABOUT TO RUN ...\n")
				testhelpers.ExecuteCommand(
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
						adapters.NewFlagInfo("concise ensures that output is compressed", "c", false),
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
						recover()
					}()
					adapters.NewParamSet[InvalidParameterSet](widgetCommand)

					Fail("❌ expected panic due to attempt to create a param set with a non struct")
				})
			})
		})

		Context("ParamSet.Validators", func() {
			It("🧪 should: get validators", func() {
				Expect(paramSet.Validators()).NotTo(BeNil())
			})
		})

		Context("Register ParamSet", func() {
			var container *adapters.CobraContainer
			const cname = "widget"
			const psname = cname + "-ps"

			BeforeEach(func() {
				container = adapters.NewCobraContainer(rootCommand)
			})

			It("🧪 should: be able get registered param set", func() {
				const cname = "widget"
				const psname = cname + "-ps"
				container.RegisterParamSet(psname, paramSet)

				testhelpers.ExecuteCommand(
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
						recover()
					}()

					container.RegisterParamSet(psname, paramSet)
					container.RegisterParamSet(psname, paramSet)
					Fail("❌ expected panic due to attempt to register parameter that already exists")
				})
			})

			When("parameter set registered with non pointer type", func() {
				It("🧪 should: panic", func() {
					defer func() {
						recover()
					}()
					const foo = "bar"
					container.RegisterParamSet(psname, foo)
					Fail("❌ expected panic due to attempt to register non pointer type")
				})
			})

			When("parameter set registered with non struct", func() {
				It("🧪 should: panic", func() {
					defer func() {
						recover()
					}()
					foo := 42
					container.RegisterParamSet(psname, &foo)
					Fail("❌ expected panic due to attempt to register pointer to non struct")
				})
			})
		})

		Context("NewFlagInfoOnFlagSet", func() {
			It("🧪 should: bind flag to alternative flag set", func() {
				paramSet.BindString(
					adapters.NewFlagInfoOnFlagSet("pattern", "p", "default-pattern",
						widgetCommand.PersistentFlags()), &paramSet.Native.Pattern,
				)
				commandLine := "--pattern=*music.infex*"
				testhelpers.ExecuteCommand(
					rootCommand, "widget", "/usr/fuse/home/music", commandLine,
				)
				Expect(paramSet.Native.Pattern).To(Equal("*music.infex*"))
			})
		})
	})
})
