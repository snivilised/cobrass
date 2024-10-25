package lab

import (
	"io/fs"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	Perms = struct {
		File fs.FileMode
		Dir  fs.FileMode
	}{
		File: 0o666, //nolint:mnd // ok (pedantic)
		Dir:  0o777, //nolint:mnd // ok (pedantic)
	}
)

func Path(parent, relative string) string {
	segments := strings.Split(relative, "/")
	return filepath.Join(append([]string{parent}, segments...)...)
}

func Repo(relative string) string {
	_, filename, _, _ := runtime.Caller(0) //nolint:dogsled // use of 3 _ is out of our control
	return Path(filepath.Dir(filename), relative)
}
