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
	testFlag           = flag.Bool("test", false, "generate code in test location?")
	cwdFlag            = flag.String("cwd", "", "current working directory")
	testPath           = filepath.Join("generators", "gola", "out", "assistant")
	sourcePath         = filepath.Join("src", "assistant")
	outputPathNotFound = "Output path '%v', not found"
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

// ???
// https://askgolang.com/how-to-get-current-directory-in-golang/

func main() {
	flag.Usage = Usage
	flag.Parse()

	outputPath := lo.Ternary(*testFlag, testPath, sourcePath)

	if *cwdFlag == "" {
		fail("🔥 current working directory not specified")
	}

	absolutePath, _ := filepath.Abs(*cwdFlag)
	absolutePath = filepath.Join(absolutePath, outputPath)

	if !utils.FileExists(absolutePath) {
		callback := func() {
			fmt.Printf("💥 --->      CWD: '%v' \n", *cwdFlag)
			fmt.Printf("💥 --->   OUTPUT: '%v' \n", outputPath)
			fmt.Printf("💥 ---> RESOLVED: '%v' \n", absolutePath)
		}
		fail(fmt.Sprintf(outputPathNotFound, absolutePath), callback)

		return
	}

	sourceCode := gola.NewSourceCodeContainer()
	mode := lo.Ternary(*testFlag, "🧪 Test", "🎁 Source")

	fmt.Printf("☑️ --->      CWD: '%v' \n", *cwdFlag)
	fmt.Printf("☑️ --->   OUTPUT: '%v' \n", outputPath)
	fmt.Printf("☑️ ---> RESOLVED: '%v' \n", absolutePath)
	fmt.Printf("---> 🐲 cobrass generator (%v, to: %v)\n", mode, absolutePath)

	if !*testFlag {
		if sourceCode.AnyMissing(absolutePath) {
			sourceCode.Verify(absolutePath, func(entry *gola.SourceCodeData) {
				exists := entry.Exists(absolutePath)
				indicator := lo.Ternary(exists, "✔️", "❌")
				status := lo.Ternary(exists, "exists", "missing")
				path := entry.FullPath(absolutePath)
				message := fmt.Sprintf("%v source file: '%v' %v", indicator, path, status)

				fmt.Printf("%v\n", message)
			})
		} else {
			fmt.Printf("✅ ---> ALL-PRESENT-AT: '%v' \n", absolutePath)
		}
	}

	sourceCode.Generator().Run()

	logicalEnum := gola.LogicalType{
		TypeName:           "Enum",
		GoType:             "string",
		DisplayType:        "enum",
		UnderlyingTypeName: "String",
		FlagName:           "Format",
		Short:              "f",
		Def:                "xml",
	}

	fmt.Printf("---> 🐲 cobrass generator (enum: %+v)\n", logicalEnum)
}
