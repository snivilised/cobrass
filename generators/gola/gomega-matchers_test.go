package gola_test

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/generators/gola"
	"github.com/snivilised/cobrass/generators/gola/internal/storage"

	. "github.com/onsi/gomega/types"
)

type (
	AreAllSourceCodeFilesPresentMatcher struct {
		fs        storage.VirtualFS
		directory string
	}
)

// ContainAllSourceCodeFilesAt
func ContainAllSourceCodeFilesAt(fs storage.VirtualFS, directory string) GomegaMatcher {
	return &AreAllSourceCodeFilesPresentMatcher{
		fs:        fs,
		directory: directory,
	}
}

func (m *AreAllSourceCodeFilesPresentMatcher) Match(actual interface{}) (bool, error) {
	sourceCode, ok := actual.(*gola.SourceCodeContainer)
	if !ok {
		return false, fmt.Errorf("matcher expected a SourceCodeContainer value (actual: '%v')", actual)
	}

	return !sourceCode.AnyMissing(), nil
}

func (m *AreAllSourceCodeFilesPresentMatcher) report(
	negated bool,
	sourceCode *gola.SourceCodeContainer,
) string {
	builder := strings.Builder{}
	not := lo.Ternary(negated, " NOT ", " ")
	builder.WriteString(
		fmt.Sprintf("🔥 Expected all source code files %vto be present\n", not),
	)

	sourceCode.ForEach(func(entry *gola.SourceCodeData) {
		exists := m.fs.FileExists(entry.FullPath())
		indicator := lo.Ternary(exists, "✔️", "❌")
		status := lo.Ternary(exists, "exists", "missing")
		path := entry.FullPath()
		message := fmt.Sprintf("%v source file: '%v' %v\n", indicator, path, status)
		builder.WriteString(message)
	})

	return builder.String()
}

func (m *AreAllSourceCodeFilesPresentMatcher) FailureMessage(actual interface{}) string {
	sourceCode, ok := actual.(*gola.SourceCodeContainer)
	if !ok {
		return fmt.Sprintf("matcher expected a SourceCodeContainer value (actual: '%v')", actual)
	}

	return m.report(false, sourceCode)
}

func (m *AreAllSourceCodeFilesPresentMatcher) NegatedFailureMessage(actual interface{}) string {
	sourceCode, ok := actual.(*gola.SourceCodeContainer)
	if !ok {
		return fmt.Sprintf("matcher expected a SourceCodeContainer value (actual: '%v')", actual)
	}

	return m.report(true, sourceCode)
}

type (
	HashMatcher struct {
		expectedHash string
	}
)

// MatchRegisteredHash
func MatchRegisteredHash(expected string) GomegaMatcher {
	return &HashMatcher{
		expectedHash: expected,
	}
}

func (m *HashMatcher) Match(actual interface{}) (bool, error) {
	actual, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("matcher expected a string value (actual: '%v')", actual)
	}

	return m.expectedHash == actual, nil
}

func (m *HashMatcher) report(negated bool, expected, actual string) string {
	builder := strings.Builder{}
	not := lo.Ternary(negated, " NOT ", " ")
	builder.WriteString(
		fmt.Sprintf("🔥 Expected hashes%vto be equal\n", not),
	)

	builder.WriteString(
		fmt.Sprintf("===> [🤖]  EXPECTED-HASH: '%v'\n", expected),
	)
	builder.WriteString(
		fmt.Sprintf("===> [👾]    ACTUAL-HASH: '%v'\n", actual),
	)

	return builder.String()
}

func (m *HashMatcher) FailureMessage(actual interface{}) string {
	actualHash, ok := actual.(string)
	if !ok {
		return fmt.Sprintf("matcher expected a string value (actual: '%v')", actual)
	}

	return m.report(false, m.expectedHash, actualHash)
}

func (m *HashMatcher) NegatedFailureMessage(actual interface{}) string {
	actualHash, ok := actual.(string)
	if !ok {
		return fmt.Sprintf("matcher expected a string value (actual: '%v')", actual)
	}

	return m.report(true, m.expectedHash, actualHash)
}
