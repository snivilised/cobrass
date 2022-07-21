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
	Printer   *message.Printer
}

var languages LanguageInfo

func init() {
	british := language.MustParse("en-GB")
	us := language.MustParse("en-US")

	detectedLang, err := jibberjabber.DetectLanguage()
	territory, _ := jibberjabber.DetectTerritory()

	detectedLangTag, _ := language.Parse(fmt.Sprintf("%v-%v", detectedLang, territory))

	current := lo.Ternary(err == nil, detectedLangTag, british)
	languages = LanguageInfo{
		Default:   british,
		Detected:  detectedLangTag,
		Territory: territory,
		Current:   current,
		Supported: []language.Tag{british, us},
		Printer:   message.NewPrinter(current),
	}
}

func UseTag(tag language.Tag) error {
	_, found := lo.Find(languages.Supported, func(t language.Tag) bool {
		return t == tag
	})

	if found {
		languages.Current = tag
		languages.Printer = message.NewPrinter(tag)
	} else {
		return fmt.Errorf("language '%v' not supported", tag)
	}

	return nil
}

func GetLanguageInfo() LanguageInfo {
	return languages
}
