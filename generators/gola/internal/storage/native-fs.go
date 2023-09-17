package storage

import (
	"os"
)

type nativeFS struct {
	backend VirtualBackend
}

func UseNativeFS() VirtualFS {
	return &nativeFS{
		backend: "native",
	}
}

func (fs *nativeFS) FileExists(path string) bool {
	result := false
	if info, err := os.Lstat(path); err == nil {
		result = !info.IsDir()
	}

	return result
}

func (fs *nativeFS) DirectoryExists(path string) bool {
	result := false
	if info, err := os.Lstat(path); err == nil {
		result = info.IsDir()
	}

	return result
}

func (fs *nativeFS) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs *nativeFS) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (fs *nativeFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (fs *nativeFS) ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}
