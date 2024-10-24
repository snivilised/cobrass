package gola_test

import (
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok

	"github.com/snivilised/cobrass/generators/gola"
	nef "github.com/snivilised/nefilim"
)

func Path(parent, relative string) string {
	segments := strings.Split(relative, "/")
	return filepath.Join(append([]string{parent}, segments...)...)
}

func Repo(relative string) string {
	_, filename, _, _ := runtime.Caller(0) //nolint:dogsled // use of 3 _ is out of our control
	return Path(filepath.Dir(filename), relative)
}

var _ = Describe("SourceCodeData", Ordered, func() {
	var (
		repo, testPath, sourcePath string
		fS                         nef.UniversalFS
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
		fS = nef.NewUniversalABS()

		_ = testPath
	})

	Context("AnyMissing", func() {
		When("source mode", func() {
			XIt("🧪 should: find all source code files are present", func() {
				outputPath := filepath.Join(repo, sourcePath)
				templatesSubPath := ""
				sourceContainer := gola.NewSourceCodeContainer(fS, outputPath, templatesSubPath)

				Expect(sourceContainer).To(ContainAllSourceCodeFilesAt(
					fS, outputPath,
				))
			})
		})
	})
})
