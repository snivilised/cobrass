package assistant_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/snivilised/cobrass/src/assistant"
)

var _ = Describe("ValidatorContainer", func() {
	var validators *assistant.ValidatorContainer
	var rootCommand *cobra.Command
	var widgetCommand *cobra.Command
	var paramSet *assistant.ParamSet[WidgetParameterSet]

	Context("NewValidatorContainer", func() {
		When("options not specified", func() {
			It("🧪 should: create ValidatorContainer with default options", func() {
				validators = assistant.NewValidatorContainer()
				Expect(validators).ToNot(BeNil())
			})
		})

		When("options specified", func() {
			It("🧪 should: create ValidatorContainer", func() {
				validators = assistant.NewValidatorContainer(func(o *assistant.ValidatorContainerOptions) {
					o.Size = 10
				})
				Expect(validators).ToNot(BeNil())
			})
		})
	})

	Context("ValidatorContainer", func() {
		BeforeEach(func() {
			validators = assistant.NewValidatorContainer()

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
				RunE: func(_ *cobra.Command, args []string) error {
					paramSet.Native.Directory = args[0]
					return nil
				},
			}
			rootCommand.AddCommand(widgetCommand)
			paramSet = assistant.NewParamSet[WidgetParameterSet](widgetCommand)

		})

		When("validator is added with a flag that already exists", func() {
			It("🧪 should: only panic when duplicate added", func() {
				flag := "pattern"
				validator := paramSet.BindValidatedString(
					assistant.NewFlagInfo(flag, "p", "default-pattern"),
					&paramSet.Native.Pattern,
					func(_ string, _ *pflag.Flag) error {
						return nil
					},
				)

				validators.Add(flag, validator)
				{
					defer func() {
						_ = recover()
					}()
					validators.Add(flag, validator)
					Fail("❌ Expected a panic")
				}
			})
		})
		Context("Run", func() {
			When("a validator fails", func() {
				It("🧪 should: return error", func() {
					wrapper := assistant.StringOptionValidator{
						Fn: func(_ string, _ *pflag.Flag) error {
							return fmt.Errorf("directory does not exist")
						},
						Value: &paramSet.Native.Directory,
					}
					Expect(wrapper.Validate()).Error().NotTo(BeNil())
				})

				It("🧪 should: (via paramset) return error", func() {
					validators.Add("Directory", paramSet.BindValidatedString(
						assistant.NewFlagInfo("directory", "d", "/foo-bar"),
						&paramSet.Native.Directory,
						func(_ string, _ *pflag.Flag) error {
							return fmt.Errorf("directory does not exist")
						},
					))
					Expect(paramSet.Validate()).Error().NotTo(BeNil())
				})
			})

			When("all validators pass", func() {
				It("🧪 should: return nil", func() {
					validators.Add("Directory", assistant.StringOptionValidator{
						Fn: func(_ string, _ *pflag.Flag) error {
							return nil
						},
						Value: &paramSet.Native.Directory,
					})
					validators.Add("Count", assistant.UintOptionValidator{
						Fn: func(_ uint, _ *pflag.Flag) error {
							return nil
						},
						Value: &paramSet.Native.Count,
					})
					Expect(paramSet.Validate()).Error().To(BeNil())
				})
				It("🧪 should: (via paramset) return nil", func() {
					validators.Add("Directory", paramSet.BindValidatedString(
						assistant.NewFlagInfo("directory", "d", "/foo-bar"),
						&paramSet.Native.Directory,
						func(_ string, _ *pflag.Flag) error {
							return nil
						},
					))
					validators.Add("Count", paramSet.BindValidatedUint(
						assistant.NewFlagInfo("count", "c", uint(0)),
						&paramSet.Native.Count,
						func(_ uint, _ *pflag.Flag) error {
							return nil
						},
					))
					Expect(paramSet.Validate()).Error().To(BeNil())
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
						assistant.NewFlagInfo("directory", "d", "/foo-bar"),
						&paramSet.Native.Directory,
						func(_ string, _ *pflag.Flag) error {
							return nil
						},
					))
					Expect(validators.Get("Directory")).ToNot(BeNil())
				})
			})
		})
	})
})
