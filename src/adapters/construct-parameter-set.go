package adapters

import (
	"fmt"
	"reflect"
)

func constructParameterSet[T comparable](
	target *T,
	params GenericParameterSet,
	options ParameterSetCreateOptions) *T {

	refElemStruct := reflect.ValueOf(target).Elem()
	refTypeOfStruct := refElemStruct.Type()

	if reflect.TypeOf(*target).Kind() == reflect.Struct {
		for i, n := 0, refElemStruct.NumField(); i < n; i++ {
			name := refTypeOfStruct.Field(i).Name

			if value, found := params[name]; found {
				nativeType := refElemStruct.Field(i).Type()
				paramType := reflect.TypeOf(value)

				if nativeType != paramType {
					message := fmt.Sprintf(errorTypes[mismatchNativeMemberValueTypeEn].Info[FormatEn],
						name, nativeType, paramType)

					panic(message)
				}

				refElemStruct.Field(i).Set(reflect.ValueOf(value))
			} else {
				if options.Strict {
					message := fmt.Sprintf(errorTypes[missingNativeMemberValueEn].Info[FormatEn], name)
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
