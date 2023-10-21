package store

import (
	"strings"

	"github.com/snivilised/cobrass/src/assistant"
)

type longFlagName = string
type shortFlagName = string

type flagDefinitions map[longFlagName]shortFlagName

var shortFlags = flagDefinitions{
	// worker pool family
	//
	"cpu": "C",
	"now": "N",

	// preview family
	//
	"dry-run": "D",

	// filter family
	//
	"files-gb":   "G",
	"files-rx":   "X",
	"folders-gb": "z",
	"folders-rx": "y",

	// parameter profile
	//
	"profile": "P",
}

func newFlagInfo[T any](usage string, defaultValue T) *assistant.FlagInfo {
	name := strings.Split(usage, " ")[0]
	short := shortFlags[name]

	return assistant.NewFlagInfo(usage, short, defaultValue)
}
