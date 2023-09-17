package gola

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/generators/gola/internal/storage"
)

// parseInline does not need to use the filesystem to acquire
// its content. It is provided as a result of the code
// generation process instead.
func parseInline(contents CodeContent) (*SignatureResult, error) {
	return parseContents(contents)
}

// parseFromFS parses content acquired from the filesystem.
func parseFromFS(vfs storage.VirtualFS, directoryPath string) (*SignatureResult, error) {
	var (
		entries         []fs.DirEntry
		contents        CodeContent
		readErr, acqErr error
	)

	if entries, readErr = readEntries(vfs, directoryPath); readErr != nil {
		return nil, readErr
	}

	if contents, acqErr = acquire(vfs, directoryPath, entries); acqErr != nil {
		return nil, acqErr
	}

	return parseContents(contents)
}

func readEntries(vfs storage.VirtualFS, directoryPath string) ([]fs.DirEntry, error) {
	entries, err := vfs.ReadDir(directoryPath)

	if err != nil {
		return nil, err
	}

	autoPattern := "*auto*.go"
	testPattern := "*auto*_test.go"
	entries = lo.Filter(entries, func(item fs.DirEntry, index int) bool {
		auto, _ := filepath.Match(autoPattern, item.Name())
		test, _ := filepath.Match(testPattern, item.Name())
		return !item.IsDir() && auto && !test
	})

	if len(entries) == 0 {
		return nil, fmt.Errorf(
			"found no applicable source files at: '%v'",
			directoryPath,
		)
	}

	return entries, nil
}

func acquire(vfs storage.VirtualFS, directoryPath string, entries []fs.DirEntry) (CodeContent, error) {
	contents := make(CodeContent, len(entries))

	for _, entry := range entries {
		name := entry.Name()
		sourcePath := filepath.Join(directoryPath, name)
		c, err := vfs.ReadFile(sourcePath)

		if err != nil {
			return nil, err
		}

		contents[CodeFileName(name)] = string(c)
	}

	return contents, nil
}

func parseContents(contents CodeContent) (*SignatureResult, error) {
	ending := "\n"
	hashBuilder := strings.Builder{}
	metrics := make(SignatureCountsBySource)
	totals := &SignatureCounts{}

	for _, name := range contents.Keys() {
		c := contents[name]

		metrics[name] = &SignatureCounts{}
		lines := strings.Split(c, ending)

		for _, line := range lines {
			if strings.HasPrefix(line, "func") {
				index := strings.LastIndex(line, " {")
				if index >= 0 {
					signature := line[0 : index+1]
					hashBuilder.WriteString(fmt.Sprintf("%v\n", strings.TrimSpace(signature)))
					metrics[name].Func++
				}
			} else if strings.HasPrefix(line, "type") {
				hashBuilder.WriteString(fmt.Sprintf("%v\n", strings.TrimSpace(line)))
				metrics[name].Type++
			}
		}

		totals.Func += metrics[name].Func
		totals.Type += metrics[name].Type
	}

	sha256hash := hash(hashBuilder.String())
	outputBuilder := strings.Builder{}
	outputBuilder.WriteString(fmt.Sprintf("===> [🤖]        THIS-HASH: '%v'\n", sha256hash))
	outputBuilder.WriteString(fmt.Sprintf("===> [👾]  REGISTERED-HASH: '%v'\n", RegisteredHash))

	for _, n := range metrics.Keys() {
		metric := metrics[n]
		funcs := fmt.Sprintf("'%v'", metric.Func)
		types := fmt.Sprintf("'%v'", metric.Type)
		line := fmt.Sprintf(
			"---> 🍄 [%36v] Signature Counts - 🍅functions: %6v, 🥦types: %6v",
			n, funcs, types,
		)

		outputBuilder.WriteString(fmt.Sprintf("%v\n", line))
	}

	funcs := fmt.Sprintf("'%v'", totals.Func)
	types := fmt.Sprintf("'%v'", totals.Type)
	totalLine := fmt.Sprintf(
		"---> 🍄 Total Counts - 🍅functions: %6v, 🥦types: %6v",
		funcs, types,
	)

	outputBuilder.WriteString(fmt.Sprintf("%v\n", totalLine))

	status := lo.Ternary(sha256hash == RegisteredHash,
		"✔️ Hashes are equal",
		"💥 Api changes detected",
	)
	outputBuilder.WriteString(fmt.Sprintf(">>>> Status: %v\n", status))

	return &SignatureResult{
		Totals:   totals,
		Counters: metrics,
		Hash:     sha256hash,
		Status:   status,
		Output:   outputBuilder.String(),
	}, nil
}

func hash(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))

	return hex.EncodeToString(hasher.Sum(nil))
}
