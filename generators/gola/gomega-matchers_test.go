package gola_test

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/generators/gola"

	. "github.com/onsi/gomega/types"
)

type (
	AreAllSourceCodeFilesPresentMatcher struct {
		directory string
	}
)

func ContainAllSourceCodeFilesAt(directory string) GomegaMatcher {
	return &AreAllSourceCodeFilesPresentMatcher{
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
	not := lo.Ternary(negated, "NOT ", " ")
	builder.WriteString(
		fmt.Sprintf("🔥 Expected all source code files %vto be present\n", not),
	)

	sourceCode.ForEach(func(entry *gola.SourceCodeData) {
		exists := entry.Exists()
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
