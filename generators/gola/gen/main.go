package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/generators/gola"
	"github.com/snivilised/cobrass/generators/gola/internal/utils"
)

const (
	appName                    = "cobrass-gen"
	outputPathNotFoundExitCode = 2
)

var (
	testFlag         = flag.Bool("test", false, "generate code in test location?")
	cwdFlag          = flag.String("cwd", "", "current working directory")
	templatesSubPath = flag.String("templates", "", "templates sub path")
	write            = flag.Bool("write", false, "write generated code?")
	testPath         = filepath.Join("generators", "gola", "out", "assistant")
	sourcePath       = filepath.Join("src", "assistant")
)

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

func main() {
	flag.Usage = Usage
	flag.Parse()

	outputPath := lo.Ternary(*testFlag, testPath, sourcePath)

	if *cwdFlag == "" {
		fail("🔥 current working directory not specified")
	}

	absolutePath, _ := filepath.Abs(*cwdFlag)
	absolutePath = filepath.Join(absolutePath, outputPath)

	if !utils.FolderExists(absolutePath) {
		callback := func() {
			fmt.Printf("💥 --->      CWD: '%v' \n", *cwdFlag)
			fmt.Printf("💥 --->   OUTPUT: '%v' \n", outputPath)
			fmt.Printf("💥 ---> RESOLVED: '%v' \n", absolutePath)
		}
		fail(fmt.Sprintf("Output path '%v', not found", absolutePath), callback)

		return
	}

	sourceCode := gola.NewSourceCodeContainer(absolutePath, *templatesSubPath)
	mode := lo.Ternary(*testFlag, "🧪 Test", "🎁 Source")

	fmt.Printf("☑️ --->      CWD: '%v' \n", *cwdFlag)
	fmt.Printf("☑️ --->   OUTPUT: '%v' \n", outputPath)
	fmt.Printf("☑️ ---> RESOLVED: '%v' \n", absolutePath)
	fmt.Printf("---> 🐲 cobrass generator (%v, to: %v)\n", mode, absolutePath)

	if !*testFlag {
		if sourceCode.AnyMissing() {
			sourceCode.Verify(func(entry *gola.SourceCodeData) {
				exists := entry.Exists()
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

	if err := sourceCode.Generator(*write).Run(); err != nil {
		fail(err.Error())
	}
}
