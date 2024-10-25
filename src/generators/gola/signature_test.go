package gola_test

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok
	. "github.com/onsi/gomega"    //nolint:revive // ok
	"github.com/snivilised/cobrass/src/generators/gola"
	"github.com/snivilised/cobrass/src/internal/lab"
	nef "github.com/snivilised/nefilim"
)

type setupFile struct {
	path string
	data []byte
}

func setup(fS nef.UniversalFS, directoryPath string, files ...setupFile) {
	if e := fS.MakeDirAll(directoryPath, lab.Perms.Dir); e != nil {
		Fail(fmt.Sprintf("%q, path: %q", e.Error(), directoryPath))
	}

	for _, f := range files {
		if e := fS.WriteFile(f.path, f.data, lab.Perms.File); e != nil {
			Fail(fmt.Sprintf("%q, path: %q", e.Error(), f.path))
		}
	}
}

// By using a virtual file system, we can write tests that can read existing
// source from the native file system but write only to an in memory file system
// The cobrass-gen tool, when invoked directly (as opposed to running a test),
// can be used to generate test content to the test location.
// (./generators/gola/out/assistant). The rationale behind this is that when
// development occurs, new code can be generated into the test location without
// having to over-write the source code. Only when the new defined generated code
// is deemed to be correct, the existing code can be overridden.

var _ = Describe("Signature", Ordered, func() {
	var (
		repo,
		testPath,
		sourcePath,
		outputPath string
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("src", "generators", "gola", "out", "assistant")
		sourcePath = "assistant"
	})

	Context("sign", func() {
		When("standalone", func() { // 📍 can't use Generate, only Signature()
			Context("and: Test mode", func() {
				Context("and: without write", func() {
					It("🧪 should: return hash result of newly generated content", func() {
						fS := NewTestMemFS()
						templatesSubPath := ""
						outputPath = filepath.Join(repo, testPath)

						// 👻 TODO: !! what about the other files?; this is short term and will
						// need to accommodate the other source files.
						//
						path := filepath.Join(testPath, "option-validator-auto.go")

						if data, err := os.ReadFile(path); err != nil {
							setup(fS, testPath, setupFile{
								path: path,
								data: data,
							})
						}
						Expect(fS.FileExists(path)).To(BeTrue())

						sourceCode := gola.NewSourceCodeContainer(fS, testPath, templatesSubPath)
						result, err := sourceCode.Signature()

						Expect(err).Error().To(BeNil())
						Expect(result.Hash).NotTo(BeEmpty())

						// 👻 Can't expect this to match yet, since the registered hash is
						// generated from all 3 sources, but only the generation of
						// option-validator-auto.go has been implemented so far
						//
						Expect(result.Hash).NotTo(MatchRegisteredHash(gola.RegisteredHash))
					})
				})
			})

			Context("and: Source mode", func() {
				Context("and: without write", func() {
					It("🧪 should: return hash result of src/assistant/*auto*.go", func() {
						fS := nef.NewUniversalABS()
						templatesSubPath := ""
						outputPath = filepath.Join(repo, sourcePath)

						sourceCode := gola.NewSourceCodeContainer(fS, outputPath, templatesSubPath)
						result, err := sourceCode.Signature()

						fmt.Printf("===> [👾] REGISTERED-HASH: '%v'\n", gola.RegisteredHash)
						fmt.Printf("===> Output:\n%v\n", result.Output)

						Expect(err).Error().To(BeNil())
						Expect(result.Output).NotTo(BeEmpty())
						Expect(result.Hash).To(MatchRegisteredHash(gola.RegisteredHash))
					})
				})
			})
		})

		When("with generator", func() { // 📍 must use Generator only
			Context("and: Test mode", func() {
				Context("and: without write", func() {
					It("🧪 should: return hash result of parsed contents sources", func() {
						const (
							doWrite = false
						)
						fS := nef.NewUniversalABS()
						templatesSubPath := ""
						outputPath = filepath.Join(repo, testPath)
						sourceCode := gola.NewSourceCodeContainer(fS, outputPath, templatesSubPath)
						result, err := sourceCode.Generator(doWrite).Run()

						Expect(err).Error().To(BeNil())
						Expect(result.Output).NotTo(BeEmpty())
						Expect(result.Hash).NotTo(MatchRegisteredHash(gola.RegisteredHash)) // 👻
					})
				})

				Context("and: with write", func() {
					It("🧪 should: return hash result of generators/gola/out/assistant/*auto*.go", func() {
						const (
							doWrite = true
						)

						fS := NewTestMemFS()
						templatesSubPath := ""
						outputPath = filepath.Join(repo, testPath)

						// only use testPath as this is relative and required when using mapFS
						sourceCode := gola.NewSourceCodeContainer(fS, testPath, templatesSubPath)
						result, err := sourceCode.Generator(doWrite).Run()
						path := filepath.Join(testPath, "option-validator-auto.go") // ☢️

						Expect(fS.FileExists(path)).To(BeTrue(),
							fmt.Sprintf("⛔⛔⛔ generated file '%v' not found\n", path),
						)

						fmt.Printf("===> [👾] REGISTERED-HASH: '%v'\n", gola.RegisteredHash)
						fmt.Printf("===> Output:\n%v\n", result.Output)

						Expect(err).Error().To(BeNil())
						Expect(result.Output).NotTo(BeEmpty())
						Expect(result.Hash).NotTo(MatchRegisteredHash(gola.RegisteredHash)) // 👻
					})
				})
			})
		})
	})
})
