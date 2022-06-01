package adapters

import (
	"fmt"
)

/*
GenericParameterSet is a map of flag names to their value as represented in Cobra. Note
flag values are reprented by strings, which are internally converted to the required
type on the native parameter set, via reflection
*/
type GenericParameterSet map[string]any

const identifierPattern = "[A-Z][A-Za-z_\\d]*"

type errorReportEnum int

const (
	missingNativeMemberValueEn errorReportEnum = iota
	mismatchNativeMemberValueTypeEn
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
	missingNativeMemberValueEn: errorReportInfo{
		Info: func() []string {
			const prefix = "Native object member '"
			const suffix = "', does not have a defined member inside the generic params, possible name case issue?"

			return []string{
				fmt.Sprintf("%v%v%v", prefix, "%v", suffix),
				fmt.Sprintf("%v%v%v", prefix, identifierPattern, suffix),
			}
		}(),
	},
	mismatchNativeMemberValueTypeEn: errorReportInfo{
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
type ParameterSetCreateOptions struct {
	// Strict defines whether it is permissable for native parameter set
	// member variable has a correspond entry inside the generic param
	// map. When true, each member of the native parameter set MUST have
	// an entry in the generic params.
	//
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
	constructParameterSet(target, params, options)

	return target
}
