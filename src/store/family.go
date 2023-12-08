package store

import (
	"strings"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/pflag"
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
	"folders-gb": "Z",
	"folders-rx": "Y",

	// profile family
	//
	"profile": "P",
	"scheme":  "S",
}

func newFlagInfo[T any](usage string, defaultValue T) *assistant.FlagInfo {
	name := strings.Split(usage, " ")[0]
	short := shortFlags[name]

	return assistant.NewFlagInfo(usage, short, defaultValue)
}

func newFlagInfoOnFlagSet[T any](usage string, defaultValue T,
	alternativeFlagSet *pflag.FlagSet,
) *assistant.FlagInfo {
	name := strings.Split(usage, " ")[0]
	short := shortFlags[name]

	return assistant.NewFlagInfoOnFlagSet(usage, short, defaultValue, alternativeFlagSet)
}

func resolveNewFlagInfo[T any](usage string, defaultValue T,
	alternativeFlagSet ...*pflag.FlagSet,
) *assistant.FlagInfo {
	return lo.TernaryF(len(alternativeFlagSet) == 0,
		func() *assistant.FlagInfo {
			return newFlagInfo(usage, defaultValue)
		},
		func() *assistant.FlagInfo {
			return newFlagInfoOnFlagSet(usage, defaultValue, alternativeFlagSet[0])
		},
	)
}
