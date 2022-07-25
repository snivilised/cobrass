package assistant

import (
	"fmt"

	"github.com/cubiest/jibberjabber"
	"github.com/samber/lo"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type LanguageInfo struct {
	Default   language.Tag
	Detected  language.Tag
	Territory string
	Current   language.Tag
	Supported []language.Tag
}

var languages LanguageInfo
var p *message.Printer

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

	p = message.NewPrinter(current)

	message.SetString(language.AmericanEnglish,
		"greetings '%v', welcome to internationalisation",
		"greetings '%v', welcome to internationalization",
	)
}

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
