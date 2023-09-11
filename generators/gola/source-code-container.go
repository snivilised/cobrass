package gola

import (
	"fmt"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/samber/lo"
)

type sourceCodeDataCollection map[CodeFileName]*SourceCodeData

type SourceCodeContainer struct {
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
		path := page.templates()

		if d.templatesSubPath != "" {
			path = filepath.Join(d.templatesSubPath, path)
		}

		if templ, err := template.New(string(page.name)).Funcs(page.funcs).ParseGlob(path); err == nil {
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

func (d *SourceCodeContainer) sourceNames() []string {
	keys := lo.Keys(d.collection)
	sorted := lo.Map(keys, func(item CodeFileName, index int) string {
		return string(item)
	})
	sort.Strings(sorted)

	return sorted
}

func (d *SourceCodeContainer) AnyMissing() bool {
	// TODO: this needs to verified after all has been built
	//
	return d.ForEachUntil(func(data *SourceCodeData) bool {
		return !data.Exists()
	})
}

func (d *SourceCodeContainer) ForEach(fn func(entry *SourceCodeData)) {
	names := d.sourceNames()

	for _, name := range names {
		sourceCodeName := CodeFileName(name)
		data := (d.collection)[sourceCodeName]

		fn(data)
	}
}

// ForEachUntil returns true if exit's early, false otherwise
func (d *SourceCodeContainer) ForEachUntil(fn func(entry *SourceCodeData) bool) bool {
	names := d.sourceNames()

	for _, name := range names {
		sourceCodeName := CodeFileName(name)
		data := (d.collection)[sourceCodeName]

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
		doWrite:              doWrite,
		sourceCodeCollection: &d.collection,
	}
	generator.init()

	return generator
}

func NewSourceCodeContainer(
	absolutePath string,
	templatesSubPath string,
) *SourceCodeContainer {
	container := &SourceCodeContainer{
		absolutePath:     absolutePath,
		templatesSubPath: templatesSubPath,
	}
	container.init()

	return container
}
