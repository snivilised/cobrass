package translate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"golang.org/x/text/language"

	xi18n "github.com/snivilised/extendio/i18n"

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
// other than the one automatically detected.
// TODO: rename to just Use
func UseTag(tag language.Tag) error {
	// TODO: delegate to extendio.Use
	//
	_, found := lo.Find(languages.Supported, func(t language.Tag) bool {
		return t == tag
	})

	if found {
		languages = createIncrementalLanguageInfo(tag, languages)
		localiser = createLocaliser(languages)
	} else {
		return xi18n.NewLanguageNotAvailableNativeError(tag)
	}

	return nil
}

// GetLanguageInfo gets LanguageInfo.
func GetLanguageInfo() *LanguageInfo {
	return languages
}

// GetLocaliser gets the current go-i18n localizer instance.
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
