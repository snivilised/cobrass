package adapters

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

type EnumInfo[E ~int] struct {
	// The address of Source should be used with BindEnum. The reason why we
	// need an alternative for the 'to' parameter on the binder method is that
	// the associated native member is going to be the pseudo enum type, which
	// is not compatible with the string value that the user provides on the
	// comand line. So Source is just a temporary place holder for the value,
	// which subsequently needs to be converted and injected into the native
	// parameter set(see Value() method)
	//
	Source string

	acceptables   AcceptableEnumValues[E]
	reverseLookup lookupEnumValue[E]
}

// En, returns the underlying int based enum associated with the provided
// string value as defined by the Acceptables.
//
func (info *EnumInfo[E]) En(value string) E {
	return E(info.reverseLookup[value])
}

// Value, returns the value of the enum that is stored within the EnumInfo
// captured from the command line.
//
func (info *EnumInfo[E]) Value() E {
	return E(info.reverseLookup[info.Source])
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
// EnumInfo[OutputFormatEnum]
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
				panic(fmt.Sprintf("'%v' already exists, invalid enum info specified", existing))
			}

			info.reverseLookup[acc] = enum
		}
	}

	return info
}

// FlagInfo collates together the paramters passed into the bind methods
// The Bind methods are just a wrapper around invoking the type based methods
// on the cobra flag set in order to define flags. They only cover the versions
// of those methods that allow the definition of a short code. If the client
// needs to use those versions without the short, they can do so directly.
//
type FlagInfo struct {
	Name    string
	Usage   string
	Short   string
	Default any
}

// NewFlagInfo factory function for FlagInfo
//
func NewFlagInfo(usage string, short string, def any) *FlagInfo {
	n := usage
	i := strings.Index(n, " ")
	if i >= 0 {
		n = n[:i]
	}
	return &FlagInfo{
		Name:    n,
		Usage:   usage,
		Short:   short,
		Default: def,
	}
}

// ParamSet represents a set of flags/options/positional args for a cmmand.
// The term 'parameter set' is used really to distinguish from other established
// abstractions (flags/options/positional args, otherwise to be referred to as
// inputs). The ParamSet is used to ensure that all these inputs are collated
// into a single entity that the application can refer to as required. A command
// can have multiple parameter sets associated with it, but will probably best
// be used with a single parameter set, where inputs not provided by the end user
// are defaulted, perhaps from config. If its essential to distinguish between
// different activation scenarios (ie which set of parameters that the user provides)
// then the client can define mutiple parameter sets to reflect this.
//
// The binder mnethods are defined explicitly for each type as 'go' does not
// allow for generic parameters defined at the method level as opposed to
// being defined on the receiver struct.
//
// The generic parameter N represents the client defined native parameter set. Eg:
//
// type WidgetParameterSet struct {
// 	 Directory string
// 	 Output    string
// 	 Format    OutputFormatEnum
// 	 Shape     InfexionShapeEnum
// 	 Concise   bool
// 	 Strategy  TravseralStratgeyEnum
// 	 Overwrite bool
// 	 Pattern   string}
//
// ... is known as the 'native' parameter set for a 'widget' command which
// would be used to instantiate ParamSet in a declaration as follows:
//
// var paramSet *ParamSet[WidgetParameterSet]
//
type ParamSet[N any] struct {
	// Native is the native client defined parameter set instance, which
	// must be a struct.
	//
	Native  *N
	flagSet *pflag.FlagSet
}

// NewParamSet is the factory function, which creates a 'parameter set' for
// a command. Each command can have multiple command sets, reflecting the
// different ways a command can be used
//
// paramSet = NewParamSet[WidgetParameterSet](widgetCommand)
//
// The generic parameter N represents the client defined native parameter set.
//
func NewParamSet[N any](command *cobra.Command) (ps *ParamSet[N]) {
	ps = new(ParamSet[N])
	ps.flagSet = command.Flags()
	ps.Native = new(N)

	if reflect.TypeOf(*ps.Native).Kind() != reflect.Struct {
		name := reflect.TypeOf(*ps.Native).Name()
		panic(fmt.Sprintf("the native param set object ('%v') must be a structure", name))
	}
	return ps
}

// BindString, binds string flag with a shorthand
//
func (params *ParamSet[N]) BindString(info *FlagInfo, to *string) *ParamSet[N] {
	params.flagSet.StringVarP(to, info.Name, info.Short, info.Default.(string), info.Usage)

	return params
}

// BindInt, binds int flag with a shorthand
//
func (params *ParamSet[N]) BindInt(info *FlagInfo, to *int) *ParamSet[N] {
	params.flagSet.IntVarP(to, info.Name, info.Short, info.Default.(int), info.Usage)

	return params
}

// BindInt8, binds int8 flag with a shorthand
//
func (params *ParamSet[N]) BindInt8(info *FlagInfo, to *int8) *ParamSet[N] {
	params.flagSet.Int8VarP(to, info.Name, info.Short, info.Default.(int8), info.Usage)

	return params
}

// BindInt16, binds int16 flag with a shorthand
//
func (params *ParamSet[N]) BindInt16(info *FlagInfo, to *int16) *ParamSet[N] {
	params.flagSet.Int16VarP(to, info.Name, info.Short, info.Default.(int16), info.Usage)

	return params
}

// BindInt32, binds int32 flag with a shorthand
//
func (params *ParamSet[N]) BindInt32(info *FlagInfo, to *int32) *ParamSet[N] {
	params.flagSet.Int32VarP(to, info.Name, info.Short, info.Default.(int32), info.Usage)

	return params
}

// BindInt64, binds int64 flag with a shorthand
//
func (params *ParamSet[N]) BindInt64(info *FlagInfo, to *int64) *ParamSet[N] {
	params.flagSet.Int64VarP(to, info.Name, info.Short, info.Default.(int64), info.Usage)

	return params
}

// BindUint, binds uint flag with a shorthand
//
func (params *ParamSet[N]) BindUint(info *FlagInfo, to *uint) *ParamSet[N] {
	params.flagSet.UintVarP(to, info.Name, info.Short, info.Default.(uint), info.Usage)

	return params
}

// BindUint8, binds int8 flag with a shorthand
//
func (params *ParamSet[N]) BindUint8(info *FlagInfo, to *uint8) *ParamSet[N] {
	params.flagSet.Uint8VarP(to, info.Name, info.Short, info.Default.(uint8), info.Usage)

	return params
}

// BindUint16, binds int16 flag with a shorthand
//
func (params *ParamSet[N]) BindUint16(info *FlagInfo, to *uint16) *ParamSet[N] {
	params.flagSet.Uint16VarP(to, info.Name, info.Short, info.Default.(uint16), info.Usage)

	return params
}

// BindUint32, binds int32 flag with a shorthand
//
func (params *ParamSet[N]) BindUint32(info *FlagInfo, to *uint32) *ParamSet[N] {
	params.flagSet.Uint32VarP(to, info.Name, info.Short, info.Default.(uint32), info.Usage)

	return params
}

// BindUint64, binds int64 flag with a shorthand
//
func (params *ParamSet[N]) BindUint64(info *FlagInfo, to *uint64) *ParamSet[N] {
	params.flagSet.Uint64VarP(to, info.Name, info.Short, info.Default.(uint64), info.Usage)

	return params
}

// BindBool, binds bool flag with a shorthand. Note, the default value for
// a bool based flag is ignored as it is 'sensibly' defaulted to false. When
// the user does not specify the flag, the value will be its default (false).
// When present, it is set to true as would be expected.
//
func (params *ParamSet[N]) BindBool(info *FlagInfo, to *bool) *ParamSet[N] {
	params.flagSet.BoolVarP(to, info.Name, info.Short, false, info.Usage)

	return params
}

// BindFloat64, binds float64 flag with a shorthand
//
func (params *ParamSet[N]) BindFloat64(info *FlagInfo, to *float64) *ParamSet[N] {
	params.flagSet.Float64VarP(to, info.Name, info.Short, info.Default.(float64), info.Usage)

	return params
}

// BindFloat32, binds float32 flag with a shorthand
//
func (params *ParamSet[N]) BindFloat32(info *FlagInfo, to *float32) *ParamSet[N] {
	params.flagSet.Float32VarP(to, info.Name, info.Short, info.Default.(float32), info.Usage)

	return params
}

// BindEnum, binds pseudo int based enum flag with a shorthand. Note that normally
// the client would bind to a member of the native parameter set. However, since
// there is a discrepency between the type of the native int based pseudo enum
// member and the equivalent acceptable string value typed by the user on the
// command line (idiomatically stored on the enum info), the client needs to extract
// the enum value from the enum info, something like this:
//
// paramSet.Native.Format = OutputFormatEnumInfo.Value()
//
// The best place to put this would be inside the PreRun/PreRunE function, assuming the
// paramset and the enum info are both in scope. Actually, every int based enum
// flag, would need to have this assignment performed.
//
func (params *ParamSet[N]) BindEnum(info *FlagInfo, to *string) {
	params.flagSet.StringVarP(to, info.Name, info.Short, info.Default.(string), info.Usage)
}

// BindDuration, binds time.Duration flag with a shorthand
//
func (params *ParamSet[N]) BindDuration(info *FlagInfo, to *time.Duration) *ParamSet[N] {
	params.flagSet.DurationVarP(to, info.Name, info.Short, info.Default.(time.Duration), info.Usage)

	return params
}

// BindIp, binds time.Duration flag with a shorthand
//
func (params *ParamSet[N]) BindIp(info *FlagInfo, to *net.IP) *ParamSet[N] {
	params.flagSet.IPVarP(to, info.Name, info.Short, info.Default.(net.IP), info.Usage)

	return params
}

// BindIpMask, binds net.IPMask flag with a shorthand
//
func (params *ParamSet[N]) BindIpMask(info *FlagInfo, to *net.IPMask) *ParamSet[N] {
	params.flagSet.IPMaskVarP(to, info.Name, info.Short, info.Default.(net.IPMask), info.Usage)

	return params
}

// BindStringSlice, binds string slice flag with a shorthand
//
func (params *ParamSet[N]) BindStringSlice(info *FlagInfo, to *[]string) {
	params.flagSet.StringSliceVarP(to, info.Name, info.Short, info.Default.([]string), info.Usage)
}

// BindIntSlice, binds int slice flag with a shorthand
//
func (params *ParamSet[N]) BindIntSlice(info *FlagInfo, to *[]int) {
	params.flagSet.IntSliceVarP(to, info.Name, info.Short, info.Default.([]int), info.Usage)
}

// BindUintSlice, binds uint slice flag with a shorthand
//
func (params *ParamSet[N]) BindUintSlice(info *FlagInfo, to *[]uint) {
	params.flagSet.UintSliceVarP(to, info.Name, info.Short, info.Default.([]uint), info.Usage)
}

// BindBoolSlice, binds uint slice flag with a shorthand
//
func (params *ParamSet[N]) BindBoolSlice(info *FlagInfo, to *[]bool) {
	params.flagSet.BoolSliceVarP(to, info.Name, info.Short, info.Default.([]bool), info.Usage)
}

// BindFloat64Slice, binds float64 slice flag with a shorthand
//
func (params *ParamSet[N]) BindFloat64Slice(info *FlagInfo, to *[]float64) {
	params.flagSet.Float64SliceVarP(to, info.Name, info.Short, info.Default.([]float64), info.Usage)
}

// BindFloat32Slice, binds float32 slice flag with a shorthand
//
func (params *ParamSet[N]) BindFloat32Slice(info *FlagInfo, to *[]float32) {
	params.flagSet.Float32SliceVarP(to, info.Name, info.Short, info.Default.([]float32), info.Usage)
}

// BindIpSlice, binds float32 slice flag with a shorthand
//
func (params *ParamSet[N]) BindIpSlice(info *FlagInfo, to *[]net.IP) {
	params.flagSet.IPSliceVarP(to, info.Name, info.Short, info.Default.([]net.IP), info.Usage)
}
