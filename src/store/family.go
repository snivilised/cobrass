package store

import (
	"strings"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/pflag"
)

type LongFlagName = string
type ShortFlagName = string

type FlagDefinitions map[LongFlagName]ShortFlagName

var ShortFlags = FlagDefinitions{
	// worker pool family
	//
	"cpu": "",
	"now": "",

	// preview family
	//
	"dry-run": "D",

	// filter family
	//
	"files":      "F",
	"files-gb":   "G",
	"files-rx":   "X",
	"folders-gb": "Z",
	"folders-rx": "Y",

	// profile family
	//
	"profile": "P",
	"scheme":  "S",

	// i18n family (niche option, so let's define without a short code)
	//
	"language": "",

	// depth family
	//
	"depth":      "",
	"no-recurse": "N",
}

func newFlagInfo[T any](usage string, defaultValue T) *assistant.FlagInfo {
	name := strings.Split(usage, " ")[0]
	short := ShortFlags[name]

	return assistant.NewFlagInfo(usage, short, defaultValue)
}

func newFlagInfoOnFlagSet[T any](usage string, defaultValue T,
	alternativeFlagSet *pflag.FlagSet,
) *assistant.FlagInfo {
	name := strings.Split(usage, " ")[0]
	short := ShortFlags[name]

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
