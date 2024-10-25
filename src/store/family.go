package store

import (
	"strings"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/internal/third/lo"
	"github.com/spf13/pflag"
)

type LongFlagName = string
type ShortFlagName = string

type FlagDefinitions map[LongFlagName]ShortFlagName

var ShortFlags = FlagDefinitions{
	"cpu":        "", // worker pool family
	"now":        "",
	"dry-run":    "D", // preview family
	"files":      "F", // filter family
	"files-gb":   "G",
	"files-rx":   "X",
	"folders-gb": "Z",
	"folders-rx": "Y",
	"profile":    "P", // profile family
	"scheme":     "S",
	"language":   "", // i18n family
	"depth":      "", // depth family
	"no-recurse": "N",
	"sample":     "", // sampling
	"no-files":   "",
	"no-folders": "",
	"last":       "",
	"tui":        "", // interaction
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
