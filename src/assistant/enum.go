//nolint:revive // receiver naming hint doesn't make sense
package assistant

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/snivilised/cobrass/src/assistant/translate"
)

// AcceptableEnumValues maps values of enum type to an array of
// string values that are acceptable to be interpreted as that enum
// value. E is the int based pseudo enum type.
//
// This type allows multiple string values to be taken as equivalent
// values. The rationale behing this, is for ease of use for the end user.
// The code can use long winded expressive enum names, without imposing
// the burden of requiring the user to have to exactly type in that long
// enum name. Eg an enum value of XmlFormatEn is not end user friendly,
// we wouldnt them to have to type that in the command line. "xml", or
// even just "x" would be much easier, but we couldnt really use those
// names as enum values in code, because they are too generic. This
// is where AcceptableEnumValues comes in. It provides a mapping from
// user friendly names to internal enum names, eg, we want to define
// an acceptable values collection for our predefined output format
// enum: 'OutputFormatEnum' as follows:
//
// given:
//
// type OutputFormatEnum int
// const (
//   _ OutputFormatEnum = iota
// 	 XmlFormatEn
// 	 JsonFormatEn
// 	 TextFormatEn
// 	 ScribbleFormatEn)
//
// we can define our acceptables asfollows:
//
// AcceptableEnumValues[OutputFormatEnum]{
// 	 XmlFormatEn:      []string{"xml", "x"},
// 	 JsonFormatEn:     []string{"json", "j"},
// 	 TextFormatEn:     []string{"text", "tx"},
// 	 ScribbleFormatEn: []string{"scribble", "scribbler", "scr"}}
//
// Note, when composing the list of acceptable string values for an enum, it is
// recommended to make the first item to be the most expressive (ie the longest
// string value), because whenever an enum needs to be printed, the client
// can use the 'NameOf' method to display the useful name of the enum rather
// than getting the integer value, which is not very expressive but
// is what you get by default using the %v formatter for print functions.
//
type AcceptableEnumValues[E ~int] map[E][]string

// lookupEnumValue provides a reverse lookup of an enum value
// given a string value.
//
type lookupEnumValue[E ~int] map[string]E

// https://medium.com/trendyol-tech/contributing-the-go-compiler-adding-new-tilde-operator-f66d0c6cff7
//

// EnumInfo represents the meta data for a pseduo int based enum type.
//
type EnumInfo[E ~int] struct {
	// (this should probably go into a different module as it's use goes beyond
	// the context of cobra)
	//
	acceptables   AcceptableEnumValues[E]
	reverseLookup lookupEnumValue[E]
}

// NewEnumInfo is the factory function that creates an EnumInfo instance given
// a client defined acceptable values collection. Builds the reverse lookup
// which then allows the client to lookup the enum value for a string.
// AcceptableEnumValues is defined from enum => slice of acceptables, because
// the literal representation of this is less verbose than defining the
// mapping the other way around, ie from acceptable-value => enum, as this
// would require multiple entries for each enum value. When we defined it
// as acceptable-value => enum, we can group together the acceptables values
// into a string array and thus is more concise.
//
// The generic variable E represents the int based enum type, so given:
//
// type OutputFormatEnum int
// const (
// 	 XmlFormatEn OutputFormatEnum = iota + 1
// 	 JsonFormatEn
// 	 TextFormatEn
// 	 ScribbleFormatEn)
//
// ... the user should define an EnumInfo for it as:
//
// EnumInfo[OutputFormatEnum].
//
func NewEnumInfo[E ~int](acceptables AcceptableEnumValues[E]) *EnumInfo[E] {
	info := new(EnumInfo[E])
	info.acceptables = acceptables
	info.reverseLookup = make(lookupEnumValue[E])

	// build the reverse lookup that will allow the client to lookup the enum
	// for a string value. acceptables is only a map of the enum value to an
	// array of strings that its associated with
	//
	for enum, values := range acceptables {
		for _, acc := range values {
			if existing, found := info.reverseLookup[acc]; found {
				panic(translate.GetEnumValueAlreadyExistsErrorMessage(info.NameOf(existing), int(existing)))
			}

			info.reverseLookup[acc] = enum
		}
	}

	return info
}

// NewValue creates a new EnumValue associated with this EnumInfo.
//
func (info *EnumInfo[E]) NewValue() EnumValue[E] {
	return EnumValue[E]{Info: info}
}

// NewSlice creates a new EnumSlice associated with this EnumInfo.
//
func (info *EnumInfo[E]) NewSlice() EnumSlice[E] {
	return EnumSlice[E]{Info: info, Source: []string{}}
}

// En, returns the underlying int based enum associated with the provided
// string value as defined by the Acceptables.
//
func (info *EnumInfo[E]) En(value string) E {
	return info.reverseLookup[value]
}

// IsValid returns true if the string is an acceptable value for this enum
// false otherwise.
//
func (info *EnumInfo[E]) IsValid(value string) bool {
	_, found := info.reverseLookup[value]
	return found
}

// IsValidOrEmpty returns true if the string is an acceptable value for this enum
// or the empty string false otherwise.
//
func (info *EnumInfo[E]) IsValidOrEmpty(value string) bool {
	if value == "" {
		return true
	}

	_, found := info.reverseLookup[value]

	return found
}

// String returns a string representing contents of all acceptable values for
// the enum.
//
func (info *EnumInfo[E]) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "=== (TYPE: %T) ===\n", info)

	for k, v := range info.reverseLookup {
		fmt.Fprintf(&builder, "--- [%v] => [%v] (%v) ---\n", k, info.NameOf(v), v)
	}

	return builder.String()
}

// NameOf returns the first acceptable name for the enum value specified.
// Ideally, there would be a way in go reflection to obtain the name of a
// variable (as opposed to type name), but this isnt possible. Go reflection
// currently can only query type names not variable or function names, so
// the NameOf method is used as a workaround.
//
func (info *EnumInfo[E]) NameOf(enum E) string {
	return info.acceptables[enum][0]
}

type EnumValue[E ~int] struct {
	// Info is the EnumInfo associated with this enam value
	//
	Info *EnumInfo[E]

	// The address of Source should be used with BindEnum. The reason why we
	// need an alternative for the 'to' parameter on the binder method is that
	// the associated native member is going to be the pseudo enum type, which
	// is not compatible with the string value that the user provides on the
	// comand line. So Source is just a temporary place holder for the value,
	// which subsequently needs to be converted and injected into the native
	// parameter set(see Value() method)
	//
	Source string
}

// Value, returns the value of the enum that is stored within the EnumValue
// captured from the command line.
//
func (ev *EnumValue[E]) Value() E {
	return ev.Info.reverseLookup[ev.Source]
}

// IsValid returns true if the string is an acceptable value for this enum
// false otherwise.
//
func (ev *EnumValue[E]) IsValid() bool {
	return ev.Info.IsValid(ev.Source)
}

// IsValidOrEmpty returns true if the string is an acceptable value for this enum
// or the empty string false otherwise.
//
func (ev *EnumValue[E]) IsValidOrEmpty() bool {
	return ev.Info.IsValidOrEmpty(ev.Source)
}

// String returns the content of Source, assuming it is a valid acceptable enum
// value. If not valid or not set yet causes panic. As it curently stands, the
// client needs to validate incoming input as performed in a binder operation.
//
func (ev *EnumValue[E]) String() string {
	if _, found := ev.Info.reverseLookup[ev.Source]; !found {
		panic(fmt.Errorf(translate.GetIsNotValidEnumValueErrorMessage(ev.Source)))
	} else {
		return ev.Source
	}
}

// EnumSlice represents a collection of EnumValues. Note that this abstraction is
// not the same as defining a slice of EnumValues, ie []EnumValues.
type EnumSlice[E ~int] struct {
	// Info is the EnumInfo associated with this enam value
	//
	Info *EnumInfo[E]

	// The address of Source should be used with BindEnum. The reason why we
	// need an alternative for the 'to' parameter on the binder method is that
	// the associated native member is going to be the pseudo enum type, which
	// is not compatible with the string value that the user provides on the
	// comand line. So Source is just a temporary place holder for the value,
	// which subsequently needs to be converted and injected into the native
	// parameter set(see Value() method)
	//
	Source []string
}

// Values, returns an an array of int based enum values replicating the slice of
// string values stored in Source.
//
func (es *EnumSlice[E]) Values() []E {
	return lo.Map(es.Source, func(v string, _ int) E {
		return es.Info.reverseLookup[v]
	})
}

// AllAreValid returns true if the Source strings are all acceptable values for
// this enum false otherwise.
//
func (es *EnumSlice[E]) AllAreValid() bool {
	return lo.EveryBy(es.Source, func(v string) bool {
		return es.Info.IsValid(v)
	})
}

// AllAreValidOrEmpty returns true if the Source strings are all acceptable values
// for this enum or the empty string false otherwise.
//
func (es *EnumSlice[E]) AllAreValidOrEmpty() bool {
	return lo.EveryBy(es.Source, func(v string) bool {
		return es.Info.IsValidOrEmpty(v)
	})
}
