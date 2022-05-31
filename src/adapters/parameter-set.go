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
