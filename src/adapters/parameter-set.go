package adapters

import (
	"fmt"
	"reflect"
)

type GenericParameterSet map[string]any

const prefix = "Native object member '"
const suffix = "', does not have a defined member inside the generic params, possible name case issue?"
const IdentifierPattern = "[A-Z][A-Za-z_\\d]+"

var MissingNativeParamValueFormat = fmt.Sprintf("%v%v%v", prefix, "%v", suffix)
var MissingNativeParamValuePattern = fmt.Sprintf("%v%v%v", prefix, IdentifierPattern, suffix)

// NewParameterSet is a generic function on comparable type T which when
// given a map containing field name to values, creates the native
// object required by the client. It is intended that the user should use
// Cobra's VisitAll facility on the command to create the values associated
// with each flag/argument. The comparable type T is the native parameter
// set object that can be futher used by the cli without tight coupling to
// Cobra.
// Panics, if instantiated with anything other than a struct.
//
func NewParameterSet[T comparable](params GenericParameterSet) *T {
	target := new(T)

	refElemStruct := reflect.ValueOf(target).Elem()
	refTypeOfStruct := refElemStruct.Type()

	if reflect.TypeOf(*target).Kind() == reflect.Struct {
		for i, n := 0, refElemStruct.NumField(); i < n; i++ {
			name := refTypeOfStruct.Field(i).Name

			if value, found := params[name]; found {
				refElemStruct.Field(i).Set(reflect.ValueOf(value))
			} else {
				panic(fmt.Sprintf(MissingNativeParamValueFormat, name))
			}

		}
	} else {
		name := reflect.TypeOf(*target).Name()
		panic(fmt.Sprintf("the native parameter set object ('%v') must be a structure", name))
	}

	return target
}
