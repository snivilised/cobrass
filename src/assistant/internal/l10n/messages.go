package l10n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// must be struct
//
type PsObjectMustBeStructTemplData struct {
	Name string
}

func (td PsObjectMustBeStructTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ps-object-must-be-struct.cobrass",
		Description: "The native Parameter Set object denoted, must be a struct.",
		Other:       "the native param set object ('{{.Name}}') must be a struct",
	}
}

// language not supported
//
type LanguageNotSupportedTemplData struct {
	Language string
}

func (td LanguageNotSupportedTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "language-not-supported.cobrass",
		Description: "The language specified is not supported; no translations for this language.",
		Other:       "language '{{.Language}}' not supported",
	}
}

// This file will be scanned by goi18n extract in order to create translations file(s)
// The strings from i18n.Message{} are the values that are extracted for translation.
//

// --- Out Of Range
//
type OutOfRangeOV struct {
	Flag  string
	Value any
	Lo    any
	Hi    any
}

type WithinOptValidationTemplData struct {
	OutOfRangeOV
}

func (td WithinOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-within.cobrass",
		Description: "'Within' Option validation has failed due to Value being out of range.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', out of range: [{{.Lo}}]..[{{.Hi}}]",
	}
}

type NotWithinOptValidationTemplData struct {
	OutOfRangeOV
}

func (td NotWithinOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-not-within.cobrass",
		Description: "'Within' Option validation has failed due to Value being inside of range.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', is inside of range: [{{.Lo}}]..[{{.Hi}}]",
	}
}

// --- ContainmentOV
//
type ContainmentOV struct {
	Flag       string
	Value      any
	Collection []any
}

type ContainsOptValidationTemplData struct {
	ContainmentOV
}

func (td ContainsOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-contains.cobrass",
		Description: "'Contains' Option validation has failed due to Value not being a member of collection.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not a member of: [{{.Collection}}]",
	}
}

type NotContainsOptValidationTemplData struct {
	ContainmentOV
}

func (td NotContainsOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-not-contains.cobrass",
		Description: "'Contains' Option validation has failed due to Value being a member of collection.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', is a member of: [{{.Collection}}]",
	}
}

// --- Relational
//
type RelationalOV struct {
	Flag      string
	Value     any
	Threshold any
}

type GreaterThanOptValidationTemplData struct {
	RelationalOV
}

func (td GreaterThanOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-greater-than.cobrass",
		Description: "'GreaterThan' Option validation has failed due to Value not being greater than threshold.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not greater than: [{{.Threshold}}]",
	}
}

// this is just an example of how to instantiate a struct literal
// with embedded type
//
var GT = GreaterThanOptValidationTemplData{
	RelationalOV: RelationalOV{"flag", 1, 2},
}

type AtLeastOptValidationTemplData struct {
	RelationalOV
}

func (td AtLeastOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-at-least.cobrass",
		Description: "'AtLeast' Option validation has failed due to Value not being at least the threshold.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not at least: [{{.Threshold}}]",
	}
}

type LessThanOptValidationTemplData struct {
	RelationalOV
}

func (td LessThanOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-less-than.cobrass",
		Description: "'LessThan' Option validation has failed due to Value not being less than the threshold.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not less than: [{{.Threshold}}]",
	}
}

type AtMostOptValidationTemplData struct {
	RelationalOV
}

func (td AtMostOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-at-most.cobrass",
		Description: "'AtMost' Option validation has failed due to Value not being at most the threshold.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not at most: [{{.Threshold}}]",
	}
}
