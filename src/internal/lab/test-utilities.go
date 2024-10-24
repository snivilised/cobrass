package lab

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
	_, filename, _, _ := runtime.Caller(0) //nolint:dogsled // use of 3 _ is out of our control
	return Path(filepath.Dir(filename), relative)
}
