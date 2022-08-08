package translate

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localisable interface {
	Message() *i18n.Message
}

// This file will be scanned by goi18n extract in order to create translations file(s)
// The strings from i18n.Message{} are the values that are extracted for translation.
//

// --- language not supported.
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

// --- already exists, invalid enum info specified.
//
type EnumValueValueAlreadyExistsTemplData struct {
	Number int
	Value  string
}

func (td EnumValueValueAlreadyExistsTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "enum-value-already-exists.cobrass",
		Description: "Invalid enum value",
		Other:       "'{{.Value}} ({{.Number}})' already exists, invalid enum info specified",
	}
}

// --- is not a valid enum value
//

type IsNotValidEnumValueTemplData struct {
	Source string
}

func (td IsNotValidEnumValueTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "is-not-valid-enum-value.cobrass",
		Description: "Not a valid enum value",
		Other:       "'{{.Source}}' is not a valid enum value",
	}
}

// --- failed to add validator for flag, because it already exists.
//
type FailedToAddValidatorAlreadyExistsTemplData struct {
	Flag string
}

func (td FailedToAddValidatorAlreadyExistsTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "failed-to-add-validator-already-exists.cobrass",
		Description: "Failed to add validator for flag because it already exists",
		Other:       "failed to add validator for flag: '{{.Flag}}', because it already exists",
	}
}

// --- command already registered.
//
type CommandAlreadyRegisteredTemplData struct {
	Name string
}

func (td CommandAlreadyRegisteredTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "command-already-registered.cobrass",
		Description: "Cobra command already registered in container",
		Other:       "'cobra container: command '{{.Name}}' already registered",
	}
}

// --- parent command not registered
//

type ParentCommandNotRegisteredTemplData struct {
	Parent string
}

func (td ParentCommandNotRegisteredTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "parent-command-not-registered.cobrass",
		Description: "Parent Cobra command not registered in container",
		Other:       "cobra container: parent command '{{.Parent}}' not registered",
	}
}

// --- param set already registered.
//
type ParamSetAlreadyRegisteredTemplData struct {
	Name string
}

func (td ParamSetAlreadyRegisteredTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "param-set-already-registered.cobrass",
		Description: "Param set already registered in container",
		Other:       "parameter set '{{.Name}}' already registered",
	}
}

// --- param set must be struct.
//
type ParamSetObjectMustBeStructTemplData struct {
	Name string
	Type string
}

func (td ParamSetObjectMustBeStructTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "param-set-object-must-be-struct.cobrass",
		Description: "The native Parameter Set object denoted, must be a struct.",
		Other:       "the native param set object ('{{.Name}}') must be a struct, actual type: '{{.Type}}'",
	}
}

// --- param set must be pointer.
//
type ParamSetObjectMustBePointerTemplData struct {
	Name string
	Type string
}

func (td ParamSetObjectMustBePointerTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "param-set-object-must-be-pointer.cobrass",
		Description: "The native Parameter Set object denoted, must be a pointer.",
		Other:       "the native param set object ('{{.Name}}') must be a pointer, actual type: '{{.Type}}'",
	}
}

// --- param set not found.
//
type ParamSetNotFoundTemplData struct {
	Name string
}

func (td ParamSetNotFoundTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "param-set-not-found.cobrass",
		Description: "The parameter set not found.",
		Other:       "parameter set '{{.Name}}' not found",
	}
}

// --- Within.
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

// --- ContainmentOV.
//
type ContainmentOV[T any] struct {
	Flag       string
	Value      any
	Collection []T
}

type ContainsOptValidationTemplData[T any] struct {
	ContainmentOV[T]
}

func (td ContainsOptValidationTemplData[T]) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-contains.cobrass",
		Description: "'Contains' Option validation has failed due to Value not being a member of collection.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not a member of: [{{.Collection}}]",
	}
}

type NotContainsOptValidationTemplData[T any] struct {
	ContainmentOV[T]
}

func (td NotContainsOptValidationTemplData[T]) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-not-contains.cobrass",
		Description: "'Contains' Option validation has failed due to Value being a member of collection.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', is a member of: [{{.Collection}}]",
	}
}

// --- Match.
//
type MatchOV struct {
	Flag    string
	Value   string
	Pattern string
}

type MatchOptValidationTemplData struct {
	MatchOV
}

func (td MatchOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-match.cobrass",
		Description: "'Match' Option validation has failed due to Value not matching the regex pattern.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', does not match: [{{.Pattern}}]",
	}
}

type NotMatchOptValidationTemplData struct {
	MatchOV
}

func (td NotMatchOptValidationTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-not-match.cobrass",
		Description: "'Match' Option validation has failed due to Value matching the regex pattern.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', matches: [{{.Pattern}}]",
	}
}

// --- Relational.
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
