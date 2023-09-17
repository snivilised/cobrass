package storage

import (
	"os"

	"github.com/avfs/avfs/vfs/memfs"
)

type memFS struct {
	backend VirtualBackend
	mfs     *memfs.MemFS
}

func UseMemFS() VirtualFS {
	return &memFS{
		backend: "mem",
		mfs:     memfs.New(),
	}
}

func (fs *memFS) FileExists(path string) bool {
	result := false
	if info, err := fs.mfs.Lstat(path); err == nil {
		result = !info.IsDir()
	}

	return result
}

func (fs *memFS) DirectoryExists(path string) bool {
	result := false
	if info, err := fs.mfs.Lstat(path); err == nil {
		result = info.IsDir()
	}

	return result
}

func (fs *memFS) MkdirAll(path string, perm os.FileMode) error {
	return fs.mfs.MkdirAll(path, perm)
}

func (fs *memFS) WriteFile(name string, data []byte, perm os.FileMode) error {
	return fs.mfs.WriteFile(name, data, perm)
}

func (fs *memFS) ReadFile(name string) ([]byte, error) {
	return fs.mfs.ReadFile(name)
}

func (fs *memFS) ReadDir(name string) ([]os.DirEntry, error) {
	return fs.mfs.ReadDir(name)
}
