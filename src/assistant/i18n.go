package assistant

import (
	"encoding/json"
	"fmt"

	"github.com/cubiest/jibberjabber"
	"github.com/samber/lo"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/snivilised/cobrass/src/assistant/internal/l10n"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// the active file should be in the same directory at the item that is
// loading the bundle
//
// Create the "active.en.json" file from internal/i18n:
// cd internal/i18n
// goi18n extract -format json

// do merge
// goi18n merge -outdir out -format json active.en.json translate.en-US.json
//

type LanguageInfo struct {
	Default   language.Tag
	Detected  language.Tag
	Territory string
	Current   language.Tag
	Supported []language.Tag
}

var languages LanguageInfo
var p *message.Printer
var Localiser *i18n.Localizer

func init() {
	detectedLang, err := jibberjabber.DetectLanguage()
	territory, _ := jibberjabber.DetectTerritory()

	detectedLangTag, _ := language.Parse(fmt.Sprintf("%v-%v", detectedLang, territory))

	current := lo.Ternary(err == nil, detectedLangTag, language.BritishEnglish)
	languages = LanguageInfo{
		Default:   language.BritishEnglish,
		Detected:  detectedLangTag,
		Territory: territory,
		Current:   current,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}

	// TODO: printer is not used, delete this
	//
	p = message.NewPrinter(current)

	bundle := i18n.NewBundle(languages.Default)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("./internal/l10n/out/translate.en-US.json")

	supported := lo.Map(languages.Supported, func(t language.Tag, _ int) string {
		return t.String()
	})
	Localiser = i18n.NewLocalizer(bundle, supported...)
}

// not necessry when using go-i18n
//
func UseTag(tag language.Tag) error {
	_, found := lo.Find(languages.Supported, func(t language.Tag) bool {
		return t == tag
	})

	if found {
		languages.Current = tag
		p = message.NewPrinter(tag)
	} else {
		return fmt.Errorf("language '%v' not supported", tag)
	}

	return nil
}

func GetLanguageInfo() LanguageInfo {
	return languages
}

// do not use this function, it is a temporary function, to help define
// correct way of doing i18n with go-i18n
//
func GetOutOfRangeErrorMessage(flag string, value, lo, hi any) string {

	data := l10n.WithinOptValidationTemplData{
		OutOfRangeOV: l10n.OutOfRangeOV{
			Flag: flag, Value: value, Lo: lo, Hi: hi,
		},
	}
	return Localiser.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: data.Message(),
		TemplateData:   data,
	})
}
