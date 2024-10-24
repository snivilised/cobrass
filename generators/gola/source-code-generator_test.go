package gola_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok

	"github.com/snivilised/cobrass/generators/gola"
	nef "github.com/snivilised/nefilim"
)

var _ = Describe("SourceCodeGenerator", Ordered, func() {
	var (
		repo, testPath, sourcePath string
		fs                         nef.UniversalFS
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
		fs = nef.NewUniversalABS() // use mapFS

		_ = sourcePath
	})

	Context("AnyMissing", func() {
		When("test mode", func() {
			XIt("should: find all source code files are present", func() {
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
