package gola

import (
	_ "embed"
	"fmt"
	"path"

	"github.com/snivilised/cobrass/src/internal/lab"
	"github.com/snivilised/cobrass/src/internal/third/lo"
	nef "github.com/snivilised/nefilim"
)

var (
	noData = struct{}{}
	//go:embed signature.GO-HASH.txt
	RegisteredHash string
)

type SourceCodeGenerator struct {
	fS                   nef.UniversalFS
	sourceCodeCollection sourceCodeDataCollection
	types                typeCollection
	operators            operatorCollection
	doWrite              bool
}

func (g *SourceCodeGenerator) init(sourceCode sourceCodeDataCollection) {
	g.sourceCodeCollection = sourceCode
	g.operators = buildOperators()
	g.types = buildTypes()
}

func (g *SourceCodeGenerator) preDefinedOrderedTypeKeys() []TypeNameID {
	// TODO: We might want to re-define this order, for now, leave it
	// to match the same order as the PowerShell generator, at least
	// until we have verified the signature hashing matches (a difference
	// in the order will probably impact the has value)
	//
	return []TypeNameID{
		"Duration", "Enum",
		"Float32", "Float64",
		"Int", "Int16", "Int32", "Int64", "Int8",
		"IPMask", "IPNet",
		"String",
		"Uint16", "Uint32", "Uint64", "Uint8", "Uint",
	}
}

func (g *SourceCodeGenerator) Run() (*SignatureResult, error) {
	// g.sourceCodeCollection needs to be populated off the container,
	// we can't just make an empty collection, that's just pointless
	// an does not work.
	//
	contents := make(CodeContent, len(g.sourceCodeCollection))

	for _, k := range g.sourceCodeCollection.Keys() {
		page := g.sourceCodeCollection[k]
		yield := &generatedYield{}
		overwrite := lo.Ternary(g.fS.FileExists(page.FullPath()), "♻️ overwrite", "✨ new")

		if !page.active {
			fmt.Printf("===> 📛 (%v) SKIPPING generation of code to '%v' (%v)\n",
				page.name,
				page.OutputFileName(),
				overwrite,
			)

			continue
		}

		fmt.Printf("===> 🚀 (%v) generating code to '%v' (%v)\n",
			page.name,
			page.OutputFileName(),
			overwrite,
		)

		if err := g.static("header", page, yield); err != nil {
			return nil, err
		}

		for _, specTypeKey := range g.preDefinedOrderedTypeKeys() {
			// TODO: the executionInfo needs to match the structure of the dot
			// and child templates. executionInfo is the dot object and everything else
			// needs to be relative to this dot.
			//
			dot := executionInfo{
				Spec:      g.types[specTypeKey],
				Operators: g.operators,
			}

			if err := g.render("body", page, yield, &dot); err != nil {
				return nil, err
			}
		}

		if err := g.static("footer", page, yield); err != nil {
			return nil, err
		}

		contents[page.name] = yield.Content()

		if g.doWrite {
			if err := g.flush(page.FullPath(), yield); err != nil {
				fmt.Printf("---> 🔥 Write Error occurred for: '%v' (%v), aborting\n",
					page.name, err,
				)

				return nil, err
			}
		}
	}

	return g.signature(contents)
}

func (g *SourceCodeGenerator) static(
	section string, page *SourceCodeData, yield *generatedYield,
) error {
	if err := page.templ.ExecuteTemplate(
		&yield.buffer,
		page.section(section),
		noData,
	); err != nil {
		fmt.Printf("---> 🔥 Error executing static section template for: '%v' (%v), aborting\n",
			page.name, err,
		)

		return err
	}

	return nil
}

func (g *SourceCodeGenerator) render(
	section string, page *SourceCodeData, yield *generatedYield, dot *executionInfo,
) error {
	if err := page.templ.ExecuteTemplate(
		&yield.buffer,
		page.section(section),
		dot,
	); err != nil {
		fmt.Printf("---> 🔥 Error executing section template for: '%v' (%v), aborting\n",
			page.name, err,
		)

		return err
	}

	return nil
}

func (g *SourceCodeGenerator) signature(content CodeContent) (*SignatureResult, error) {
	return parseInline(content)
}

func (g *SourceCodeGenerator) flush(outputPath string, yield *generatedYield) error {
	directory := path.Dir(outputPath)

	if err := g.fS.MakeDirAll(
		directory,
		lab.Perms.Dir,
	); err != nil {
		return fmt.Errorf("failed to ensure parent directory '%v' exists (%v)", directory, err)
	}

	if err := g.fS.WriteFile(
		outputPath,
		yield.buffer.Bytes(),
		lab.Perms.File,
	); err != nil {
		return fmt.Errorf("failed to write generated code to '%v' (%v)", outputPath, err)
	}

	return nil
}
