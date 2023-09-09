package gola

import (
	_ "embed"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/generators/gola/internal/utils"
)

// https://pkg.go.dev/text/template
// https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax

type CodeFileName string

var (
	//go:embed templates/top/option-validator-auto.templ.top.txt
	optionValidatorAutoTop string

	//go:embed templates/top/option-validator-auto_test.templ.top.txt
	optionValidatorAutoTestTop string

	//go:embed templates/top/param-set-auto.templ.top.txt
	paramSetAutoTop string

	//go:embed templates/top/param-set-auto_test.templ.top.txt
	paramSetAutoTestTop string

	//go:embed templates/top/param-set-binder-helpers-auto.templ.top.txt
	paramSetBinderHelpersAutoTop string

	//go:embed templates/top/param-set-binder-helpers-auto_test.templ.top.txt
	paramSetBinderHelpersAutoTestTop string
)

type SourceCodeData struct {
	name  CodeFileName
	top   string
	templ *template.Template
}

func (d *SourceCodeData) FileName() string {
	return string(d.name) + ".go"
}

func (d *SourceCodeData) IsTest() bool {
	return strings.HasSuffix(string(d.name), "_test")
}

func (d *SourceCodeData) Exists(absolutePath string) bool {
	return utils.FileExists(d.FullPath(absolutePath))
}

func (d *SourceCodeData) FullPath(absolutePath string) string {
	filename := d.FileName()
	return filepath.Join(absolutePath, filename)
}

type sourceCodeDataCollection map[CodeFileName]*SourceCodeData

type SourceCodeContainer struct {
	collection sourceCodeDataCollection
}

func (d *SourceCodeContainer) init() {
	d.collection = sourceCodeDataCollection{
		"option-validator-auto": &SourceCodeData{
			name: "option-validator-auto",
			top:  optionValidatorAutoTop,
		},
		"option-validator-auto_test": &SourceCodeData{
			name: "option-validator-auto_test",
			top:  optionValidatorAutoTestTop,
		},
		"param-set-auto": &SourceCodeData{
			name: "param-set-auto",
			top:  paramSetAutoTop,
		},
		"param-set-auto_test": &SourceCodeData{
			name: "param-set-auto_test",
			top:  paramSetAutoTestTop,
		},
		"param-set-binder-helpers-auto": &SourceCodeData{
			name: "param-set-binder-helpers-auto",
			top:  paramSetBinderHelpersAutoTop,
		},
		"param-set-binder-helpers-auto_test": &SourceCodeData{
			name: "param-set-binder-helpers-auto_test",
			top:  paramSetBinderHelpersAutoTestTop,
		},
	}

	for _, data := range d.collection {
		if templ, err := template.New(string(data.name)).Parse(data.top); err == nil {
			data.templ = templ
		}
	}
}

func (d *SourceCodeContainer) sourceNames() []string {
	keys := lo.Keys(d.collection)
	sorted := lo.Map(keys, func(item CodeFileName, index int) string {
		return string(item)
	})
	sort.Strings(sorted)

	return sorted
}

func (d *SourceCodeContainer) AnyMissing(absolutePath string) bool {
	names := d.sourceNames()

	for _, name := range names {
		sourceCodeName := CodeFileName(name)
		data := (d.collection)[sourceCodeName]
		exists := data.Exists(absolutePath)

		if !exists {
			return true
		}
	}

	return false
}

func (d *SourceCodeContainer) ReportAll(fn ...func(entry *SourceCodeData)) {
	names := d.sourceNames()

	for _, name := range names {
		sourceCodeName := CodeFileName(name)
		data := (d.collection)[sourceCodeName]

		if len(fn) > 0 {
			fn[0](data)
		}
	}
}

func (d *SourceCodeContainer) Verify(absolutePath string, fn ...func(entry *SourceCodeData)) bool {
	result := d.AnyMissing(absolutePath)

	if !result {
		return result
	}

	names := d.sourceNames()

	for _, name := range names {
		sourceCodeName := CodeFileName(name)
		data := (d.collection)[sourceCodeName]

		if len(fn) > 0 {
			fn[0](data)
		}
	}

	return result
}

func (d *SourceCodeContainer) Generator() *SourceCodeGenerator {
	return &SourceCodeGenerator{}
}

func NewSourceCodeContainer() *SourceCodeContainer {
	data := &SourceCodeContainer{}
	data.init()

	return data
}
