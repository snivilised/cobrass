package adapters

import (
	"fmt"
)

/*
GenericParamSet is a map of flag names to their value as represented in Cobra. Note
flag values are reprented by strings, which are internally converted to the required
type on the native parameter set, via reflection
*/
type GenericParamSet map[string]any

const identifierPattern = "[A-Z][A-Za-z_\\d]*"

type ErrorReportEnum int

const (
	MissingNativeMemberValueEn ErrorReportEnum = iota
	MismatchNativeMemberValueTypeEn
)

type ErrorReportFieldEnum int

const (
	FormatEn ErrorReportFieldEnum = iota
	PatternEn
)

type errorReportInfo struct {
	Info []string
}

type errorReportInfoMap map[ErrorReportEnum]errorReportInfo

var ErrorTypes errorReportInfoMap = errorReportInfoMap{
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

// VisitStrategyEnum defines the valid values that can be used on the Extract method of
// ParamSetFactory.
//
type VisitStrategyEnum int

const (
	// ActiveFlagsEn extract only the active flags for the command
	//
	ActiveFlagsEn VisitStrategyEnum = iota

	// AllFlagsEn extract all defined flags, whether they are active
	// or not
	//
	AllFlagsEn
)

// CreatePsOptions options for function CreateParameterSetWith
// Allows the client to control the behaviour of CreateParameterSetWith.
//
type CreatePsOptions struct {
	// Strict defines whether it is permissable for native parameter set
	// member variable has a correspond entry inside the generic param
	// map. When true, each member of the native parameter set MUST have
	// an entry in the generic params.
	//
	Strict bool
}

// ExtractPsOptions specifies parameter set extraction options. Also see
// ParameterSetCreateOptions
//
type ExtractPsOptions struct {
	// Create species parameterset creation options
	//
	CreatePsOptions

	// Strategy determines which Visit method on the cobra FlagSet will
	// be invoked
	//
	Strategy VisitStrategyEnum
}
