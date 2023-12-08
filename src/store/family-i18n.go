package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
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
			xi18n.Text(i18n.LanguageParamUsageTemplData{}),
			defaultLanguage,
			flagSet...,
		),
		&parent.Native.Language,
		func(s string, f *pflag.Flag) error {
			if _, err := language.Parse(s); err != nil {
				return err
			}

			return nil
		},
	)
}
