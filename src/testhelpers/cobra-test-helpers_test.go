package testhelpers_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/src/testhelpers"
	"github.com/spf13/cobra"
)

func emptyRun(*cobra.Command, []string) {}

const onetwo = "one two"

var _ = Describe("CobraTestHelpers", func() {

	Context("given: ", func() {
		It("๐งช should: ", func() {

			var rootCmdArgs []string
			rootCmd := &cobra.Command{
				Use:  "root",
				Args: cobra.ExactArgs(2),
				Run:  func(_ *cobra.Command, args []string) { rootCmdArgs = args },
			}
			aCmd := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: emptyRun}
			bCmd := &cobra.Command{Use: "b", Args: cobra.NoArgs, Run: emptyRun}
			rootCmd.AddCommand(aCmd, bCmd)

			output, err := testhelpers.ExecuteCommand(rootCmd, "one", "two")

			Expect(output).To(Equal(""), fmt.Sprintf(
				"โ Unexpected output '%v'", output,
			))

			Expect(err).To(BeNil(), fmt.Sprintf(
				"โ Unexpected error '%v'", err,
			))

			Expect(strings.Join(rootCmdArgs, " ")).To(Equal(onetwo))
		})
	})
})
