package gola

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/snivilised/cobrass/generators/gola/internal/utils"
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

func (d *SourceCodeData) GeneratedFileName() string {
	return string(d.name) + ".go"
}

func (d *SourceCodeData) IsTest() bool {
	return strings.HasSuffix(string(d.name), "_test")
}

func (d *SourceCodeData) Exists() bool {
	return utils.FileExists(d.FullPath())
}

func (d *SourceCodeData) FullPath() string {
	return filepath.Join(d.directory, d.GeneratedFileName())
}

func (d *SourceCodeData) templates() string {
	return fmt.Sprintf("templates/%v/*.go.tmpl", d.name)
}

func (d *SourceCodeData) child(_ string) string {
	panic("child template name not yet defined (<name>-XXX.go.tmpl?)")
}

func (d *SourceCodeData) section(s string) string {
	return fmt.Sprintf("%v-%v.go.tmpl", d.name, s)
}
