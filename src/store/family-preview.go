package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type PreviewParameterSet struct {
	DryRun bool
}

func (f *PreviewParameterSet) BindAll(
	parent *assistant.ParamSet[PreviewParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --dry-run(D)
	//
	const (
		defaultDryRun = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(locale.DryRunParamUsageTemplData{}),
			defaultDryRun,
			flagSet...,
		),
		&parent.Native.DryRun,
	)
}
