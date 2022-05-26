package adapters

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// see: https://pkg.go.dev/github.com/stretchr/testify/suite
//
type ContainerSuite struct {
	suite.Suite
	Container *CobraContainer

	DummyCommand *cobra.Command
}

func (suite *ContainerSuite) SetupTest() {
	// TODO: investigate why we can't use the new operator instead of
	// using a NewXXX function. new returns a pointer to structure with
	// zero valued members, not initial constructed values.
	//
	suite.Container = NewCobraContainer(&cobra.Command{
		Use:   "root",
		Short: "A root command",
		Long:  "The root command for test case",
	})

	suite.Container.Root().AddCommand(
		&cobra.Command{
			Use:   "child",
			Short: "A child command",
			Long:  "A child command for test case",
		},
	)

	suite.DummyCommand = &cobra.Command{
		Use:   "dummy",
		Short: "A dummy command",
		Long:  "A dummy command for test case",
	}
}

func TestContainerSuite(t *testing.T) {
	suite.Run(t, new(ContainerSuite))
}

func (suite *ContainerSuite) TestContainerCreation() {

	if suite.Container == nil {
		assert.Fail(suite.T(), "Container is nil")
	}
}

func (suite *ContainerSuite) TestRoot() {
	if suite.Container.Root() == nil {
		assert.Fail(suite.T(), "Container Root is nil")
	}
}

func (suite *ContainerSuite) TestCommandForExistingCommand() {
	name := "foo"
	var fooCmd = &cobra.Command{
		Use:   name,
		Short: "Short descrption",
		Long:  "Long description",
		RunE: func(command *cobra.Command, args []string) error {
			return nil
		},
	}
	suite.Container.insert(fooCmd)

	if suite.Container.Command(name) == nil {
		message := fmt.Sprintf(
			"Container should contain previously added command '%v'", name,
		)
		assert.Fail(suite.T(), message)
	}
}

func (suite *ContainerSuite) TestCommandForNonExistingCommand() {
	name := "foo"

	if suite.Container.Command(name) != nil {
		message := fmt.Sprintf(
			"Container should not contain command '%v' that was not previously added", name,
		)
		assert.Fail(suite.T(), message)
	}
}

func (suite *ContainerSuite) TestInsertNonExistingCommand() {

	name := "test"
	var testCmd = &cobra.Command{
		Use:   name,
		Short: "Short descrption",
		Long:  "Long description",
		RunE: func(command *cobra.Command, args []string) error {
			return nil
		},
	}

	suite.Container.insert(testCmd)
	exists := suite.Container.IsPresent(name)
	message := fmt.Sprintf("added command '%v'not found", name)
	assert.True(suite.T(), exists, message)
}

func (suite *ContainerSuite) TestRegisterRootChildCommand() {

	if err := suite.Container.RegisterRootChildCommand(suite.DummyCommand); err != nil {
		message := fmt.Sprintf("failed to register command: '%v' (%v)",
			suite.DummyCommand.Name(), err)
		assert.Fail(suite.T(), message)
	}
}

func (suite *ContainerSuite) TestParentRegisterCommand() {
	parent := "par"

	pc := &cobra.Command{
		Use:   string(parent),
		Short: "A parent command",
		Long:  "A parent command for test case",
	}
	suite.Container.Root().AddCommand(pc)

	suite.Container.insert(pc)

	if err := suite.Container.RegisterCommand(parent, suite.DummyCommand); err != nil {
		message := fmt.Sprintf("failed to register command: '%v' (%v)",
			suite.DummyCommand.Name(), err)
		assert.Fail(suite.T(), message)
	}
	message := fmt.Sprintf("failed to register command: '%v'", suite.DummyCommand.Name())
	assert.True(suite.T(), suite.Container.IsPresent(string(parent)), message)

	message = "Root command should have child commands but doesn't"
	assert.True(suite.T(), pc.HasSubCommands(), message)

	found := IsChild(pc, suite.DummyCommand)

	message = fmt.Sprintf("registered command: '%v' is mssing from the parent: '%v'",
		suite.DummyCommand.Name(), parent)
	assert.True(suite.T(), found, message)
}

func (suite *ContainerSuite) TestParentRegisterCommands() {
	parent := "par"

	pc := &cobra.Command{
		Use:   string(parent),
		Short: "A parent command",
		Long:  "A parent command for test case",
	}
	suite.Container.RegisterRootChildCommand(pc)

	alphaName := "alpha"
	alpha := &CobraCommandSpec{
		Command: &cobra.Command{
			Use:   fmt.Sprintf("%v usage", alphaName),
			Short: "Alpha command",
			Long:  "Alpha child command for test case",
		},
	}

	betaName := "beta"
	beta := &CobraCommandSpec{
		Command: &cobra.Command{
			Use:   fmt.Sprintf("%v usage", betaName),
			Short: "Beta command",
			Long:  "Beta child command for test case",
		},
	}

	deltaName := "delta"
	delta := &CobraCommandSpec{
		Command: &cobra.Command{
			Use:   fmt.Sprintf("%v usage", deltaName),
			Short: "Delta command",
			Long:  "Delta child command for test case",
		},
	}

	specs := []*CobraCommandSpec{alpha, beta, delta}

	if err := suite.Container.RegisterCommands(parent, specs...); err != nil {
		message := fmt.Sprintf("failed to register commands: '%v'", err)
		assert.Fail(suite.T(), message)
	}

	dictionary := map[string]*CobraCommandSpec{
		"alpha": alpha,
		"beta":  beta,
		"delta": delta,
	}

	for name := range dictionary {
		message := fmt.Sprintf("registered command: '%v' is mssing from the parent: '%v'",
			name, string(parent))

		assert.True(suite.T(), IsChild(pc, dictionary[name].Command), message)
	}
}
