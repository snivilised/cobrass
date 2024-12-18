package assistant

import (
	"fmt"
	"sort"
	"strings"

	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/cobrass/src/internal/third/lo"
)

// AcceptableEnumValues maps values of enum type to an array of
// string values that are acceptable to be interpreted as that enum
// value. E is the int based pseudo enum type.
//
// This type allows multiple string values to be taken as equivalent
// values. The rationale behind this, is for ease of use for the end user.
// The code can use long winded expressive enum names, without imposing
// the burden of requiring the user to have to exactly type in that long
// enum name. Eg an enum value of XmlFormatEn is not end user friendly,
// we wouldn't them to have to type that in the command line. "xml", or
// even just "x" would be much easier, but we couldn't really use those
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
//
//	  _ OutputFormatEnum = iota
//		 XmlFormatEn
//		 JsonFormatEn
//		 TextFormatEn
//		 ScribbleFormatEn)
//
// we can define our acceptables as follows:
//
//	AcceptableEnumValues[OutputFormatEnum]{
//		 XmlFormatEn:      []string{"xml", "x"},
//		 JsonFormatEn:     []string{"json", "j"},
//		 TextFormatEn:     []string{"text", "tx"},
//		 ScribbleFormatEn: []string{"scribble", "scribbler", "scr"}}
//
// Note, when composing the list of acceptable string values for an enum, it is
// recommended to make the first item to be the most expressive (ie the longest
// string value), because whenever an enum needs to be printed, the client
// can use the 'NameOf' method to display the useful name of the enum rather
// than getting the integer value, which is not very expressive but
// is what you get by default using the %v formatter for print functions.
type AcceptableEnumValues[E ~int] map[E][]string

// lookupEnumValue provides a reverse lookup of an enum value
// given a string value.
type lookupEnumValue[E ~int] map[string]E

// https://medium.com/trendyol-tech/contributing-the-go-compiler-adding-new-tilde-operator-f66d0c6cff7
//

// EnumInfo represents the meta data for a pseudo int based enum type.
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
//
//	XmlFormatEn OutputFormatEnum = iota + 1
//	JsonFormatEn
//	TextFormatEn
//	ScribbleFormatEn)
//
// ... the user should define an EnumInfo for it as:
//
// EnumInfo[OutputFormatEnum].
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
				panic(locale.NewEnumValueValueAlreadyExistsNativeError(
					info.NameOf(existing), int(existing)),
				)
			}

			info.reverseLookup[strings.ToLower(acc)] = enum
		}
	}

	return info
}

// NewValue creates a new EnumValue associated with this EnumInfo.
func (info *EnumInfo[E]) NewValue() EnumValue[E] {
	return EnumValue[E]{Info: info}
}

// NewValue creates new enum value initialised with the provided string value
func (info *EnumInfo[E]) NewWith(value string) EnumValue[E] {
	enum := EnumValue[E]{Info: info}
	enum.Source = value

	return enum
}

// NewSlice creates a new EnumSlice associated with this EnumInfo.
func (info *EnumInfo[E]) NewSlice() EnumSlice[E] {
	return EnumSlice[E]{Info: info, Source: []string{}}
}

// En, returns the underlying int based enum associated with the provided
// string value as defined by the Acceptables.
func (info *EnumInfo[E]) En(value string) E {
	return info.reverseLookup[strings.ToLower(value)]
}

// IsValid returns true if the string is an acceptable value for this enum
// false otherwise.
func (info *EnumInfo[E]) IsValid(value string) bool {
	_, found := info.reverseLookup[strings.ToLower(value)]
	return found
}

// IsValidOrEmpty returns true if the string is an acceptable value for this enum
// or the empty string false otherwise.
func (info *EnumInfo[E]) IsValidOrEmpty(value string) bool {
	if value == "" {
		return true
	}

	_, found := info.reverseLookup[strings.ToLower(value)]

	return found
}

// String returns a string representing contents of all acceptable values for
// the enum.
func (info *EnumInfo[E]) String() string {
	keys := lo.Keys(info.reverseLookup)

	return lo.Reduce(keys, func(agg string, current string, _ int) string {
		return fmt.Sprintf("%v%v(%v), ", agg, current, info.En(current))
	}, "")
}

// Acceptable returns a string that indicates the set of acceptable values
// for this enum. Since the client can create multiple values for a single
// enum value, the client can choose whether they want all possible values
// or just the primary value for each enum entry. The client should use
// this method if all values for all enumerations in this enum info should
// be represented in the returned string. This method (and AcceptablePrimes)
// are typically used to prompt the end user of the acceptable values for an
// enum based option in a cli application.
func (info *EnumInfo[E]) Acceptable() string {
	keys := lo.Keys(info.reverseLookup)
	sort.Strings(keys)

	return slashes + lo.Reduce(keys, func(agg, current string, _ int) string {
		return fmt.Sprintf("%v%v/", agg, current)
	}, "") + "/"
}

// AcceptablePrimes returns a string that indicates the set of acceptable values
// for this enum. Similar to Acceptable except that the returned string only
// indicates the primary entry for each enumeration in the enum info.
func (info *EnumInfo[E]) AcceptablePrimes() string {
	l := len(info.acceptables)
	elements := make([]string, l)
	keys := lo.Keys(info.acceptables)

	// type func(a E, b E) bool of func(a E, b E) bool {…}
	// does not match inferred type:
	// func(a E, b E) int for func(a E, b E) int
	//
	// slices.SortFunc(keys, func(a E, b E) bool {
	// 	return a > b
	// })
	//
	// This error does manifest itself locally, rather, it breaks
	// downstream projects eg arcadia. Without sorting, the order
	// of the acceptable primes can not be assured.
	//

	for i, v := range keys {
		// can't reduce collection[E], because E is not accumulate-able
		//
		elements[i] = info.acceptables[v][0]
	}

	return slashes + strings.Join(elements, "/") + slashes
}

// NameOf returns the first acceptable name for the enum value specified.
// Ideally, there would be a way in go reflection to obtain the name of a
// variable (as opposed to type name), but this isn't possible. Go reflection
// currently can only query type names not variable or function names, so
// the NameOf method is used as a workaround.
func (info *EnumInfo[E]) NameOf(enum E) string {
	return info.acceptables[enum][0]
}

type EnumValue[E ~int] struct {
	// Info is the EnumInfo associated with this enum value
	//
	Info *EnumInfo[E]

	// The address of Source should be used with BindEnum. The reason why we
	// need an alternative for the 'to' parameter on the binder method is that
	// the associated native member is going to be the pseudo enum type, which
	// is not compatible with the string value that the user provides on the
	// command line. So Source is just a temporary place holder for the value,
	// which subsequently needs to be converted and injected into the native
	// parameter set(see Value() method)
	//
	Source string
}

// Value, returns the value of the enum that is stored within the EnumValue
// captured from the command line.
func (ev *EnumValue[E]) Value() E {
	return ev.Info.reverseLookup[ev.Source]
}

// IsValid returns true if the string is an acceptable value for this enum
// false otherwise.
func (ev *EnumValue[E]) IsValid() bool {
	return ev.Info.IsValid(ev.Source)
}

// IsValidOrEmpty returns true if the string is an acceptable value for this enum
// or the empty string false otherwise.
func (ev *EnumValue[E]) IsValidOrEmpty() bool {
	return ev.Info.IsValidOrEmpty(ev.Source)
}

// String returns the content of Source, assuming it is a valid acceptable enum
// value. If not valid or not set yet causes panic. As it currently stands, the
// client needs to validate incoming input as performed in a binder operation.
func (ev *EnumValue[E]) String() string {
	if _, found := ev.Info.reverseLookup[ev.Source]; !found {
		panic(locale.NewIsNotValidEnumValueNativeError(ev.Source))
	}

	return ev.Source
}

// EnumSlice represents a collection of EnumValues. Note that this abstraction is
// not the same as defining a slice of EnumValues, ie []EnumValues.
type EnumSlice[E ~int] struct {
	// Info is the EnumInfo associated with this enum value
	//
	Info *EnumInfo[E]

	// The address of Source should be used with BindEnum. The reason why we
	// need an alternative for the 'to' parameter on the binder method is that
	// the associated native member is going to be the pseudo enum type, which
	// is not compatible with the string value that the user provides on the
	// command line. So Source is just a temporary place holder for the value,
	// which subsequently needs to be converted and injected into the native
	// parameter set(see Value() method)
	//
	Source []string
}

// Values, returns an an array of int based enum values replicating the slice of
// string values stored in Source.
func (es *EnumSlice[E]) Values() []E {
	return lo.Map(es.Source, func(v string, _ int) E {
		return es.Info.reverseLookup[v]
	})
}

// AllAreValid returns true if the Source strings are all acceptable values for
// this enum false otherwise.
func (es *EnumSlice[E]) AllAreValid() bool {
	return lo.EveryBy(es.Source, func(v string) bool {
		return es.Info.IsValid(v)
	})
}

// AllAreValidOrEmpty returns true if the Source strings are all acceptable values
// for this enum or the empty string false otherwise.
func (es *EnumSlice[E]) AllAreValidOrEmpty() bool {
	return lo.EveryBy(es.Source, func(v string) bool {
		return es.Info.IsValidOrEmpty(v)
	})
}

const (
	slashes = "//"
)
