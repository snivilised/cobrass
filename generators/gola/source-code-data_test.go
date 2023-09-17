package gola_test

import (
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/generators/gola"
	"github.com/snivilised/cobrass/generators/gola/internal/storage"
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
		fs                         storage.VirtualFS
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
		fs = storage.UseNativeFS()

		_ = testPath
	})

	Context("AnyMissing", func() {
		When("source mode", func() {
			It("🧪 should: find all source code files are present", func() {
				outputPath := filepath.Join(repo, sourcePath)
				templatesSubPath := ""
				sourceContainer := gola.NewSourceCodeContainer(fs, outputPath, templatesSubPath)

				Expect(sourceContainer).To(ContainAllSourceCodeFilesAt(
					fs, outputPath,
				))
			})
		})
	})
})
