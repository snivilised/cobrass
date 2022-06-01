package adapters

import (
	"fmt"
	"reflect"
)

type GenericParameterSet map[string]any

const prefix = "Native object member '"
const suffix = "', does not have a defined member inside the generic params, possible name case issue?"
const identifierPattern = "[A-Z][A-Za-z_\\d]*"

var MissingNativeParamValueFormat = fmt.Sprintf("%v%v%v", prefix, "%v", suffix)
var MissingNativeParamValuePattern = fmt.Sprintf("%v%v%v", prefix, identifierPattern, suffix)

type errorReportEnum int

const (
	MissingNativeMemberValueEn errorReportEnum = iota
	MismatchNativeMemberValueTypeEn
)

type errorReportFieldEnum int

const (
	FormatEn errorReportFieldEnum = iota
	PatternEn
)

type errorReportInfo struct {
	Info []string
}

type errorReportInfoMap map[errorReportEnum]errorReportInfo

var errorTypes errorReportInfoMap = errorReportInfoMap{
	MissingNativeMemberValueEn: errorReportInfo{
		Info: func() []string {
			const prefix = "Native object member '"
			const suffix = "', does not have a defined member inside the generic params, possible name case issue?"

			return []string{
				fmt.Sprintf("%v%v%v", prefix, "%v", suffix),
				fmt.Sprintf("%v%v%v", prefix, identifierPattern, suffix),
			}
		}(),
	},
	MismatchNativeMemberValueTypeEn: errorReportInfo{
		Info: func() []string {
			const prefix = "mismatching types for '"
			const suffix = "')"

			return []string{
				fmt.Sprintf("%v%v' (native: '%v', generic: '%v%v", prefix, "%v", "%v", "%v", suffix),
				fmt.Sprintf("%v.+native.+generic", prefix),
			}
		}(),
	},
}

// ParameterSetCreateOptions options for function CreateParameterSetWith
// Allows the client to control the behaviour of CreateParameterSetWith.
//
// Available options:
//
// "Strict": each member of the native parameter set MUST have an entry in
// the generic params.
//
type ParameterSetCreateOptions struct {
	Strict bool
}

// CreateParameterSet creates native parameter set object from generic
// params with the default options. See CreateParameterSetWith for more
// details.
//
// Default options:
//
// "Strict": true
//
func CreateParameterSet[T comparable](params GenericParameterSet) *T {

	return CreateParameterSetWith[T](params, ParameterSetCreateOptions{
		Strict: true,
	})
}

// CreateParameterSet creates native parameter set object from generic
// params using the options specified.
//
// CreateParameterSet is a generic function on comparable type T which when
// given a map containing field name to values, creates the native
// object required by the client. It is intended that the user should use
// Cobra's Visit facility on the command to create the values associated
// with each flag/argument. The comparable type T is the native parameter
// set object that can be futher used by the cli without tight coupling to
// Cobra.
//
// The parameter set be composed 1 of 2 ways. For each cobra command, there
// may be a different set of parameters that are active. This means there is
// a 1 to many relationship between a cobra command and its associated
// parameter sets. If the client uses 'Visit' method on Cobra's 'FlagSet'
// then on flags that the user has specified will be visited. That means
// there will not be any value set for that corresponding member on the
// native object T.
//
// So the 2 approaches previusly aluded to are as follows:
//
// 1) Optional: Define a single parameter set T which will be used for all combinations
// of valid setrs of flags that the user provides on the command line. When
// using this model, it stands to reason that some members of the native
// object T will not have a value for a flag provided on the command line.
// In this scenario, the client should set ParameterSetCreateOptions.Strict = false,
// which results in no panic occurring when the native member variable does
// not have a corresponding entry in the generic params.
//
// 2) Strict: Define multiple different native pamater sets that represent the different
// ways a particular command can be invoked with different sets of active parameters.
// This is a trickier option to use because the client will need to determine th correct
// logic to detect a particular parameter set and create to create accordingly. In
// this scenario, CreateParameterSetWith should be invoked with Strict set to true
// to ensure that internal logic errors don;'t silently arise due to incorrect
// collation of generic params for the active native parameter set.
//
// Further points of note:
//
// - the native member means the exported (capitalised) member variable of T
//
// - the generic member is the corresponding entry that is linked to the native member
//
// - the name of the native member must match exactly (including by case) the
// corresponding entry in the params map (panic will occur otherwise, if Strict)
//
// - there may be more generic entries in params than there are members in the native type T
//
// Panics in the following circumstances:
//
// - instantiated with anything other than a struct
//
// - params does not contain a corresponding value for native member, when Strict
//
// - the type of the value in params does not match the type of the correspnond
// native member
//
func CreateParameterSetWith[T comparable](params GenericParameterSet, options ParameterSetCreateOptions) *T {
	target := new(T)

	refElemStruct := reflect.ValueOf(target).Elem()
	refTypeOfStruct := refElemStruct.Type()

	if reflect.TypeOf(*target).Kind() == reflect.Struct {
		for i, n := 0, refElemStruct.NumField(); i < n; i++ {
			name := refTypeOfStruct.Field(i).Name

			if value, found := params[name]; found {
				nativeType := refElemStruct.Field(i).Type()
				paramType := reflect.TypeOf(value)

				if nativeType != paramType {
					message := fmt.Sprintf(errorTypes[MismatchNativeMemberValueTypeEn].Info[FormatEn],
						name, nativeType, paramType)

					panic(message)
				}

				refElemStruct.Field(i).Set(reflect.ValueOf(value))
			} else {
				if options.Strict {
					message := fmt.Sprintf(errorTypes[MissingNativeMemberValueEn].Info[FormatEn], name)
					panic(message)
				}
			}
		}
	} else {
		name := reflect.TypeOf(*target).Name()
		panic(fmt.Sprintf("the native parameter set object ('%v') must be a structure", name))
	}

	return target
}

// NewParameterSet is a generic function on comparable type T which when
// given a map containing field name to values, creates the native
// object required by the client. It is intended that the user should use
// Cobra's VisitAll facility on the command to create the values associated
// with each flag/argument. The comparable type T is the native parameter
// set object that can be futher used by the cli without tight coupling to
// Cobra.
// Note:
// - the native member means the exported (capitalised) member variable of T
// - the generic member is the corresponding entry that is linked to the native member
// - the name of the native member must match exactly (including by case) the
// corresponding entry in the params map (panic will occur otherwise)
// - there may be more generic entries in params than there are members in the native type T
// - each member of T MUST have an entry in params
// Panics in the following circumstances:
// - instantiated with anything other than a struct
// - params does not contain a corresponding value for native member
// - the type of the value in params does not match the type of the correspnond
// native member
//
func NewParameterSet[T comparable](params GenericParameterSet) *T {
	target := new(T)

	refElemStruct := reflect.ValueOf(target).Elem()
	refTypeOfStruct := refElemStruct.Type()

	if reflect.TypeOf(*target).Kind() == reflect.Struct {
		for i, n := 0, refElemStruct.NumField(); i < n; i++ {
			name := refTypeOfStruct.Field(i).Name

			if value, found := params[name]; found {
				nativeType := refElemStruct.Field(i).Type()
				paramType := reflect.TypeOf(value)

				if nativeType != paramType {
					message := fmt.Sprintf(errorTypes[MismatchNativeMemberValueTypeEn].Info[FormatEn],
						name, nativeType, paramType)

					panic(message)
				}

				refElemStruct.Field(i).Set(reflect.ValueOf(value))
			} else {
				message := fmt.Sprintf(errorTypes[MissingNativeMemberValueEn].Info[FormatEn], name)
				panic(message)
			}

		}
	} else {
		name := reflect.TypeOf(*target).Name()
		panic(fmt.Sprintf("the native parameter set object ('%v') must be a structure", name))
	}

	return target
}
