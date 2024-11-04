package gola_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok

	"github.com/snivilised/cobrass/src/generators/gola"
	nef "github.com/snivilised/nefilim"
	"github.com/snivilised/nefilim/test/luna"
)

var _ = Describe("SourceCodeGenerator", Ordered, func() {
	var (
		repo, testPath string
		fS             nef.UniversalFS
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("src", "generators", "gola", "out", "assistant")
		luna.NewMemFS()
		fS = luna.NewMemFS()
	})

	Context("AnyMissing", func() {
		When("test mode", func() {
			It("should: find all source code files are present", func() {
				outputPath := filepath.Join(repo, testPath)
				templatesSubPath := ""
				codeContainer := gola.NewSourceCodeContainer(fS, outputPath, templatesSubPath)

				omitWrite := false
				_, err := codeContainer.Generator(omitWrite).Run()
				Expect(err).Error().To(BeNil())
			})
		})
	})
})
