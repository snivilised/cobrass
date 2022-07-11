package adapters_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/adapters"
)

var _ = Describe("ValidatorContainer", func() {
	var validators *adapters.ValidatorContainer
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *adapters.ParamSet[WidgetParameterSet]

	Context("NewValidatorContainer", func() {
		When("options set to nil", func() {
			It("🧪 should: create ValidatorContainer with default options", func() {
				validators = adapters.NewValidatorContainer(nil)
				Expect(validators).ToNot(BeNil())
			})
		})

		When("options specified", func() {
			It("🧪 should: create ValidatorContainer", func() {
				validators = adapters.NewValidatorContainer(&adapters.ValidatorGroupOptions{
					Size: 10,
				})
				Expect(validators).ToNot(BeNil())
			})
		})
	})

	Context("ValidatorContainer", func() {
		BeforeEach(func() {
			validators = adapters.NewValidatorContainer(nil)

			rootCommand = &cobra.Command{
				Use:   "peek",
				Short: "A brief description of your application",
				Long:  "A long description of the root peek command",
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

		When("validator is added with a flag that already exists", func() {
			It("🧪 should: only panic when duplicate added", func() {
				flag := "pattern"
				validator := paramSet.BindValidatedString(
					adapters.NewFlagInfo(flag, "p", "default-pattern"),
					&paramSet.Native.Pattern,
					func(value string) error {
						return nil
					},
				)

				validators.Add(flag, validator)
				{
					defer func() {
						recover()
					}()
					validators.Add(flag, validator)
					Fail("❌ Expected a panic")
				}
			})
		})
		Context("Run", func() {
			When("a validator fails", func() {
				It("🧪 should: return error", func() {
					wrapper := adapters.StringOptionValidator{
						Fn: func(value string) error {
							return fmt.Errorf("directory does not exist")
						},
						Value: &paramSet.Native.Directory,
					}
					Expect(wrapper.Validate()).Error().NotTo(BeNil())
				})

				It("🧪 should: (via paramset) return error", func() {
					validators.Add("Directory", paramSet.BindValidatedString(
						adapters.NewFlagInfo("directory", "d", "/foo-bar"),
						&paramSet.Native.Directory,
						func(value string) error {
							return fmt.Errorf("directory does not exist")
						},
					))
					Expect(validators.Run()).Error().NotTo(BeNil())
				})
			})

			When("all validators pass", func() {
				It("🧪 should: return nil", func() {
					validators.Add("Directory", adapters.StringOptionValidator{
						Fn: func(value string) error {
							return nil
						},
						Value: &paramSet.Native.Directory,
					})
					validators.Add("Count", adapters.UintOptionValidator{
						Fn: func(value uint) error {
							return nil
						},
						Value: &paramSet.Native.Count,
					})
					Expect(validators.Run()).Error().To(BeNil())
				})
				It("🧪 should: (via paramset) return nil", func() {
					validators.Add("Directory", paramSet.BindValidatedString(
						adapters.NewFlagInfo("directory", "d", "/foo-bar"),
						&paramSet.Native.Directory,
						func(value string) error {
							return nil
						},
					))
					validators.Add("Count", paramSet.BindValidatedUint(
						adapters.NewFlagInfo("count", "c", uint(0)),
						&paramSet.Native.Count,
						func(value uint) error {
							return nil
						},
					))
					Expect(validators.Run()).Error().To(BeNil())
				})
			})
		})

		Context("Get", func() {
			When("validator not found", func() {
				It("🧪 should: return nil value", func() {
					Expect(validators.Get("missing")).To(BeNil())
				})
			})

			When("validator is present", func() {
				It("🧪 should: return the requested validator", func() {
					validators.Add("Directory", paramSet.BindValidatedString(
						adapters.NewFlagInfo("directory", "d", "/foo-bar"),
						&paramSet.Native.Directory,
						func(value string) error {
							return nil
						},
					))
					Expect(validators.Get("Directory")).ToNot(BeNil())
				})
			})
		})
	})
})
