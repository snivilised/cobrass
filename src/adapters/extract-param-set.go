package adapters

import (
	"flag"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

// ExtactOptions specifies parameter set extraction options. Also see
// ParameterSetCreateOptions
//
type ExtactOptions struct {
	// Create species parameterset creation options
	//
	Create ParameterSetCreateOptions

	// Strategy determines which Visit method on the cobra FlagSet will
	// be invoked
	//
	Strategy VisitStrategyEnum
}

/*
ExtractParamSet is a helper function that the client can use to extract a native parameter
set object, denoted by type T. This is an alternative to using the CreateParameterSet
directly. The client needs to simply pass in the FlagSet instance on the cobra command,
ie 'command.Flags()'. The result is the native parameter set object.

Please see CreateParameterSet/CreateParameterSetWith for more details
*/
func ExtractParamSet[T comparable](flagSet *flag.FlagSet, options ExtactOptions) *T {
	target := new(T)
	genericSet := make(GenericParameterSet)

	handler := func(f *flag.Flag) {
		// The name has to be capitalised, because the native structure will have public
		// properties, which must begin with a capital letter
		//
		name := cases.Title(language.English).String(f.Name)
		genericSet[name] = f.Value
	}

	switch options.Strategy {
	case ActiveFlagsEn:
		flagSet.Visit(handler)
	case AllFlagsEn:
		flagSet.VisitAll(handler)
	}

	return constructParameterSet(target, genericSet, options.Create)
}
