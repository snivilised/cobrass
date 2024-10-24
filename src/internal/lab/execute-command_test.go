package lab_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
	"github.com/spf13/cobra"

	"github.com/snivilised/cobrass/src/internal/lab"
)

func emptyRun(*cobra.Command, []string) {}

const oneTwo = "one two"

var _ = Describe("Execute", func() {
	// strictly speaking, we don't need this test because its testing
	// test code. But it is of value here because it illustrates
	// how to write cobra tests
	//
	It("🧪 should: execute", func() {
		var rootCmdArgs []string
		rootCmd := &cobra.Command{
			Use:  "root",
			Args: cobra.ExactArgs(2),
			Run:  func(_ *cobra.Command, args []string) { rootCmdArgs = args },
		}
		aCmd := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: emptyRun}
		bCmd := &cobra.Command{Use: "b", Args: cobra.NoArgs, Run: emptyRun}
		rootCmd.AddCommand(aCmd, bCmd)
		output, err := lab.ExecuteCommand(rootCmd, "one", "two")

		Expect(output).To(Equal(""), fmt.Sprintf(
			"❌ Unexpected output '%v'", output,
		))

		Expect(err).To(BeNil(), fmt.Sprintf(
			"❌ Unexpected error '%v'", err,
		))

		Expect(strings.Join(rootCmdArgs, " ")).To(Equal(oneTwo))
	})
})
