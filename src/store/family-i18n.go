package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/li18ngo"
	"github.com/spf13/pflag"
	"golang.org/x/text/language"
)

type I18nParameterSet struct {
	Language string
}

func (f *I18nParameterSet) BindAll(
	parent *assistant.ParamSet[I18nParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --language
	//
	const (
		defaultLanguage = ""
	)

	parent.BindValidatedString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.LanguageParamUsageTemplData{}),
			defaultLanguage,
			flagSet...,
		),
		&parent.Native.Language,
		func(s string, _ *pflag.Flag) error {
			if _, err := language.Parse(s); err != nil {
				return err
			}

			return nil
		},
	)
}
