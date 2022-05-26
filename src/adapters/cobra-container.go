package adapters

import (
	"errors"
	"fmt"

	"github.com/snivilised/cobrass/src/utils"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

const RootName = "root"

/*
CobraCommandSpec is a wrapper around the cobra command, require to register
multiple commands at he same time, see RegisterCommands.
- Command: a pointer to the underlying cobra command
*/
type CobraCommandSpec struct {
	Command *cobra.Command
}

/*
CobraContainer is a wrapper around the collection of cobra commands.

Please see unit tests for examples of how to use the CobraContainer.
*/
type CobraContainer struct {
	root     *cobra.Command
	commands map[string]*cobra.Command
}

/*
NewCobraContainer is a factory function for the CobraContainer. The client
must pass in the root Cobra command.
- root: the root Cobra command.
*/
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

/*
RegisterCommand stores a command inside the container. The client passes in the
name of the parent command and the command is added to that parent.
- parent: the name of the parent command. The name can be derived by calling the Name()
member function of the Cobra command.
- command: the Cobra command to register.

Returns an error if the there is no command currently registered with the name of parent
*/
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

/*
RegisterRootChildCommand stores a command inside the container as a direct descendent
of the root Cobra command and is added to the root command itself.
- command: the Cobra command to register.

Returns an error if the command with the same name has already been registered.
*/
func (container *CobraContainer) RegisterRootChildCommand(
	command *cobra.Command,
) error {
	return container.RegisterCommand("root", command)
}

/*
IsPresent checks whether a command has been registered anywhere within the
command tree. NB, the container stores all commands in a flat hierarchy as opposed
to Cobra which stores commands in a tree like hierarchy.
- name: the name of the command to check.

Returns true if present, false otherwise
*/
func (container *CobraContainer) IsPresent(name string) bool {
	_, exists := container.commands[name]
	return exists
}

/*
Root returns the root command
*/
func (container *CobraContainer) Root() *cobra.Command {
	return container.root
}

/*
Command return the command registered with the name specified
- name: the name of the Cobra command to check. The name can be derived by
calling the Name() function on the command.

Returns the command identified by the name, nil if the command does not exist
*/
func (container *CobraContainer) Command(name string) *cobra.Command {
	if name == RootName {
		return container.Root()
	}
	command, exists := container.commands[name]

	return utils.TernaryIf(exists, command, nil)
}

/*
IsChild determines if child command is a direct descendent of the parent
- parent: the parent Cobra command
- child: the child Cobra command

Returns true if child is direct descendent of the parent, false otherwise.
*/
func IsChild(parent *cobra.Command, child *cobra.Command) bool {
	_, exists := slices.BinarySearchFunc(parent.Commands(), child, func(a, b *cobra.Command) int {

		return utils.TernaryIfIf(a.Name() == b.Name(), 0, a.Name() < b.Name(), -1, 1)
	})

	return exists
}
