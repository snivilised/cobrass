package adapters

import (
	"errors"
	"fmt"

	"github.com/snivilised/cobrass/src/utils"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

const RootName = "root"

type CobraCommandSpec struct {
	Command *cobra.Command
}

type CobraContainer struct {
	root     *cobra.Command
	commands map[string]*cobra.Command
}

func NewCobraContainer(root *cobra.Command) *CobraContainer {
	return &CobraContainer{
		root:     root,
		commands: make(map[string]*cobra.Command),
	}
}

func (container *CobraContainer) insert(command *cobra.Command) error {
	name := command.Name()
	if _, exists := container.commands[name]; exists {
		message := fmt.Sprintf("cobra container: command '%v' already registered", name)
		return errors.New(message)
	}

	container.commands[name] = command

	return nil
}

func (container *CobraContainer) RegisterCommand(parent string, command *cobra.Command) error {

	if pc := container.Command(parent); pc != nil {
		if err := container.insert(command); err != nil {
			return err
		}
		pc.AddCommand(command)

		return nil
	}
	message := fmt.Sprintf("cobra container: parent command '%v' not registered", parent)
	return errors.New(message)
}

func (container *CobraContainer) RegisterCommands(
	parent string, specs ...*CobraCommandSpec,
) error {

	for _, spec := range specs {
		if err := container.RegisterCommand(parent, spec.Command); err != nil {
			return err
		}
	}

	return nil
}

func (container *CobraContainer) RegisterRootChildCommand(
	command *cobra.Command,
) error {
	return container.RegisterCommand("root", command)
}

func (container *CobraContainer) IsPresent(name string) bool {
	_, exists := container.commands[name]
	return exists
}

func (container *CobraContainer) Root() *cobra.Command {
	return container.root
}

func (container *CobraContainer) Command(name string) *cobra.Command {
	if name == RootName {
		return container.Root()
	}
	command, exists := container.commands[name]

	return utils.TernaryIf(exists, command, nil)
}

func IsChild(parent *cobra.Command, child *cobra.Command) bool {
	_, exists := slices.BinarySearchFunc(parent.Commands(), child, func(a, b *cobra.Command) int {

		return utils.TernaryIfIf(a.Name() == b.Name(), 0, a.Name() < b.Name(), -1, 1)
	})

	return exists
}
