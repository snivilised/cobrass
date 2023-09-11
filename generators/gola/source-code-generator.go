package gola

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/samber/lo"
)

var (
	noData = struct{}{}
)

type typeCollection map[TypeNameID]*TypeSpec
type operatorCollection []*Operator

type SourceCodeGenerator struct {
	sourceCodeCollection *sourceCodeDataCollection
	types                *typeCollection
	operators            operatorCollection
	doWrite              bool
}

func (g *SourceCodeGenerator) init() {
	g.buildOperators()
	g.types = buildTypes()
}

func (g *SourceCodeGenerator) sortedTypeKeys() []TypeNameID {
	keys := lo.Keys(*g.types)
	sort.Strings(keys)

	return keys
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

func (g *SourceCodeGenerator) Run() error {
	for _, page := range *g.sourceCodeCollection {
		var buffer bytes.Buffer

		overwrite := lo.Ternary(page.Exists(), "♻️ overwrite", "✨ new")
		if page.active {
			fmt.Printf("===> 🚀 (%v) generating code to '%v' (%v)\n",
				page.name,
				page.GeneratedFileName(),
				overwrite,
			)

			if err := g.renderStatic("header", page, &buffer); err != nil {
				return err
			}

			for _, specTypeKey := range g.preDefinedOrderedTypeKeys() {
				fmt.Printf("---> 🍬🍬🍬 generating code for type '%v'\n", specTypeKey)
				// TODO: the executionInfo needs to match the structure of the dot
				// and child templates. executionInfo is the dot object and everything else
				// needs to be relative to this dot.
				//
				dot := executionInfo{
					Spec:      (*g.types)[specTypeKey],
					Operators: g.operators,
				}

				if err := g.renderSection("body", page, &dot, &buffer); err != nil {
					return err
				}
			}

			if err := g.renderStatic("footer", page, &buffer); err != nil {
				return err
			}

			if g.doWrite {
				if err := g.flush(page.FullPath(), &buffer); err != nil {
					fmt.Printf("---> 🔥 Write Error occurred for: '%v' (%v), aborting\n",
						page.name, err,
					)

					return err
				}
			} else {
				fmt.Printf("===> 🧊🧊🧊 Generated %v content: ...\n",
					page.name,
				)
				fmt.Printf("%v\n", buffer.String())
			}
		} else {
			fmt.Printf("===> 📛 (%v) SKIPPING generation of code to '%v' (%v)\n",
				page.name,
				page.GeneratedFileName(),
				overwrite,
			)
		}
	}

	return nil
}

func (g *SourceCodeGenerator) renderStatic(
	section string, page *SourceCodeData, buffer *bytes.Buffer,
) error {
	static := page.section(section)
	if err := page.templ.ExecuteTemplate(
		buffer,
		static,
		noData,
	); err != nil {
		fmt.Printf("---> 🔥 Error executing static section template for: '%v' (%v), aborting\n",
			page.name, err,
		)

		return err
	}

	return nil
}

func (g *SourceCodeGenerator) renderSection(
	section string, page *SourceCodeData, dot *executionInfo, buffer *bytes.Buffer,
) error {
	body := page.section(section)
	if err := page.templ.ExecuteTemplate(
		buffer,
		body,
		dot,
	); err != nil {
		fmt.Printf("---> 🔥 Error executing section template for: '%v' (%v), aborting\n",
			page.name, err,
		)

		return err
	}

	return nil
}

func (g *SourceCodeGenerator) buildOperators() {
	g.operators = operatorCollection{
		&Operator{
			Name:          "Within",
			Documentation: "fails validation if the option value does not lie within 'low' and 'high' (inclusive)",
		},
	}
}

func (g *SourceCodeGenerator) flush(filepath string, content *bytes.Buffer) error {
	faydeaudeau := 0o777
	directory := path.Dir(filepath)

	if err := os.MkdirAll(directory, os.FileMode(faydeaudeau)); err != nil {
		return fmt.Errorf("failed to ensure parent directory '%v' exists (%v)", directory, err)
	}

	beezledub := 0o666

	if err := os.WriteFile(filepath, content.Bytes(), os.FileMode(beezledub)); err != nil {
		return fmt.Errorf("failed to write generated code to '%v' (%v)", filepath, err)
	}

	return nil
}
