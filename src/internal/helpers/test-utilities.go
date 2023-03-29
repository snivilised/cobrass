package helpers

import (
	"path/filepath"
	"runtime"
	"strings"
)

func Path(parent, relative string) string {
	segments := strings.Split(relative, "/")
	return filepath.Join(append([]string{parent}, segments...)...)
}

func Repo(relative string) string {
	_, filename, _, _ := runtime.Caller(0)
	return Path(filepath.Dir(filename), relative)
}
