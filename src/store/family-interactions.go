package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/li18ngo"
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
			li18ngo.Text(locale.TextualInteractionIsNoTUIUsageTemplData{}),
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
			li18ngo.Text(locale.CliInteractionIsTUIUsageTemplData{}),
			defIsTUI,
			flagSet...,
		),
		&parent.Native.IsTUI,
	)
}
