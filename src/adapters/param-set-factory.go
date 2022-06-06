package adapters

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ParamSetFactory[T comparable] struct {
	Params GenericParamSet
}

/*
Extract is a helper method that the client can use to extract a native parameter
set object, denoted by type T. This is an alternative to using the CreateParameterSet
directly. The client needs to simply pass in the FlagSet instance on the cobra command,
ie 'command.Flags()'. The result is the native parameter set object.

Please see CreateParameterSet/CreateParameterSetWith functions for more details
*/
func (factory *ParamSetFactory[T]) Extract(flagSet pflag.FlagSet, options ExtractPsOptions) *T { // New

	target := new(T)
	factory.Params = make(GenericParamSet)

	handler := func(f *pflag.Flag) {
		// The name has to be capitalised, because the native structure will have public
		// properties, which must begin with a capital letter
		//
		// SHOULD WE USE SAFE INSERT HERE
		//
		name := cases.Title(language.English).String(f.Name)
		factory.Params[name] = f.Value

		fmt.Printf("===> ⚠️ FLAG - name: '%v', value: '%v'\n", name, f.Value)
	}

	switch options.Strategy {
	case ActiveFlagsEn:
		flagSet.Visit(handler)
	case AllFlagsEn:
		flagSet.VisitAll(handler)
	}

	return factory.constructParameterSet(target, options.CreatePsOptions)
}

func (factory *ParamSetFactory[T]) constructParameterSet(
	target *T,
	options CreatePsOptions) *T {

	refElemStruct := reflect.ValueOf(target).Elem()
	refTypeOfStruct := refElemStruct.Type()

	if reflect.TypeOf(*target).Kind() == reflect.Struct {
		for i, n := 0, refElemStruct.NumField(); i < n; i++ {
			name := refTypeOfStruct.Field(i).Name

			if value, found := factory.Params[name]; found {
				nativeType := refElemStruct.Field(i).Type()
				paramType := reflect.TypeOf(value)

				fmt.Printf("===> 🚀 native: '%v', param: '%v'\n", nativeType, paramType)
				fmt.Printf("===> 🚀 param - name: '%v', value: '%v'\n", name, reflect.ValueOf(value))

				if paramType.String() == "*pflag.stringValue" {
					refElemStruct.Field(i).SetString(reflect.ValueOf(value).String())
				} else {
					if nativeType != paramType {
						message := fmt.Sprintf(ErrorTypes[MismatchNativeMemberValueTypeEn].Info[FormatEn],
							name, nativeType, paramType)

						panic(message)
					}
					refElemStruct.Field(i).Set(reflect.ValueOf(value))
				}
			} else {
				if options.Strict {
					message := fmt.Sprintf(ErrorTypes[MissingNativeMemberValueEn].Info[FormatEn], name)
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

// CreateParamSet creates native parameter set object from generic
// params with the default options. See CreateParameterSetWith for more
// details.
//
// Default options:
//
// "Strict": true
//
func (factory *ParamSetFactory[T]) Create() *T {

	return factory.CreateWith(CreatePsOptions{
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
// this scenario, CreateParamSetWith should be invoked with Strict set to true
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
func (factory *ParamSetFactory[T]) CreateWith(options CreatePsOptions) *T {
	target := new(T)
	factory.constructParameterSet(target, options)

	return target
}
