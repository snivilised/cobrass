package storage

import (
	"os"
)

type filepathAPI interface {
	// Intended only for those filepath methods that actually affect the
	// filesystem. Eg: there is no point in replicating methods like
	// filepath.Join here they are just path helpers that do not read/write
	// to the filesystem.
	// Currently, there is no requirement for using any filepath methods
	// with the golang generator, hence nothing is defined here. We may
	// want to replicate this filesystem model in other contexts, so this
	// will serve as a reminder in the intended use of this interface.
}

type ExistsInFS interface {
	FileExists(path string) bool
	DirectoryExists(path string) bool
}

type ReadFromFS interface {
	ReadFile(name string) ([]byte, error)
	ReadDir(name string) ([]os.DirEntry, error)
}

type WriteToFS interface {
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type ReadOnlyVirtualFS interface {
	filepathAPI
	ExistsInFS
	ReadFromFS
}

type VirtualFS interface {
	filepathAPI
	ExistsInFS
	ReadFromFS
	WriteToFS
}

type VirtualBackend string
