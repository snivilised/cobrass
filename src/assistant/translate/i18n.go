package translate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"golang.org/x/text/language"

	"github.com/snivilised/cobrass/src/assistant/internal/l10n"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var languages *LanguageInfo
var localiser *i18n.Localizer

type LanguageInitOptions struct {
	// TranslationFilename overrides the default filename to load for a language
	// When not set, the default filename is active.<ietf language>.json
	//
	TranslationFilename string

	// Path denoting install location
	//
	Path string

	// Detected language tag
	//
	Detected language.Tag
}

// ValidatorContainerOptionFn definition of a client defined function to
// set ValidatorContainer options.
//
type LanguageInitOptionFn func(*LanguageInitOptions)

func Initialise(options ...LanguageInitOptionFn) {
	o := LanguageInitOptions{}
	for _, fo := range options {
		fo(&o)
	}

	languages = createInitialLanguageInfo(o)
	localiser = createLocaliser(languages)
}

// the active file should be in the same directory at the item that is
// loading the bundle
//
// Create the "active.en.json" file from internal/i18n:
// cd internal/i18n
// goi18n extract -format json

// do merge
// goi18n merge -outdir out -format json active.en.json translate.en-US.json
//
// do rename out/translate.en-US.json -> out/active.en-US.json (or copy the text content over)
// mv translate.en-US.json active.en-US.json
//

// LanguageInfo indicates information relating to current language. See members for
// details.
//
type LanguageInfo struct {
	// TranslationFilename overrides the default filename to load for a language
	// When not set, the default filename is active.<ietf language>.json
	//
	TranslationFilename string

	// Path denoting where to load language file from
	//
	Path string

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
		return fmt.Errorf(GetLanguageNotSupportedErrorMessage(tag))
	}

	return nil
}

// GetLanguageInfo gets LanguageInfo.
//
func GetLanguageInfo() *LanguageInfo {
	return languages
}

// GetLocaliser gets the current go-i18n localizer instance.
//
func GetLocaliser() *i18n.Localizer {
	return localiser
}

func createInitialLanguageInfo(options LanguageInitOptions) *LanguageInfo {
	return &LanguageInfo{
		TranslationFilename: options.TranslationFilename,
		Path:                options.Path,
		Default:             language.BritishEnglish,
		Detected:            options.Detected,
		Current:             options.Detected,

		// TODO: this has to be read in from config
		//
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createIncrementalLanguageInfo(requested language.Tag, existing *LanguageInfo) *LanguageInfo {
	return &LanguageInfo{
		TranslationFilename: existing.TranslationFilename, // ???
		Path:                existing.Path,
		Default:             language.BritishEnglish,
		Detected:            existing.Detected,
		Current:             requested,
		Supported:           existing.Supported,
	}
}

func createLocaliser(li *LanguageInfo) *i18n.Localizer {
	bundle := i18n.NewBundle(li.Current)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	const Prefix = "cobrass."

	if li.Current != li.Default {
		filename := lo.Ternary(li.TranslationFilename != "",
			li.TranslationFilename, fmt.Sprintf("%vactive.%v.json", Prefix, li.Current))

		exe, _ := os.Executable()
		resolved, _ := filepath.Abs(li.Path)
		directory := lo.Ternary(li.Path != "", resolved, filepath.Dir(exe))
		path := filepath.Join(directory, filename)

		_, err := bundle.LoadMessageFile(path)

		if err != nil {
			// Since, translations failed to load, we will ever be in a situation where
			// this error message is able to be generated in translated form, so
			// we are force to generated an error message in the default language.
			//
			panic(fmt.Errorf("could not load translations for '%v', from: '%v'", li.Current, path))
		}
	}

	supported := lo.Map(li.Supported, func(t language.Tag, _ int) string {
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

// --- language not supported

func GetLanguageNotSupportedErrorMessage(tag language.Tag) string {
	data := l10n.LanguageNotSupportedTemplData{
		Language: tag.String(),
	}

	return localise(data)
}

// --- already exists, invalid enum info specified

func GetEnumValueAlreadyExistsErrorMessage(value string, number int) string {
	data := l10n.EnumValueValueAlreadyExistsTemplData{
		Value:  value,
		Number: number,
	}

	return localise(data)
}

// --- is not a valid enum value

func GetIsNotValidEnumValueErrorMessage(source string) string {
	data := l10n.IsNotValidEnumValueTemplData{
		Source: source,
	}

	return localise(data)
}

// --- failed to add validator for flag, because it already exists

func GetFailedToGetValidatorForFlagAlreadyExistsErrorMessage(flag string) string {
	data := l10n.FailedToAddValidatorAlreadyExistsTemplData{
		Flag: flag,
	}

	return localise(data)
}

// --- command already registered

func GetCommandAlreadyRegisteredErrorMessage(name string) string {
	data := l10n.CommandAlreadyRegisteredTemplData{
		Name: name,
	}

	return localise(data)
}

// --- parent command not registered

func GetParentCommandNotRegisteredErrorMessage(parent string) string {
	data := l10n.CommandAlreadyRegisteredTemplData{
		Name: parent,
	}

	return localise(data)
}

// --- param set already registered

func GetParamSetAlreadyRegisteredErrorMessage(name string) string {
	data := l10n.ParamSetAlreadyRegisteredTemplData{
		Name: name,
	}

	return localise(data)
}

// --- param set must be struct

func GetParamSetMustBeStructErrorMessage(name string, actualType string) string {
	data := l10n.ParamSetObjectMustBeStructTemplData{
		Name: name,
		Type: actualType,
	}

	return localise(data)
}

// --- param set must be pointer

func GetParamSetMustBePointerErrorMessage(name, actualType string) string {
	data := l10n.ParamSetObjectMustBePointerTemplData{
		Name: name,
		Type: actualType,
	}

	return localise(data)
}

// --- param set not found

func GetParamSetNotFoundErrorMessage(name string) string {
	data := l10n.ParamSetNotFoundTemplData{
		Name: name,
	}

	return localise(data)
}

// --- Within

func GetWithinErrorMessage(flag string, value, lo, hi any) string {
	data := l10n.WithinOptValidationTemplData{
		OutOfRangeOV: l10n.OutOfRangeOV{
			Flag: flag, Value: value, Lo: lo, Hi: hi,
		},
	}

	return localise(data)
}

func GetNotWithinErrorMessage(flag string, value, lo, hi any) string {
	data := l10n.NotWithinOptValidationTemplData{
		OutOfRangeOV: l10n.OutOfRangeOV{
			Flag: flag, Value: value, Lo: lo, Hi: hi,
		},
	}

	return localise(data)
}

// --- Containment

func GetContainsErrorMessage[T any](flag string, value T, collection []T) string {
	data := l10n.ContainsOptValidationTemplData[T]{
		ContainmentOV: l10n.ContainmentOV[T]{
			Flag: flag, Value: value, Collection: collection,
		},
	}

	return localise(data)
}

func GetNotContainsErrorMessage[T any](flag string, value T, collection []T) string {
	data := l10n.NotContainsOptValidationTemplData[T]{
		ContainmentOV: l10n.ContainmentOV[T]{
			Flag: flag, Value: value, Collection: collection,
		},
	}

	return localise(data)
}

// --- Match

func GetMatchErrorMessage(flag string, value string, pattern string) string {
	data := l10n.MatchOptValidationTemplData{
		MatchOV: l10n.MatchOV{
			Flag: flag, Value: value, Pattern: pattern,
		},
	}

	return localise(data)
}

func GetNotMatchErrorMessage(flag string, value string, pattern string) string {
	data := l10n.NotMatchOptValidationTemplData{
		MatchOV: l10n.MatchOV{
			Flag: flag, Value: value, Pattern: pattern,
		},
	}

	return localise(data)
}

// --- Relational

func GetGreaterThanErrorMessage(flag string, value, threshold any) string {
	data := l10n.GreaterThanOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}

	return localise(data)
}

func GetAtLeastErrorMessage(flag string, value, threshold any) string {
	data := l10n.AtLeastOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}

	return localise(data)
}

func GetLessThanErrorMessage(flag string, value, threshold any) string {
	data := l10n.LessThanOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}

	return localise(data)
}

func GetAtMostErrorMessage(flag string, value, threshold any) string {
	data := l10n.AtMostOptValidationTemplData{
		RelationalOV: l10n.RelationalOV{
			Flag: flag, Value: value, Threshold: threshold,
		},
	}

	return localise(data)
}
