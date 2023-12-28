package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type DepthParameterSet struct {
	Depth uint
	Skim  bool
}

func (f *DepthParameterSet) BindAll(
	parent *assistant.ParamSet[DepthParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --depth
	//
	const (
		defaultDepth = uint(0)
	)

	parent.BindUint(
		resolveNewFlagInfo(
			xi18n.Text(i18n.DepthDepthParamUsageTemplData{}),
			defaultDepth,
			flagSet...,
		),
		&parent.Native.Depth,
	)

	// --skim(K)
	//
	const (
		defaultSkim = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(i18n.DepthSkimParamUsageTemplData{}),
			defaultSkim,
			flagSet...,
		),
		&parent.Native.Skim,
	)

	parent.Command.MarkFlagsMutuallyExclusive("depth", "skim")
}
