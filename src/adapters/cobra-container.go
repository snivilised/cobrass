package adapters

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/snivilised/cobrass/src/utils"
	"github.com/spf13/cobra"
)

// CobraCommandSpec is a wrapper around the cobra command, require to register
// multiple commands at he same time, see MustRegisterCommands.
//
type CobraCommandSpec struct {
	// Command: a pointer to the underlying cobra command
	//
	Command *cobra.Command
}

type paramSetsCollection map[string]any
type commandsCollection map[string]*cobra.Command

// CobraContainer is a wrapper around the collection of cobra commands.
// Please see unit tests for examples of how to use the CobraContainer.
//
type CobraContainer struct {
	root      *cobra.Command
	commands  commandsCollection
	paramSets paramSetsCollection
}

// NewCobraContainer is a factory function for the CobraContainer. The client
// must pass in the root Cobra command.
//
// - root: the root Cobra command.
//
func NewCobraContainer(root *cobra.Command) *CobraContainer {
	return &CobraContainer{
		root:      root,
		commands:  make(commandsCollection),
		paramSets: make(paramSetsCollection),
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

// MustRegisterCommand stores a command inside the container. The client passes in the
// name of the parent command and the command is added to that parent.
//
// - parent: the name of the parent command. The name can be derived by calling the Name()
// member function of the Cobra command.
//
// - command: the Cobra command to register.
//
// panics if the there is no command currently registered with the name of parent
//
func (container *CobraContainer) MustRegisterCommand(parent string, command *cobra.Command) {

	if pc := container.Command(parent); pc != nil {
		if err := container.insert(command); err != nil {
			panic(err)
		}
		pc.AddCommand(command)
	} else {
		message := fmt.Sprintf("cobra container: parent command '%v' not registered", parent)
		panic(message)
	}
}

// MustRegisterCommands invokes MustRegisterCommand for each command in the list
//
func (container *CobraContainer) MustRegisterCommands(parent string, specs ...*CobraCommandSpec) {

	for _, spec := range specs {
		container.MustRegisterCommand(parent, spec.Command)
	}
}

// MustRegisterRootedCommand stores a command inside the container as a direct descendent
// of the root Cobra command and is added to the root command itself.
//
// - command: the Cobra command to register.
//
// panics if the command with the same name has already been registered.
//
func (container *CobraContainer) MustRegisterRootedCommand(command *cobra.Command) {
	container.MustRegisterCommand(container.root.Name(), command)
}

// IsPresent checks whether a command has been registered anywhere within the
// command tree. NB, the container stores all commands in a flat hierarchy as opposed
// to Cobra which stores commands in a tree like hierarchy.
//
// - name: the name of the command to check.
//
// Returns true if present, false otherwise
//
func (container *CobraContainer) IsPresent(name string) bool {
	_, exists := container.commands[name]
	return exists
}

// Root returns the root command
//
func (container *CobraContainer) Root() *cobra.Command {
	return container.root
}

// Command returns the command registered with the name specified
//
// - name: the name of the Cobra command to check. The name can be derived by
// calling the Name() function on the cobra command.
//
// Returns the command identified by the name, nil if the command does not exist
//
func (container *CobraContainer) Command(name string) *cobra.Command {
	if name == container.root.Name() {
		return container.Root()
	}
	command, exists := container.commands[name]

	return utils.TernaryIf(exists, command, nil)
}

// MustRegisterParamSet stores the parameter set under the provided name. Used
// to reduce the number of floating global variables that the client needs
// to manage when using cobra.
//
// panics if param set already registered, or attempt to register with
// an inappropriate type.
//
func (container *CobraContainer) MustRegisterParamSet(name string, ps any) {

	if _, exists := container.paramSets[name]; exists {
		panic(fmt.Errorf("parameter set '%v' already registered", name))
	}

	if reflect.TypeOf(ps).Kind() != reflect.Ptr {
		typeOf := reflect.TypeOf(ps)

		panic(fmt.Errorf("cant register parameter set '%v' with non pointer type: '%v'",
			name, typeOf))
	}

	if reflect.TypeOf(ps).Elem().Kind() != reflect.Struct {
		typeOf := reflect.TypeOf(ps)

		panic(fmt.Errorf("cant register parameter set '%v' with non struct type: '%v'",
			name, typeOf))
	}

	container.paramSets[name] = ps
}

// Native retrieves the Native parameter set that was previously registered
//
func (container *CobraContainer) Native(name string) any {

	// Need to use reflection to get the Native property. The collection of
	// parameter sets can't be defined as a generic, because collections
	// of generics are homogeneous, but we need a heterogeneous collection of
	// parameter sets. This is why we need to use reflection to get hold of
	// the Native property.
	//
	if paramSet, found := container.paramSets[name]; found {
		paramSetStruct := reflect.ValueOf(paramSet).Elem()

		return paramSetStruct.FieldByName("Native").Interface()
	} else {
		panic(fmt.Errorf("parameter set '%v' not found", name))
	}
}

// MustGetParamSet like Native, except that it returns the parameter set
// wrapper. The client must perform a type assertion on the
// returned pointer to translate it back into the native type,
// ie MustGetParamSet[N] (as opposed to N)
//
func (container *CobraContainer) MustGetParamSet(name string) any {
	if paramSet, found := container.paramSets[name]; found {
		return paramSet
	} else {
		panic(fmt.Errorf("parameter set '%v' not found", name))
	}
}
