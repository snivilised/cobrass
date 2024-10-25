package gola

import (
	"fmt"
	"path/filepath"

	"text/template"

	"github.com/snivilised/cobrass/src/generators/gola/internal/collections"
	"github.com/snivilised/cobrass/src/internal/third/lo"
	nef "github.com/snivilised/nefilim"
)

type sourceCodeDataCollection = collections.OrderedKeysMap[CodeFileName, *SourceCodeData]

type SourceCodeContainer struct {
	fS               nef.UniversalFS
	absolutePath     string
	templatesSubPath string
	collection       sourceCodeDataCollection
}

func (d *SourceCodeContainer) init() {
	d.collection = sourceCodeDataCollection{
		"option-validator-auto": &SourceCodeData{
			name:      "option-validator-auto",
			active:    true,
			directory: d.absolutePath,
			// We only define simple functions as template functions
			// Anything more should be defined on the typeSpec
			//
			funcs: map[string]any{
				"getValidatorFn": func(typeName string) string {
					return typeName + "ValidatorFn"
				},
				"getValidatorStruct": func(typeName string) string {
					return typeName + "OptionValidator"
				},
				"getDisplayType": func(spec TypeSpec) string {
					return lo.Ternary(spec.DisplayType != "",
						spec.DisplayType,
						spec.GoType,
					)
				},
				"getSliceTypeName": func(spec TypeSpec) string {
					return spec.TypeName + "Slice"
				},
				"getSliceType": func(spec TypeSpec) string {
					return "[]" + spec.GoType
				},
				"getSliceValidationFn": func(spec TypeSpec) string {
					return spec.TypeName + "SliceValidatorFn"
				},
			},
		},
		"option-validator-auto_test": &SourceCodeData{
			name:      "option-validator-auto_test",
			directory: d.absolutePath,
			funcs:     map[string]any{},
		},
		"param-set-auto": &SourceCodeData{
			name:      "param-set-auto",
			directory: d.absolutePath,
			funcs:     map[string]any{},
		},
		"param-set-auto_test": &SourceCodeData{
			name:      "param-set-auto_test",
			directory: d.absolutePath,
			funcs:     map[string]any{},
		},
		"param-set-binder-helpers-auto": &SourceCodeData{
			name:      "param-set-binder-helpers-auto",
			directory: d.absolutePath,
			funcs:     map[string]any{},
		},
		"param-set-binder-helpers-auto_test": &SourceCodeData{
			name:      "param-set-binder-helpers-auto_test",
			directory: d.absolutePath,
			funcs:     map[string]any{},
		},
	}

	for _, page := range d.collection {
		relativeTemplPath := page.templates()

		if d.templatesSubPath != "" {
			relativeTemplPath = filepath.Join(d.templatesSubPath, relativeTemplPath)
		}

		if templ, err := template.New(string(page.name)).Funcs(page.funcs).ParseGlob(
			relativeTemplPath,
		); err == nil {
			page.templ = templ
		} else {
			panic(
				fmt.Errorf("💥 error creating templates for '%v' (%v)",
					page.name, err,
				),
			)
		}
	}
}

func (d *SourceCodeContainer) contentPath() string {
	return filepath.Join(d.absolutePath, d.templatesSubPath)
}

func (d *SourceCodeContainer) AnyMissing() bool {
	return d.ForEachUntil(func(data *SourceCodeData) bool {
		return !d.fS.FileExists(data.FullPath())
	})
}

func (d *SourceCodeContainer) ForEach(fn func(entry *SourceCodeData)) {
	names := d.collection.Keys()

	for _, name := range names {
		data := (d.collection)[name]

		fn(data)
	}
}

// ForEachUntil returns true if exit's early, false otherwise
func (d *SourceCodeContainer) ForEachUntil(fn func(entry *SourceCodeData) bool) bool {
	names := d.collection.Keys()

	for _, name := range names {
		data := (d.collection)[name]

		if fn(data) {
			return true
		}
	}

	return false
}

func (d *SourceCodeContainer) Verify(fn func(entry *SourceCodeData)) {
	if !d.AnyMissing() {
		return
	}

	d.ForEach(fn)
}

func (d *SourceCodeContainer) Generator(
	doWrite bool,
) *SourceCodeGenerator {
	generator := &SourceCodeGenerator{
		fS:                   d.fS,
		doWrite:              doWrite,
		sourceCodeCollection: d.collection,
	}

	d.init()
	generator.init(d.collection)

	return generator
}

// Signature used to compose the SHA256 hash of
// pre-generated source code.
func (d *SourceCodeContainer) Signature() (*SignatureResult, error) {
	return parseFromFS(d.fS, d.contentPath())
}

func NewSourceCodeContainer(
	fS nef.UniversalFS,
	absolutePath string,
	templatesSubPath string,
) *SourceCodeContainer {
	container := &SourceCodeContainer{
		fS:               fS,
		absolutePath:     absolutePath,
		templatesSubPath: templatesSubPath,
	}

	return container
}
