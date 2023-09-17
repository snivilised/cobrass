package gola

import (
	"fmt"
	"path/filepath"
	"text/template"
)

// https://pkg.go.dev/text/template
// https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax

type CodeFileName string

type SourceCodeData struct {
	name        CodeFileName
	active      bool // THIS IS JUST TEMPORARY
	directory   string
	rootContent string
	templ       *template.Template
	funcs       map[string]any
}

func (d *SourceCodeData) OutputFileName() string {
	return string(d.name) + ".go"
}

func (d *SourceCodeData) FullPath() string {
	return filepath.Join(d.directory, d.OutputFileName())
}

func (d *SourceCodeData) templates() string {
	return fmt.Sprintf("templates/%v/*.go.tmpl", d.name)
}

func (d *SourceCodeData) section(s string) string {
	return fmt.Sprintf("%v-%v.go.tmpl", d.name, s)
}
