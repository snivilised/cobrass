package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type TextualInteractionParameterSet struct {
	IsNoTui bool
}

func (f *TextualInteractionParameterSet) BindAll(
	parent *assistant.ParamSet[TextualInteractionParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --no-tui
	//
	const (
		defNoTUI = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(i18n.TextualInteractionIsNoTUIUsageTemplData{}),
			defNoTUI,
			flagSet...,
		),
		&parent.Native.IsNoTui,
	)
}

type CliInteractionParameterSet struct {
	IsTUI bool
}

func (f *CliInteractionParameterSet) BindAll(
	parent *assistant.ParamSet[CliInteractionParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --tui
	//
	const (
		defIsTUI = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(i18n.CliInteractionIsTUIUsageTemplData{}),
			defIsTUI,
			flagSet...,
		),
		&parent.Native.IsTUI,
	)
}
