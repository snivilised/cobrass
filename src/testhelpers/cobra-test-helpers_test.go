package testhelpers

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func emptyRun(*cobra.Command, []string) {}

const onetwo = "one two"

func TestSingleCommand(t *testing.T) {
	var rootCmdArgs []string
	rootCmd := &cobra.Command{
		Use:  "root",
		Args: cobra.ExactArgs(2),
		Run:  func(_ *cobra.Command, args []string) { rootCmdArgs = args },
	}
	aCmd := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: emptyRun}
	bCmd := &cobra.Command{Use: "b", Args: cobra.NoArgs, Run: emptyRun}
	rootCmd.AddCommand(aCmd, bCmd)

	output, err := ExecuteCommand(rootCmd, "one", "two")
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	got := strings.Join(rootCmdArgs, " ")
	if got != onetwo {
		t.Errorf("rootCmdArgs expected: %q, got: %q", onetwo, got)
	}
}
