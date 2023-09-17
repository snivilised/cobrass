package gola_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/generators/gola"
	"github.com/snivilised/cobrass/generators/gola/internal/storage"
)

var _ = Describe("SourceCodeGenerator", Ordered, func() {
	var (
		repo, testPath, sourcePath string
		fs                         storage.VirtualFS
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
		fs = storage.UseMemFS()

		_ = sourcePath
	})

	Context("AnyMissing", func() {
		When("test mode", func() {
			It("should: find all source code files are present", func() {
				outputPath := filepath.Join(repo, testPath)
				templatesSubPath := ""
				codeContainer := gola.NewSourceCodeContainer(fs, outputPath, templatesSubPath)

				omitWrite := false
				_, err := codeContainer.Generator(omitWrite).Run()
				Expect(err).Error().To(BeNil())
			})
		})
	})
})
