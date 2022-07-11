package adapters_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/cobrass/src/adapters"
	"github.com/spf13/cobra"
)

var _ = Describe("CobraContainer", func() {
	var Container *adapters.CobraContainer
	var DummyCommand *cobra.Command
	var ParentCommand *cobra.Command

	BeforeEach(func() {
		Container = adapters.NewCobraContainer(&cobra.Command{
			Use:   "root",
			Short: "A root command",
			Long:  "The root command for test case",
		})

		Container.Root().AddCommand(
			&cobra.Command{
				Use:   "child",
				Short: "A child command",
				Long:  "A child command for test case",
			},
		)

		DummyCommand = &cobra.Command{
			Use:   "dummy",
			Short: "A dummy command",
			Long:  "A dummy command for test case",
		}

		ParentCommand = &cobra.Command{
			Use:   "parent",
			Short: "A parent command",
			Long:  "A parent command for test case",
		}
	})

	Context("Root", func() {
		It("🧪 should: return valid root command", func() {
			Expect(Container.Root()).ToNot(BeNil(), "❌ Container Root is nil")
		})
	})

	Context("Command", func() {
		When("command previously registered", func() {
			It("🧪 should: return registered command", func() {
				name := DummyCommand.Name()
				Container.RegisterRootedCommand(DummyCommand)

				message := fmt.Sprintf(
					"❌ Container should contain previously added command '%v'",
					name,
				)
				Expect(Container.Command(name)).ToNot(BeNil(), message)
			})
		})

		When("command NOT previously registered", func() {
			It("🧪 should: return nil", func() {
				name := "foo"
				message := fmt.Sprintf(
					"❌ Container should not contain command '%v' that was not previously added",
					name,
				)
				Expect(Container.Command(name)).To(BeNil(), message)
			})
		})
	})

	Context("RegisterRootedCommand", func() {
		When("command previously registered", func() {
			It("🧪 should: return error", func() {
				defer func() {
					recover()
				}()
				Container.RegisterRootedCommand(DummyCommand)
				//
				Container.RegisterRootedCommand(DummyCommand)

				Fail("❌ expected panic due to command already registered")
			})
		})

		When("command NOT previously registered", func() {
			It("🧪 should: register the command as child of root command", func() {
				Container.RegisterRootedCommand(DummyCommand)
				Expect(Container.Command(DummyCommand.Name())).ToNot(BeNil())
			})
		})
	})

	Context("RegisterCommand", func() {
		Context("parent previously registered", func() {
			When("requested command previously registered", func() {
				It("🧪 should: return requested command error", func() {
					defer func() {
						recover()
					}()
					parent := ParentCommand.Name()
					Container.RegisterRootedCommand(ParentCommand)
					Container.RegisterCommand(parent, DummyCommand)
					//
					Container.RegisterCommand(parent, DummyCommand)

					Fail("❌ expected panic due to command already registered")
				})
			})

			When("requested command NOT previously registered", func() {
				It("🧪 should: register requesyted command ok", func() {
					name := DummyCommand.Name()
					parent := ParentCommand.Name()
					Container.RegisterRootedCommand(ParentCommand)
					//
					Container.RegisterCommand(parent, DummyCommand)

					message := fmt.Sprintf("❌ parent of '%v' does not match actual parent: '%v'",
						name, parent)
					Expect(DummyCommand.Parent() == ParentCommand).To(BeTrue(), message)
				})
			})
		})

		Context("parent NOT previously registered", func() {
			It("🧪 should: return requested command error", func() {
				defer func() {
					recover()
				}()
				Container.RegisterCommand("foo", DummyCommand)

				Fail("❌ expected panic due to parent not being regsitered")
			})
		})
	})

	Context("RegisterCommands", func() {

		const alphaName = "alpha"
		const betaName = "beta"
		const deltaName = "delta"

		var (
			alpha, beta, delta *adapters.CobraCommandSpec
		)

		BeforeEach(func() {
			alpha = &adapters.CobraCommandSpec{
				Command: &cobra.Command{
					Use:   fmt.Sprintf("%v usage", alphaName),
					Short: "Alpha command",
					Long:  "Alpha child command for test case",
				},
			}

			beta = &adapters.CobraCommandSpec{
				Command: &cobra.Command{
					Use:   fmt.Sprintf("%v usage", betaName),
					Short: "Beta command",
					Long:  "Beta child command for test case",
				},
			}

			delta = &adapters.CobraCommandSpec{
				Command: &cobra.Command{
					Use:   fmt.Sprintf("%v usage", deltaName),
					Short: "Delta command",
					Long:  "Delta child command for test case",
				},
			}
		})

		When("child commands NOT registered", func() {
			It("🧪 should: register all requested commands ok", func() {
				Container.RegisterRootedCommand(ParentCommand)
				parent := ParentCommand.Name()

				specs := []*adapters.CobraCommandSpec{alpha, beta, delta}
				Container.RegisterCommands(parent, specs...)

				message := fmt.Sprintf("❌ parent command: '%v' should have child commands but doesn't",
					parent)

				Expect(ParentCommand.HasSubCommands()).To(BeTrue(), message)

				for _, spec := range specs {
					message = fmt.Sprintf("❌ Parent of command: '%v' should be: '%v'",
						spec.Command.Name(), ParentCommand.Name())

					Expect(spec.Command.Parent() == ParentCommand).To(BeTrue(), message)
				}
			})
		})

		When("when 1 of the commands is already registered", func() {
			It("🧪 should: return err", func() {
				defer func() {
					recover()
				}()
				Container.RegisterRootedCommand(ParentCommand)
				parent := ParentCommand.Name()

				specs := []*adapters.CobraCommandSpec{alpha, beta, beta}
				Container.RegisterCommands(parent, specs...)

				Fail("❌ expected panic due to a command already being registered")
			})
		})
	})

	Context("IsPresent", func() {
		It("🧪 should: return bool indicating command presence", func() {
			Container.IsPresent("child")
		})
	})

	Context("Native", func() {
		When("given: a parameter set name not previously registered", func() {
			It("🧪 should: panic", func() {
				defer func() {
					recover()
				}()
				Container.Native("foo-bar")
				Fail("❌ expected panic")
			})
		})
	})
})
