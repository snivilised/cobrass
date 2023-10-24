package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type ProfileParameterSet struct {
	Profile string
}

func (f *ProfileParameterSet) BindAll(
	parent *assistant.ParamSet[ProfileParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	const (
		defaultProfile = ""
	)

	parent.BindValidatedStringIsMatch(
		resolveNewFlagInfo(
			xi18n.Text(i18n.ProfileParamUsageTemplData{}),
			defaultProfile,
			flagSet...,
		),
		&parent.Native.Profile,
		`^[\w-]+$`,
	)
}
