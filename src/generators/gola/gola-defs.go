package gola

import (
	"bytes"

	"github.com/snivilised/cobrass/src/generators/gola/internal/collections"
)

type executionInfo struct {
	Spec      *TypeSpec
	Operators operatorCollection
}

type typeCollection = collections.OrderedKeysMap[TypeNameID, *TypeSpec]

type GeneratedPage interface {
	Name() CodeFileName
	Content() string
}

type generatedYield struct {
	name   CodeFileName
	buffer bytes.Buffer
}

func (y *generatedYield) Name() CodeFileName {
	return y.name
}

func (y *generatedYield) Content() string {
	return y.buffer.String()
}

type SignatureCounts struct {
	Type int
	Func int
}

type SignatureCountsBySource = collections.OrderedKeysMap[CodeFileName, *SignatureCounts]

type SignatureResult struct {
	Totals   *SignatureCounts
	Counters SignatureCountsBySource
	Status   string
	Hash     string
	Output   string
}

// CodeContent can be used for multiple purposes. The key thing is
// that each code file can have some content associated with. What
// that content is can vary depending on context.
type CodeContent = collections.OrderedKeysMap[CodeFileName, string]
