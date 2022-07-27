package assistant

import (
	"encoding/json"
	"fmt"

	"github.com/cubiest/jibberjabber"
	"github.com/samber/lo"
	"golang.org/x/text/language"

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

// LanguageInfo indicates information relating to current language. See members for
// details.
//
type LanguageInfo struct {
	// Default language reflects the base language. If all else fails, messages will
	// be in this language. It is fixed at BritishEnglish reflecting the language this
	// package is written in.
	//
	Default language.Tag

	// Detected is the language that is automatically detected of the host machine. Assuming
	// the ost machine is configured of the user's preference, there should be no other
	// reason to divert from this language.
	//
	Detected language.Tag

	// Territory reflects the region as automatically detected.
	//
	Territory string

	// Current reflects the language currently in force. Will by default be the detected
	// language. Client can change this with the UseTag function.
	//
	Current language.Tag

	// Supported indicates the list of languages for which translations are available.
	//
	Supported []language.Tag
}

// UseTag allows the client to change the language currently in use to a language
// othr than the one automatically detected.
//
func UseTag(tag language.Tag) error {
	_, found := lo.Find(languages.Supported, func(t language.Tag) bool {
		return t == tag
	})

	if found {
		languages = createIncrementalLanguageInfo(tag, languages)
		localiser = createLocaliser(languages)
	} else {
		return fmt.Errorf(getLanguageNotSupportedErrorMessage(tag))
	}

	return nil
}

// GetLanguageInfo gets LanguageInfo
//
func GetLanguageInfo() *LanguageInfo {
	return languages
}

// GetLocaliser gets the current go-i18n localizer instance
//
func GetLocaliser() *i18n.Localizer {
	return localiser
}

type detectInfo struct {
	tag       language.Tag
	territory string
}

var languages *LanguageInfo
var localiser *i18n.Localizer

func init() {
	languages = createInitialLanguageInfo()
	localiser = createLocaliser(languages)
}

func detect() *detectInfo {
	detectedLang, _ := jibberjabber.DetectLanguage()
	territory, _ := jibberjabber.DetectTerritory()

	detectedLangTag, _ := language.Parse(fmt.Sprintf("%v-%v", detectedLang, territory))

	return &detectInfo{
		tag:       detectedLangTag,
		territory: territory,
	}
}

func createInitialLanguageInfo() *LanguageInfo {
	// TODO: the supported list should NOT be static, client can pass this list
	// in
	//
	dInfo := detect()

	return &LanguageInfo{
		Default:   language.BritishEnglish,
		Detected:  dInfo.tag,
		Territory: dInfo.territory,
		Current:   dInfo.tag,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createIncrementalLanguageInfo(requested language.Tag, existing *LanguageInfo) *LanguageInfo {

	return &LanguageInfo{
		Default:   language.BritishEnglish,
		Detected:  existing.Detected,
		Territory: existing.Territory,
		Current:   requested,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createLocaliser(li *LanguageInfo) *i18n.Localizer {
	bundle := i18n.NewBundle(languages.Current)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("./internal/l10n/out/translate.en-US.json")

	supported := lo.Map(languages.Supported, func(t language.Tag, _ int) string {
		return t.String()
	})

	return i18n.NewLocalizer(bundle, supported...)
}

func localise(data l10n.Localisable) string {
	return localiser.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: data.Message(),
		TemplateData:   data,
	})
}

func getLanguageNotSupportedErrorMessage(tag language.Tag) string {
	data := l10n.LanguageNotSupportedTemplData{
		Language: tag.String(),
	}
	return localise(data)
}

// --- Within

func getWithinErrorMessage(flag string, value, lo, hi any) string {

	data := l10n.WithinOptValidationTemplData{
		OutOfRangeOV: l10n.OutOfRangeOV{
			Flag: flag, Value: value, Lo: lo, Hi: hi,
		},
	}
	return localise(data)
}

func getNotWithinErrorMessage(flag string, value, lo, hi any) string {

	data := l10n.NotWithinOptValidationTemplData{
		OutOfRangeOV: l10n.OutOfRangeOV{
			Flag: flag, Value: value, Lo: lo, Hi: hi,
		},
	}
	return localise(data)
}

// --- Containment

func getContainsErrorMessage[T any](flag string, value T, collection []T) string {

	data := l10n.ContainsOptValidationTemplData[T]{
		ContainmentOV: l10n.ContainmentOV[T]{
			Flag: flag, Value: value, Collection: collection,
		},
	}
	return localise(data)
}

func getNotContainsErrorMessage[T any](flag string, value T, collection []T) string {

	data := l10n.NotContainsOptValidationTemplData[T]{
		ContainmentOV: l10n.ContainmentOV[T]{
			Flag: flag, Value: value, Collection: collection,
		},
	}
	return localise(data)
}

// --- Match

func getMatchErrorMessage(flag string, value string, pattern string) string {

	data := l10n.MatchOptValidationTemplData{
		MatchOV: l10n.MatchOV{
			Flag: flag, Value: value, Pattern: pattern,
		},
	}
	return localise(data)
}

func getNotMatchErrorMessage(flag string, value string, pattern string) string {

	data := l10n.NotMatchOptValidationTemplData{
		MatchOV: l10n.MatchOV{
			Flag: flag, Value: value, Pattern: pattern,
		},
	}
	return localise(data)
}

// --- Relational

func getGreaterThanErrorMessage(flag string, value, threshold any) string {

	data := l10n.GreaterThanOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}
	return localise(data)
}

func getAtLeastErrorMessage(flag string, value, threshold any) string {

	data := l10n.AtLeastOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}
	return localise(data)
}

func getLessThanErrorMessage(flag string, value, threshold any) string {

	data := l10n.LessThanOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}
	return localise(data)
}

func getAtMostErrorMessage(flag string, value, threshold any) string {

	data := l10n.AtMostOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}
	return localise(data)
}
