package i18n

import (
	"github.com/snivilised/extendio/i18n"
)

// These are user facing errors messages that occur due to
// incorrect use of the cli application. They occur as a result
// of validating the user provided options on the command line.
//

// 💧 OutOfRangeOV
type OutOfRangeOV struct {
	CobrassTemplData
	Flag  string
	Value any
	Lo    any
	Hi    any
}

// ❌ WithinOptValidationTemplData

// WithinOptValidationTemplData
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

type WithinOptValidationBehaviourQuery interface {
	error
	IsOutOfRange() bool
}

type WithinOptValidation struct {
	i18n.LocalisableError
}

func (e WithinOptValidation) IsOutOfRange() bool {
	return true
}

func NewWithinOptValidationError(flag string, value, low, high any) WithinOptValidationBehaviourQuery {
	return &WithinOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: WithinOptValidationTemplData{
				OutOfRangeOV: OutOfRangeOV{
					Flag:  flag,
					Value: value,
					Lo:    low,
					Hi:    high,
				},
			},
		},
	}
}

// ❌ NotWithinOptValidationTemplData

// NotWithinOptValidationTemplData
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

type NotWithinOptValidationBehaviourQuery interface {
	error
	IsInsideOfRange() bool
}

type NotWithinOptValidation struct {
	i18n.LocalisableError
}

func (e NotWithinOptValidation) IsInsideOfRange() bool {
	return true
}

func NewNotWithinOptValidationError(flag string, value, low, high any) NotWithinOptValidationBehaviourQuery {
	return &NotWithinOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: NotWithinOptValidationTemplData{
				OutOfRangeOV: OutOfRangeOV{
					Flag:  flag,
					Value: value,
					Lo:    low,
					Hi:    high,
				},
			},
		},
	}
}

// 💧 ContainmentOV

// ContainmentOV
type ContainmentOV[T any] struct {
	CobrassTemplData
	Flag       string
	Value      any
	Collection []T
}

// ❌ ContainsOptValidationTemplData

// ContainsOptValidationTemplData
type ContainsOptValidationTemplData[T any] struct {
	ContainmentOV[T]
}

func (td ContainsOptValidationTemplData[T]) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-contains.cobrass",
		Description: "'Contains' Option validation has failed due to Value not being a member of collection.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', not a member of: {{.Collection}}",
	}
}

type ContainsOptValidationBehaviourQuery[T any] interface {
	error
	IsAMemberOf() bool
}

type ContainsOptValidation[T any] struct {
	i18n.LocalisableError
}

func (e ContainsOptValidation[T]) IsAMemberOf() bool {
	return true
}

func NewContainsOptValidationError[T any](flag string, value any, collection []T) ContainsOptValidationBehaviourQuery[T] {
	return &ContainsOptValidation[T]{
		LocalisableError: i18n.LocalisableError{
			Data: ContainsOptValidationTemplData[T]{
				ContainmentOV: ContainmentOV[T]{
					Flag:       flag,
					Value:      value,
					Collection: collection,
				},
			},
		},
	}
}

// ❌ NotContainsOptValidationTemplData

// NotContainsOptValidationTemplData
type NotContainsOptValidationTemplData[T any] struct {
	ContainmentOV[T]
}

func (td NotContainsOptValidationTemplData[T]) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "ov-failed-not-contains.cobrass",
		Description: "'Contains' Option validation has failed due to Value being a member of collection.",
		Other:       "({{.Flag}}): option validation failed, '{{.Value}}', is a member of: {{.Collection}}",
	}
}

type NotContainsOptValidationBehaviourQuery[T any] interface {
	error
	IsNotAMemberOf() bool
}

type NotContainsOptValidation[T any] struct {
	i18n.LocalisableError
}

func (e NotContainsOptValidation[T]) IsNotAMemberOf() bool {
	return true
}

func NewNotContainsOptValidationError[T any](flag string, value any, collection []T) NotContainsOptValidationBehaviourQuery[T] {
	return &NotContainsOptValidation[T]{
		LocalisableError: i18n.LocalisableError{
			Data: NotContainsOptValidationTemplData[T]{
				ContainmentOV: ContainmentOV[T]{
					Flag:       flag,
					Value:      value,
					Collection: collection,
				},
			},
		},
	}
}

// 💧 Match

// MatchOV
type MatchOV struct {
	CobrassTemplData
	Flag    string
	Value   string
	Pattern string
}

// ❌ MatchOptValidationTemplData

// MatchOptValidationTemplData
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

type MatchOptValidationBehaviourQuery interface {
	error
	IsMatch() bool
}

type MatchOptValidation struct {
	i18n.LocalisableError
}

func (e MatchOptValidation) IsMatch() bool {
	return true
}

func NewMatchOptValidationError(flag, value, pattern string) MatchOptValidationBehaviourQuery {
	return &MatchOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: MatchOptValidationTemplData{
				MatchOV: MatchOV{
					Flag:    flag,
					Value:   value,
					Pattern: pattern,
				},
			},
		},
	}
}

// ❌ NotMatchOptValidationTemplData

// NotMatchOptValidationTemplData
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

type NotMatchOptValidationBehaviourQuery interface {
	error
	IsNotMatch() bool
}

type NotMatchOptValidation struct {
	i18n.LocalisableError
}

func (e NotMatchOptValidation) IsNotMatch() bool {
	return true
}

func NewNotMatchOptValidationError(flag, value, pattern string) NotMatchOptValidationBehaviourQuery {
	return &NotMatchOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: NotMatchOptValidationTemplData{
				MatchOV: MatchOV{
					Flag:    flag,
					Value:   value,
					Pattern: pattern,
				},
			},
		},
	}
}

// 💧 Relational
type RelationalOV struct {
	CobrassTemplData
	Flag      string
	Value     any
	Threshold any
}

// ❌ GreaterThanOptValidationTemplData

// GreaterThanOptValidationTemplData
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

type GreaterThanOptValidationBehaviourQuery interface {
	error
	IsGreaterThan() bool
}

type GreaterThanOptValidation struct {
	i18n.LocalisableError
}

func (e GreaterThanOptValidation) IsGreaterThan() bool {
	return true
}

func NewGreaterThanOptValidationError(flag string, value, threshold any) GreaterThanOptValidationBehaviourQuery {
	return &GreaterThanOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: GreaterThanOptValidationTemplData{
				RelationalOV: RelationalOV{
					Flag:      flag,
					Value:     value,
					Threshold: threshold,
				},
			},
		},
	}
}

// ❌ AtLeastOptValidationTemplData

// AtLeastOptValidationTemplData
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

type AtLeastOptValidationBehaviourQuery interface {
	error
	IsAtLeast() bool
}

type AtLeastOptValidation struct {
	i18n.LocalisableError
}

func (e AtLeastOptValidation) IsAtLeast() bool {
	return true
}

func NewAtLeastOptValidationError(flag string, value, threshold any) AtLeastOptValidationBehaviourQuery {
	return &AtLeastOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: AtLeastOptValidationTemplData{
				RelationalOV: RelationalOV{
					Flag:      flag,
					Value:     value,
					Threshold: threshold,
				},
			},
		},
	}
}

// ❌ LessThanOptValidationTemplData

// LessThanOptValidationTemplData
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

type LessThanOptValidationBehaviourQuery interface {
	error
	IsLessThan() bool
}

type LessThanOptValidation struct {
	i18n.LocalisableError
}

func (e LessThanOptValidation) IsLessThan() bool {
	return true
}

func NewLessThanOptValidationError(flag string, value, threshold any) LessThanOptValidationBehaviourQuery {
	return &LessThanOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: AtLeastOptValidationTemplData{
				RelationalOV: RelationalOV{
					Flag:      flag,
					Value:     value,
					Threshold: threshold,
				},
			},
		},
	}
}

// ❌ AtMostOptValidationTemplData

// AtMostOptValidationTemplData
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

type AtMostOptValidationBehaviourQuery interface {
	error
	IsAtMost() bool
}

type AtMostOptValidation struct {
	i18n.LocalisableError
}

func (e AtMostOptValidation) IsAtMost() bool {
	return true
}

func NewAtMostOptValidationError(flag string, value, threshold any) AtMostOptValidationBehaviourQuery {
	return &AtMostOptValidation{
		LocalisableError: i18n.LocalisableError{
			Data: AtMostOptValidationTemplData{
				RelationalOV: RelationalOV{
					Flag:      flag,
					Value:     value,
					Threshold: threshold,
				},
			},
		},
	}
}
