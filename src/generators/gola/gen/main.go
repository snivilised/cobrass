package main

import (
	"flag"
	"fmt"

	"os"
	"path/filepath"

	"github.com/snivilised/cobrass/src/generators/gola"
	"github.com/snivilised/cobrass/src/internal/third/lo"
	nef "github.com/snivilised/nefilim"
)

const (
	appName                    = "cobrass-gen"
	outputPathNotFoundExitCode = 2
)

var (
	testFlag             = flag.Bool("test", false, "generate code in test location?")
	cwdFlag              = flag.String("cwd", "", "current working directory")
	templatesSubPathFlag = flag.String("templates", "", "templates sub path")
	writeFlag            = flag.Bool("write", false, "write generated code?")
	signFlag             = flag.Bool("sign", false, "show signature of existing code only")

	testPath   = filepath.Join("src", "generators", "gola", "out", "assistant")
	sourcePath = filepath.Join("src", "assistant")
)

func main() {
	flag.Usage = Usage
	flag.Parse()

	outputPath := lo.Ternary(*testFlag, testPath, sourcePath)
	nativeFS := nef.NewUniversalABS()

	if *signFlag {
		sign(nativeFS, outputPath)
	} else {
		gen(nativeFS, outputPath)
	}

	os.Exit(0)
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Use of %v:\n", appName)
	fmt.Fprintf(os.Stderr, "run the command from the root of the repo ...\n")
	fmt.Fprintf(os.Stderr, "\t%v [Flags]\n", appName)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func fail(reason string, callback ...func()) {
	if len(callback) > 0 {
		callback[0]()
	}

	fmt.Fprintf(os.Stderr, "🔥 Failed: '%v'\n", reason)
	flag.Usage()
	os.Exit(outputPathNotFoundExitCode)
}

func gen(fS nef.UniversalFS, outputPath string) {
	if *cwdFlag == "" {
		fail("🔥 current working directory not specified")
	}

	absolutePath, _ := filepath.Abs(*cwdFlag)
	absolutePath = filepath.Join(absolutePath, outputPath)

	if !fS.DirectoryExists(absolutePath) {
		callback := func() {
			fmt.Printf("💥 --->      CWD: '%v' \n", *cwdFlag)
			fmt.Printf("💥 --->   OUTPUT: '%v' \n", outputPath)
			fmt.Printf("💥 ---> RESOLVED: '%v' \n", absolutePath)
		}
		fail(fmt.Sprintf("Output path '%v', not found", absolutePath), callback)

		return
	}

	sourceCode := gola.NewSourceCodeContainer(fS, absolutePath, *templatesSubPathFlag)
	mode := lo.Ternary(*testFlag, "🧪 Test", "🎁 Source")

	fmt.Printf("☑️ --->      CWD: '%v' \n", *cwdFlag)
	fmt.Printf("☑️ --->   OUTPUT: '%v' \n", outputPath)
	fmt.Printf("☑️ ---> RESOLVED: '%v' \n", absolutePath)
	fmt.Printf("---> 🐲 cobrass generator (%v, to: %v)\n", mode, absolutePath)

	if !*testFlag {
		if sourceCode.AnyMissing() {
			sourceCode.Verify(func(entry *gola.SourceCodeData) {
				exists := fS.FileExists(entry.FullPath())
				indicator := lo.Ternary(exists, "✔️", "❌")
				status := lo.Ternary(exists, "exists", "missing")
				path := entry.FullPath()
				message := fmt.Sprintf("%v source file: '%v' %v", indicator, path, status)

				fmt.Printf("%v\n", message)
			})
		} else {
			fmt.Printf("✅ ---> ALL-PRESENT-AT: '%v' \n", absolutePath)
		}
	}

	result, err := sourceCode.Generator(*writeFlag).Run()

	if err != nil {
		fail(err.Error())
	}

	show(result)
}

func sign(fS nef.UniversalFS, sourcePath string) {
	templatesSubPath := ""

	sourceCode := gola.NewSourceCodeContainer(fS, sourcePath, templatesSubPath)
	result, err := sourceCode.Signature()

	if err != nil {
		fail(
			fmt.Sprintf(
				"🔥 failed to calculate hash for existing source at: '%v' (%v)",
				sourcePath, err,
			),
		)
	}

	show(result)
}

func show(result *gola.SignatureResult) {
	fmt.Println(result.Output)
}
