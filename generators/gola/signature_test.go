package gola_test

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok
	. "github.com/onsi/gomega"    //nolint:revive // ok
	"github.com/snivilised/cobrass/generators/gola"
	nef "github.com/snivilised/nefilim"
)

var (
	// to go into a lab
	Perms = struct {
		File fs.FileMode
		Dir  fs.FileMode
	}{
		File: 0o666, //nolint:mnd // ok (pedantic)
		Dir:  0o777, //nolint:mnd // ok (pedantic)
	}
)

type setupFile struct {
	path string
	data []byte
}

func setup(fs nef.UniversalFS, directoryPath string, files ...setupFile) {
	if e := fs.MakeDirAll(directoryPath, Perms.Dir); e != nil {
		Fail(e.Error())
	}

	for _, f := range files {
		if e := fs.WriteFile(f.path, f.data, Perms.File); e != nil {
			Fail(e.Error())
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

var _ = XDescribe("Signature", Ordered, func() {
	var (
		repo,
		testPath,
		sourcePath,
		outputPath string
	)

	BeforeAll(func() {
		repo = Repo("../..")
		testPath = filepath.Join("generators", "gola", "out", "assistant")
		sourcePath = filepath.Join("src", "assistant")
	})

	Context("sign", func() {
		When("standalone", func() { // 📍 can't use Generate, only Signature()
			Context("and: Test mode", func() {
				Context("and: without write", func() {
					It("🧪 should: return hash result of newly generated content", func() {
						// use mapFile
						//
						mfs := nef.NewUniversalABS()
						templatesSubPath := ""
						outputPath = filepath.Join(repo, testPath)

						// ☢️ TODO: !! what about the other files?; this is short term and will
						// need to accommodate the other source files.
						//
						path := filepath.Join(outputPath, "option-validator-auto.go")

						if data, err := os.ReadFile(path); err != nil {
							setup(mfs, outputPath, setupFile{
								path: path,
								data: data,
							})
						}
						Expect(mfs.FileExists(path)).To(BeTrue())

						sourceCode := gola.NewSourceCodeContainer(mfs, outputPath, templatesSubPath)
						result, err := sourceCode.Signature()

						Expect(err).Error().To(BeNil())
						Expect(result.Hash).NotTo(BeEmpty())

						// ⚠️ Can't expect this to match yet, since the registered hash is
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
						nfs := nef.NewUniversalABS() // TODO: check
						templatesSubPath := ""
						outputPath = filepath.Join(repo, sourcePath)

						sourceCode := gola.NewSourceCodeContainer(nfs, outputPath, templatesSubPath)
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
						nfs := nef.NewUniversalABS()
						templatesSubPath := ""
						outputPath = filepath.Join(repo, testPath)
						sourceCode := gola.NewSourceCodeContainer(nfs, outputPath, templatesSubPath)
						doWrite := false
						result, err := sourceCode.Generator(doWrite).Run()

						Expect(err).Error().To(BeNil())
						Expect(result.Output).NotTo(BeEmpty())
						Expect(result.Hash).NotTo(MatchRegisteredHash(gola.RegisteredHash)) // ⚠️
					})
				})

				Context("and: with write", func() {
					It("🧪 should: return hash result of generators/gola/out/assistant/*auto*.go", func() {
						mfs := nef.NewUniversalABS() // use mapFS
						templatesSubPath := ""
						outputPath = filepath.Join(repo, testPath)

						sourceCode := gola.NewSourceCodeContainer(mfs, outputPath, templatesSubPath)
						doWrite := true
						result, err := sourceCode.Generator(doWrite).Run()
						path := filepath.Join(outputPath, "option-validator-auto.go") // ☢️

						Expect(mfs.FileExists(path)).To(BeTrue(),
							fmt.Sprintf("⛔⛔⛔ generated file '%v' not found\n", path),
						)

						fmt.Printf("===> [👾] REGISTERED-HASH: '%v'\n", gola.RegisteredHash)
						fmt.Printf("===> Output:\n%v\n", result.Output)

						Expect(err).Error().To(BeNil())
						Expect(result.Output).NotTo(BeEmpty())
						Expect(result.Hash).NotTo(MatchRegisteredHash(gola.RegisteredHash)) // ⚠️
					})
				})
			})
		})
	})
})
