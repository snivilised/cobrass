package gola_test

import (
	"io/fs"
	"os"
	"strings"
	"testing/fstest"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/src/internal/lab"
	nef "github.com/snivilised/nefilim"
)

// nefilim should export this struct for us, so we dont have to re-implement
type TestMemFS struct {
	fstest.MapFS
}

func NewTestMemFS() *TestMemFS {
	return &TestMemFS{
		MapFS: fstest.MapFS{},
	}
}

func (f *TestMemFS) FileExists(name string) bool {
	if mapFile, found := f.MapFS[name]; found && !mapFile.Mode.IsDir() {
		return true
	}

	return false
}

func (f *TestMemFS) DirectoryExists(name string) bool {
	if mapFile, found := f.MapFS[name]; found && mapFile.Mode.IsDir() {
		return true
	}

	return false
}

func (f *TestMemFS) Create(name string) (*os.File, error) {
	if _, err := f.Stat(name); err == nil {
		return nil, fs.ErrExist
	}

	file := &fstest.MapFile{
		Mode: lab.Perms.File,
	}

	f.MapFS[name] = file
	// TODO: this needs a resolution using a file interface
	// rather than using os.File which is a struct not an
	// interface
	dummy := &os.File{}

	return dummy, nil
}

func (f *TestMemFS) MakeDir(name string, perm os.FileMode) error {
	if !fs.ValidPath(name) {
		return nef.NewInvalidPathError("MakeDir", name)
	}

	if _, found := f.MapFS[name]; !found {
		f.MapFS[name] = &fstest.MapFile{
			Mode: perm | os.ModeDir,
		}
	}

	return nil
}

func (f *TestMemFS) MakeDirAll(name string, perm os.FileMode) error {
	if !fs.ValidPath(name) {
		return nef.NewInvalidPathError("MakeDirAll", name)
	}

	segments := strings.Split(name, "/")

	_ = lo.Reduce(segments,
		func(acc []string, s string, _ int) []string {
			acc = append(acc, s)
			path := strings.Join(acc, "/")

			if _, found := f.MapFS[path]; !found {
				f.MapFS[path] = &fstest.MapFile{
					Mode: perm | os.ModeDir,
				}
			}

			return acc
		}, []string{},
	)

	return nil
}

func (f *TestMemFS) WriteFile(name string, data []byte, perm os.FileMode) error {
	if _, err := f.Stat(name); err == nil {
		return fs.ErrExist
	}

	f.MapFS[name] = &fstest.MapFile{
		Data: data,
		Mode: perm,
	}

	return nil
}

func (f *TestMemFS) Change(_, _ string) error {
	panic("NOT-IMPL: TestMemFS.Change")
}

func (f *TestMemFS) Copy(_, _ string) error {
	panic("NOT-IMPL: TestMemFS.Copy")
}

func (f *TestMemFS) CopyFS(_ string, _ fs.FS) error {
	panic("NOT-IMPL: TestMemFS.CopyFS")
}

func (f *TestMemFS) Ensure(_ nef.PathAs) (string, error) {
	panic("NOT-IMPL: TestMemFS.Ensure")
}

func (f *TestMemFS) Move(_, _ string) error {
	panic("NOT-IMPL: TestMemFS.Move")
}

func (f *TestMemFS) Remove(_ string) error {
	panic("NOT-IMPL: TestMemFS.Remove")
}

func (f *TestMemFS) RemoveAll(_ string) error {
	panic("NOT-IMPL: TestMemFS.RemoveAll")
}

func (f *TestMemFS) Rename(_, _ string) error {
	panic("NOT-IMPL: TestMemFS.Rename")
}
