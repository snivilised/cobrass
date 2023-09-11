package gola_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/generators/gola"
)

var _ = Describe("SourceCodeGenerator", Ordered, func() {
	var (
		repo, testPath, sourcePath string
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
		_ = testPath
		_ = sourcePath
		_ = repo
	})

	Context("AnyMissing", func() {
		When("test mode", func() {
			It("should: find all source code files are present", func() {
				outputPath := filepath.Join(repo, testPath)
				templatesSubPath := ""
				sourceCode := gola.NewSourceCodeContainer(outputPath, templatesSubPath)

				omitWrite := false
				err := sourceCode.Generator(omitWrite).Run()
				Expect(err).Error().To(BeNil())
			})
		})
	})
})
