package helpers

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func CallE(fn any, args []any) error {
	reflFunc := reflect.ValueOf(fn)
	full := runtime.FuncForPC(reflFunc.Pointer()).Name()
	segments := strings.Split(full, "/")
	name := segments[len(segments)-1]

	// create a slice of reflect values for the arguments
	const initialSize = 0

	var reflectArgs = make([]reflect.Value, initialSize, len(args))

	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	result := reflFunc.Call(reflectArgs)

	err, ok := (result[0].Interface()).(error)
	if !ok {
		panic(
			fmt.Errorf(
				"error function invocation for '%v' (args: '%s') did not return an error",
				name, args,
			),
		)
	}

	return err
}
