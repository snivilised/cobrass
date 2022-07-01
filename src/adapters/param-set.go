package adapters

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// FlagInfo collates together the paramters passed into the bind methods
// The Bind methods are just a wrapper around invoking the type based methods
// on the cobra flag set in order to define flags.
//
type FlagInfo struct {
	Name      string
	Usage     string
	Short     string
	Default   any
	Validator StringValidatorFn
}

func extractNameFromUsage(usage string) string {
	name := usage

	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// NewFlagInfo factory function for FlagInfo
//
func NewFlagInfo(usage string, short string, def any) *FlagInfo {

	return &FlagInfo{
		Name:    extractNameFromUsage(usage),
		Usage:   usage,
		Short:   short,
		Default: def,
	}
}

// FlagName returns the name of the flag derived from the Usage
//
func (info *FlagInfo) FlagName() string {
	return info.Name
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
	validators *ValidatorContainer
	// Native is the native client defined parameter set instance, which
	// must be a struct.
	//
	Native  *N
	FlagSet *pflag.FlagSet
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
	ps.FlagSet = command.Flags()
	ps.Native = new(N)

	if reflect.TypeOf(*ps.Native).Kind() != reflect.Struct {
		name := reflect.TypeOf(*ps.Native).Name()
		panic(fmt.Sprintf("the native param set object ('%v') must be a structure", name))
	}
	ps.validators = NewValidatorContainer(nil)
	return ps
}

// Validators returns the compound validator that the client will need to invoke
// option validation (Run), typically inside the Run function defined on
// the cobra command.
//
func (params *ParamSet[N]) Validators() *ValidatorContainer {
	return params.validators
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

// func (params *ParamSet[N]) BindEnum(info *FlagInfo, to *string) {
// 	params.FlagSet.StringVarP(to, info.Name, info.Short, info.Default.(string), info.Usage)
// }

// We can also defined pre-defined validators !! such as a string value must be of limited
// length, or with a set of predefined values (like ValidArgs)
//

// BindValidatedString binds string slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of string type.
//
// func (params *ParamSet[N]) BindValidatedEnum(info *FlagInfo, to *string, validator StringValidatorFn) OptionValidator {

// 	params.BindEnum(info, to)
// 	wrapper := &StringOptionValidator{
// 		fn:    validator,
// 		value: to,
// 	}
// 	params.validatorGroup.Add(info.Name, wrapper)
// 	return wrapper
// }
