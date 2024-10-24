package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/li18ngo"
	"github.com/spf13/pflag"
)

type ProfileParameterSet struct {
	Profile string
	Scheme  string
}

func (f *ProfileParameterSet) BindAll(
	parent *assistant.ParamSet[ProfileParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --profile(P)
	//
	const (
		defaultProfile = ""
	)

	parent.BindValidatedStringIsMatch(
		resolveNewFlagInfo(
			li18ngo.Text(locale.ProfileParamUsageTemplData{}),
			defaultProfile,
			flagSet...,
		),
		&parent.Native.Profile,
		`^[\w-]+$`,
	)

	// -- scheme(S)
	//
	const (
		defaultScheme = ""
	)

	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SchemeParamUsageTemplData{}),
			defaultScheme,
			flagSet...,
		),
		&parent.Native.Scheme,
	)
}
