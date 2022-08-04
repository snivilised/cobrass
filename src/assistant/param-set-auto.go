package assistant

import (
	"net"
	"time"
)

// ----> auto generated(Build-ParamSet/gen-ps)

// BindBool binds bool slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindBool(info *FlagInfo, to *bool) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.BoolVar(to, info.FlagName(), info.Default.(bool), info.Usage)
	} else {
		flagSet.BoolVarP(to, info.FlagName(), info.Short, info.Default.(bool), info.Usage)
	}

	return params
}

// BindBoolSlice binds []bool slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindBoolSlice(info *FlagInfo, to *[]bool) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.BoolSliceVar(to, info.FlagName(), info.Default.([]bool), info.Usage)
	} else {
		flagSet.BoolSliceVarP(to, info.FlagName(), info.Short, info.Default.([]bool), info.Usage)
	}

	return params
}

// BindDuration binds time.Duration slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindDuration(info *FlagInfo, to *time.Duration) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.DurationVar(to, info.FlagName(), info.Default.(time.Duration), info.Usage)
	} else {
		flagSet.DurationVarP(to, info.FlagName(), info.Short, info.Default.(time.Duration), info.Usage)
	}

	return params
}

// BindValidatedDuration binds time.Duration slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of time.Duration type.
//
func (params *ParamSet[N]) BindValidatedDuration(info *FlagInfo, to *time.Duration, validator DurationValidatorFn) OptionValidator {
	params.BindDuration(info, to)

	wrapper := DurationOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindDurationSlice binds []time.Duration slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindDurationSlice(info *FlagInfo, to *[]time.Duration) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.DurationSliceVar(to, info.FlagName(), info.Default.([]time.Duration), info.Usage)
	} else {
		flagSet.DurationSliceVarP(to, info.FlagName(), info.Short, info.Default.([]time.Duration), info.Usage)
	}

	return params
}

// BindValidatedDurationSlice binds []time.Duration slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []time.Duration type.
//
func (params *ParamSet[N]) BindValidatedDurationSlice(info *FlagInfo, to *[]time.Duration, validator DurationSliceValidatorFn) OptionValidator {
	params.BindDurationSlice(info, to)

	wrapper := DurationSliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindEnum binds enum slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
// Note that normally the client would bind to a member of the native parameter
// set. However, since there is a discrepancy between the type of the native int
// based pseudo enum member and the equivalent acceptable string value typed by
// the user on the command line (idiomatically stored on the enum info), the
// client needs to extract the enum value from the enum info, something like this:
//
// paramSet.Native.Format = OutputFormatEnumInfo.Value()
//
// The best place to put this would be inside the PreRun/PreRunE function, assuming the
// param set and the enum info are both in scope. Actually, every int based enum
// flag, would need to have this assignment performed.
//
func (params *ParamSet[N]) BindEnum(info *FlagInfo, to *string) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.StringVar(to, info.FlagName(), info.Default.(string), info.Usage)
	} else {
		flagSet.StringVarP(to, info.FlagName(), info.Short, info.Default.(string), info.Usage)
	}

	return params
}

// BindValidatedEnum binds enum slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of enum type.
//
// Custom enum types created via the generic 'EnumInfo'/'EnumValue' come with a 'IsValid' method.
// The client can utilise this method inside a custom function passed into 'BindValidatedEnum'.
// The implementation would simply call this method, either on the EnumInfo or the EnumValue.
// Please see the readme for more details.
//
func (params *ParamSet[N]) BindValidatedEnum(info *FlagInfo, to *string, validator EnumValidatorFn) OptionValidator {
	params.BindEnum(info, to)

	wrapper := StringOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindFloat32 binds float32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindFloat32(info *FlagInfo, to *float32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float32Var(to, info.FlagName(), info.Default.(float32), info.Usage)
	} else {
		flagSet.Float32VarP(to, info.FlagName(), info.Short, info.Default.(float32), info.Usage)
	}

	return params
}

// BindValidatedFloat32 binds float32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of float32 type.
//
func (params *ParamSet[N]) BindValidatedFloat32(info *FlagInfo, to *float32, validator Float32ValidatorFn) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := Float32OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindFloat32Slice binds []float32 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindFloat32Slice(info *FlagInfo, to *[]float32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float32SliceVar(to, info.FlagName(), info.Default.([]float32), info.Usage)
	} else {
		flagSet.Float32SliceVarP(to, info.FlagName(), info.Short, info.Default.([]float32), info.Usage)
	}

	return params
}

// BindValidatedFloat32Slice binds []float32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []float32 type.
//
func (params *ParamSet[N]) BindValidatedFloat32Slice(info *FlagInfo, to *[]float32, validator Float32SliceValidatorFn) OptionValidator {
	params.BindFloat32Slice(info, to)

	wrapper := Float32SliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindFloat64 binds float64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindFloat64(info *FlagInfo, to *float64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float64Var(to, info.FlagName(), info.Default.(float64), info.Usage)
	} else {
		flagSet.Float64VarP(to, info.FlagName(), info.Short, info.Default.(float64), info.Usage)
	}

	return params
}

// BindValidatedFloat64 binds float64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of float64 type.
//
func (params *ParamSet[N]) BindValidatedFloat64(info *FlagInfo, to *float64, validator Float64ValidatorFn) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := Float64OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindFloat64Slice binds []float64 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindFloat64Slice(info *FlagInfo, to *[]float64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float64SliceVar(to, info.FlagName(), info.Default.([]float64), info.Usage)
	} else {
		flagSet.Float64SliceVarP(to, info.FlagName(), info.Short, info.Default.([]float64), info.Usage)
	}

	return params
}

// BindValidatedFloat64Slice binds []float64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []float64 type.
//
func (params *ParamSet[N]) BindValidatedFloat64Slice(info *FlagInfo, to *[]float64, validator Float64SliceValidatorFn) OptionValidator {
	params.BindFloat64Slice(info, to)

	wrapper := Float64SliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt binds int slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt(info *FlagInfo, to *int) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IntVar(to, info.FlagName(), info.Default.(int), info.Usage)
	} else {
		flagSet.IntVarP(to, info.FlagName(), info.Short, info.Default.(int), info.Usage)
	}

	return params
}

// BindValidatedInt binds int slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of int type.
//
func (params *ParamSet[N]) BindValidatedInt(info *FlagInfo, to *int, validator IntValidatorFn) OptionValidator {
	params.BindInt(info, to)

	wrapper := IntOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindIntSlice binds []int slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindIntSlice(info *FlagInfo, to *[]int) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IntSliceVar(to, info.FlagName(), info.Default.([]int), info.Usage)
	} else {
		flagSet.IntSliceVarP(to, info.FlagName(), info.Short, info.Default.([]int), info.Usage)
	}

	return params
}

// BindValidatedIntSlice binds []int slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []int type.
//
func (params *ParamSet[N]) BindValidatedIntSlice(info *FlagInfo, to *[]int, validator IntSliceValidatorFn) OptionValidator {
	params.BindIntSlice(info, to)

	wrapper := IntSliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt16 binds int16 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt16(info *FlagInfo, to *int16) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int16Var(to, info.FlagName(), info.Default.(int16), info.Usage)
	} else {
		flagSet.Int16VarP(to, info.FlagName(), info.Short, info.Default.(int16), info.Usage)
	}

	return params
}

// BindValidatedInt16 binds int16 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of int16 type.
//
func (params *ParamSet[N]) BindValidatedInt16(info *FlagInfo, to *int16, validator Int16ValidatorFn) OptionValidator {
	params.BindInt16(info, to)

	wrapper := Int16OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt32 binds int32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt32(info *FlagInfo, to *int32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int32Var(to, info.FlagName(), info.Default.(int32), info.Usage)
	} else {
		flagSet.Int32VarP(to, info.FlagName(), info.Short, info.Default.(int32), info.Usage)
	}

	return params
}

// BindValidatedInt32 binds int32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of int32 type.
//
func (params *ParamSet[N]) BindValidatedInt32(info *FlagInfo, to *int32, validator Int32ValidatorFn) OptionValidator {
	params.BindInt32(info, to)

	wrapper := Int32OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt32Slice binds []int32 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt32Slice(info *FlagInfo, to *[]int32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int32SliceVar(to, info.FlagName(), info.Default.([]int32), info.Usage)
	} else {
		flagSet.Int32SliceVarP(to, info.FlagName(), info.Short, info.Default.([]int32), info.Usage)
	}

	return params
}

// BindValidatedInt32Slice binds []int32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []int32 type.
//
func (params *ParamSet[N]) BindValidatedInt32Slice(info *FlagInfo, to *[]int32, validator Int32SliceValidatorFn) OptionValidator {
	params.BindInt32Slice(info, to)

	wrapper := Int32SliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt64 binds int64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt64(info *FlagInfo, to *int64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int64Var(to, info.FlagName(), info.Default.(int64), info.Usage)
	} else {
		flagSet.Int64VarP(to, info.FlagName(), info.Short, info.Default.(int64), info.Usage)
	}

	return params
}

// BindValidatedInt64 binds int64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of int64 type.
//
func (params *ParamSet[N]) BindValidatedInt64(info *FlagInfo, to *int64, validator Int64ValidatorFn) OptionValidator {
	params.BindInt64(info, to)

	wrapper := Int64OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt64Slice binds []int64 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt64Slice(info *FlagInfo, to *[]int64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int64SliceVar(to, info.FlagName(), info.Default.([]int64), info.Usage)
	} else {
		flagSet.Int64SliceVarP(to, info.FlagName(), info.Short, info.Default.([]int64), info.Usage)
	}

	return params
}

// BindValidatedInt64Slice binds []int64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []int64 type.
//
func (params *ParamSet[N]) BindValidatedInt64Slice(info *FlagInfo, to *[]int64, validator Int64SliceValidatorFn) OptionValidator {
	params.BindInt64Slice(info, to)

	wrapper := Int64SliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindInt8 binds int8 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindInt8(info *FlagInfo, to *int8) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int8Var(to, info.FlagName(), info.Default.(int8), info.Usage)
	} else {
		flagSet.Int8VarP(to, info.FlagName(), info.Short, info.Default.(int8), info.Usage)
	}

	return params
}

// BindValidatedInt8 binds int8 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of int8 type.
//
func (params *ParamSet[N]) BindValidatedInt8(info *FlagInfo, to *int8, validator Int8ValidatorFn) OptionValidator {
	params.BindInt8(info, to)

	wrapper := Int8OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindIPMask binds net.IPMask slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindIPMask(info *FlagInfo, to *net.IPMask) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IPMaskVar(to, info.FlagName(), info.Default.(net.IPMask), info.Usage)
	} else {
		flagSet.IPMaskVarP(to, info.FlagName(), info.Short, info.Default.(net.IPMask), info.Usage)
	}

	return params
}

// BindValidatedIPMask binds net.IPMask slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of net.IPMask type.
//
func (params *ParamSet[N]) BindValidatedIPMask(info *FlagInfo, to *net.IPMask, validator IPMaskValidatorFn) OptionValidator {
	params.BindIPMask(info, to)

	wrapper := IPMaskOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindIPNet binds net.IPNet slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindIPNet(info *FlagInfo, to *net.IPNet) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IPNetVar(to, info.FlagName(), info.Default.(net.IPNet), info.Usage)
	} else {
		flagSet.IPNetVarP(to, info.FlagName(), info.Short, info.Default.(net.IPNet), info.Usage)
	}

	return params
}

// BindValidatedIPNet binds net.IPNet slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of net.IPNet type.
//
func (params *ParamSet[N]) BindValidatedIPNet(info *FlagInfo, to *net.IPNet, validator IPNetValidatorFn) OptionValidator {
	params.BindIPNet(info, to)

	wrapper := IPNetOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindString binds string slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindString(info *FlagInfo, to *string) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.StringVar(to, info.FlagName(), info.Default.(string), info.Usage)
	} else {
		flagSet.StringVarP(to, info.FlagName(), info.Short, info.Default.(string), info.Usage)
	}

	return params
}

// BindValidatedString binds string slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of string type.
//
func (params *ParamSet[N]) BindValidatedString(info *FlagInfo, to *string, validator StringValidatorFn) OptionValidator {
	params.BindString(info, to)

	wrapper := StringOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindStringSlice binds []string slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindStringSlice(info *FlagInfo, to *[]string) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.StringSliceVar(to, info.FlagName(), info.Default.([]string), info.Usage)
	} else {
		flagSet.StringSliceVarP(to, info.FlagName(), info.Short, info.Default.([]string), info.Usage)
	}

	return params
}

// BindValidatedStringSlice binds []string slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []string type.
//
func (params *ParamSet[N]) BindValidatedStringSlice(info *FlagInfo, to *[]string, validator StringSliceValidatorFn) OptionValidator {
	params.BindStringSlice(info, to)

	wrapper := StringSliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindUint16 binds uint16 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindUint16(info *FlagInfo, to *uint16) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint16Var(to, info.FlagName(), info.Default.(uint16), info.Usage)
	} else {
		flagSet.Uint16VarP(to, info.FlagName(), info.Short, info.Default.(uint16), info.Usage)
	}

	return params
}

// BindValidatedUint16 binds uint16 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of uint16 type.
//
func (params *ParamSet[N]) BindValidatedUint16(info *FlagInfo, to *uint16, validator Uint16ValidatorFn) OptionValidator {
	params.BindUint16(info, to)

	wrapper := Uint16OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindUint32 binds uint32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindUint32(info *FlagInfo, to *uint32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint32Var(to, info.FlagName(), info.Default.(uint32), info.Usage)
	} else {
		flagSet.Uint32VarP(to, info.FlagName(), info.Short, info.Default.(uint32), info.Usage)
	}

	return params
}

// BindValidatedUint32 binds uint32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of uint32 type.
//
func (params *ParamSet[N]) BindValidatedUint32(info *FlagInfo, to *uint32, validator Uint32ValidatorFn) OptionValidator {
	params.BindUint32(info, to)

	wrapper := Uint32OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindUint64 binds uint64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindUint64(info *FlagInfo, to *uint64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint64Var(to, info.FlagName(), info.Default.(uint64), info.Usage)
	} else {
		flagSet.Uint64VarP(to, info.FlagName(), info.Short, info.Default.(uint64), info.Usage)
	}

	return params
}

// BindValidatedUint64 binds uint64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of uint64 type.
//
func (params *ParamSet[N]) BindValidatedUint64(info *FlagInfo, to *uint64, validator Uint64ValidatorFn) OptionValidator {
	params.BindUint64(info, to)

	wrapper := Uint64OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindUint8 binds uint8 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindUint8(info *FlagInfo, to *uint8) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint8Var(to, info.FlagName(), info.Default.(uint8), info.Usage)
	} else {
		flagSet.Uint8VarP(to, info.FlagName(), info.Short, info.Default.(uint8), info.Usage)
	}

	return params
}

// BindValidatedUint8 binds uint8 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of uint8 type.
//
func (params *ParamSet[N]) BindValidatedUint8(info *FlagInfo, to *uint8, validator Uint8ValidatorFn) OptionValidator {
	params.BindUint8(info, to)

	wrapper := Uint8OptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindUint binds uint slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) BindUint(info *FlagInfo, to *uint) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.UintVar(to, info.FlagName(), info.Default.(uint), info.Usage)
	} else {
		flagSet.UintVarP(to, info.FlagName(), info.Short, info.Default.(uint), info.Usage)
	}

	return params
}

// BindValidatedUint binds uint slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of uint type.
//
func (params *ParamSet[N]) BindValidatedUint(info *FlagInfo, to *uint, validator UintValidatorFn) OptionValidator {
	params.BindUint(info, to)

	wrapper := UintOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindUintSlice binds []uint slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) BindUintSlice(info *FlagInfo, to *[]uint) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.UintSliceVar(to, info.FlagName(), info.Default.([]uint), info.Usage)
	} else {
		flagSet.UintSliceVarP(to, info.FlagName(), info.Short, info.Default.([]uint), info.Usage)
	}

	return params
}

// BindValidatedUintSlice binds []uint slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of []uint type.
//
func (params *ParamSet[N]) BindValidatedUintSlice(info *FlagInfo, to *[]uint, validator UintSliceValidatorFn) OptionValidator {
	params.BindUintSlice(info, to)

	wrapper := UintSliceOptionValidator{
		Fn:    validator,
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// <---- end of auto generated
