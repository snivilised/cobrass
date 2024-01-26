package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type CascadeParameterSet struct {
	Depth     uint
	NoRecurse bool
}

func (f *CascadeParameterSet) BindAll(
	parent *assistant.ParamSet[CascadeParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --depth
	//
	const (
		defaultDepth = uint(0)
	)

	parent.BindUint(
		resolveNewFlagInfo(
			xi18n.Text(i18n.CascadeDepthParamUsageTemplData{}),
			defaultDepth,
			flagSet...,
		),
		&parent.Native.Depth,
	)

	// --no-recurse(N)
	//
	const (
		defaultNoRecurse = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(i18n.CascadeNoRecurseParamUsageTemplData{}),
			defaultNoRecurse,
			flagSet...,
		),
		&parent.Native.NoRecurse,
	)

	parent.Command.MarkFlagsMutuallyExclusive("depth", "no-recurse")
}
