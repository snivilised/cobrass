package assistant

import (
	"net"
	"time"

	"github.com/spf13/pflag"
)

// ----> auto generated(Build-Validators/gen-ov)

// DurationValidatorFn defines the validator function for time.Duration type.
//
type DurationValidatorFn func(time.Duration, *pflag.Flag) error

// DurationOptionValidator defines the struct that wraps the client defined validator function
// DurationValidatorFn for time.Duration type. This is the instance that is returned by
// validated binder function BindValidatedDuration.
//
type DurationOptionValidator GenericOptionValidatorWrapper[time.Duration]

// Validate invokes the client defined validator function for time.Duration type.
//
func (validator DurationOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// DurationSliceOptionValidator defines the validator function for DurationSlice type.
//
type DurationSliceValidatorFn func([]time.Duration, *pflag.Flag) error

// DurationSliceOptionValidator wraps the client defined validator function for type []time.Duration.
//
type DurationSliceOptionValidator GenericOptionValidatorWrapper[[]time.Duration]

// Validate invokes the client defined validator function for []time.Duration type.
//
func (validator DurationSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// EnumValidatorFn defines the validator function for enum type.
//
type EnumValidatorFn func(string, *pflag.Flag) error

// EnumOptionValidator defines the struct that wraps the client defined validator function
// EnumValidatorFn for enum type. This is the instance that is returned by
// validated binder function BindValidatedEnum.
//
type EnumOptionValidator GenericOptionValidatorWrapper[string]

// Float32ValidatorFn defines the validator function for float32 type.
//
type Float32ValidatorFn func(float32, *pflag.Flag) error

// Float32OptionValidator defines the struct that wraps the client defined validator function
// Float32ValidatorFn for float32 type. This is the instance that is returned by
// validated binder function BindValidatedFloat32.
//
type Float32OptionValidator GenericOptionValidatorWrapper[float32]

// Validate invokes the client defined validator function for float32 type.
//
func (validator Float32OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Float32SliceOptionValidator defines the validator function for Float32Slice type.
//
type Float32SliceValidatorFn func([]float32, *pflag.Flag) error

// Float32SliceOptionValidator wraps the client defined validator function for type []float32.
//
type Float32SliceOptionValidator GenericOptionValidatorWrapper[[]float32]

// Validate invokes the client defined validator function for []float32 type.
//
func (validator Float32SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Float64ValidatorFn defines the validator function for float64 type.
//
type Float64ValidatorFn func(float64, *pflag.Flag) error

// Float64OptionValidator defines the struct that wraps the client defined validator function
// Float64ValidatorFn for float64 type. This is the instance that is returned by
// validated binder function BindValidatedFloat64.
//
type Float64OptionValidator GenericOptionValidatorWrapper[float64]

// Validate invokes the client defined validator function for float64 type.
//
func (validator Float64OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Float64SliceOptionValidator defines the validator function for Float64Slice type.
//
type Float64SliceValidatorFn func([]float64, *pflag.Flag) error

// Float64SliceOptionValidator wraps the client defined validator function for type []float64.
//
type Float64SliceOptionValidator GenericOptionValidatorWrapper[[]float64]

// Validate invokes the client defined validator function for []float64 type.
//
func (validator Float64SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// IntValidatorFn defines the validator function for int type.
//
type IntValidatorFn func(int, *pflag.Flag) error

// IntOptionValidator defines the struct that wraps the client defined validator function
// IntValidatorFn for int type. This is the instance that is returned by
// validated binder function BindValidatedInt.
//
type IntOptionValidator GenericOptionValidatorWrapper[int]

// Validate invokes the client defined validator function for int type.
//
func (validator IntOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// IntSliceOptionValidator defines the validator function for IntSlice type.
//
type IntSliceValidatorFn func([]int, *pflag.Flag) error

// IntSliceOptionValidator wraps the client defined validator function for type []int.
//
type IntSliceOptionValidator GenericOptionValidatorWrapper[[]int]

// Validate invokes the client defined validator function for []int type.
//
func (validator IntSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Int16ValidatorFn defines the validator function for int16 type.
//
type Int16ValidatorFn func(int16, *pflag.Flag) error

// Int16OptionValidator defines the struct that wraps the client defined validator function
// Int16ValidatorFn for int16 type. This is the instance that is returned by
// validated binder function BindValidatedInt16.
//
type Int16OptionValidator GenericOptionValidatorWrapper[int16]

// Validate invokes the client defined validator function for int16 type.
//
func (validator Int16OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Int32ValidatorFn defines the validator function for int32 type.
//
type Int32ValidatorFn func(int32, *pflag.Flag) error

// Int32OptionValidator defines the struct that wraps the client defined validator function
// Int32ValidatorFn for int32 type. This is the instance that is returned by
// validated binder function BindValidatedInt32.
//
type Int32OptionValidator GenericOptionValidatorWrapper[int32]

// Validate invokes the client defined validator function for int32 type.
//
func (validator Int32OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Int32SliceOptionValidator defines the validator function for Int32Slice type.
//
type Int32SliceValidatorFn func([]int32, *pflag.Flag) error

// Int32SliceOptionValidator wraps the client defined validator function for type []int32.
//
type Int32SliceOptionValidator GenericOptionValidatorWrapper[[]int32]

// Validate invokes the client defined validator function for []int32 type.
//
func (validator Int32SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Int64ValidatorFn defines the validator function for int64 type.
//
type Int64ValidatorFn func(int64, *pflag.Flag) error

// Int64OptionValidator defines the struct that wraps the client defined validator function
// Int64ValidatorFn for int64 type. This is the instance that is returned by
// validated binder function BindValidatedInt64.
//
type Int64OptionValidator GenericOptionValidatorWrapper[int64]

// Validate invokes the client defined validator function for int64 type.
//
func (validator Int64OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Int64SliceOptionValidator defines the validator function for Int64Slice type.
//
type Int64SliceValidatorFn func([]int64, *pflag.Flag) error

// Int64SliceOptionValidator wraps the client defined validator function for type []int64.
//
type Int64SliceOptionValidator GenericOptionValidatorWrapper[[]int64]

// Validate invokes the client defined validator function for []int64 type.
//
func (validator Int64SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Int8ValidatorFn defines the validator function for int8 type.
//
type Int8ValidatorFn func(int8, *pflag.Flag) error

// Int8OptionValidator defines the struct that wraps the client defined validator function
// Int8ValidatorFn for int8 type. This is the instance that is returned by
// validated binder function BindValidatedInt8.
//
type Int8OptionValidator GenericOptionValidatorWrapper[int8]

// Validate invokes the client defined validator function for int8 type.
//
func (validator Int8OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// IPMaskValidatorFn defines the validator function for net.IPMask type.
//
type IPMaskValidatorFn func(net.IPMask, *pflag.Flag) error

// IPMaskOptionValidator defines the struct that wraps the client defined validator function
// IPMaskValidatorFn for net.IPMask type. This is the instance that is returned by
// validated binder function BindValidatedIPMask.
//
type IPMaskOptionValidator GenericOptionValidatorWrapper[net.IPMask]

// Validate invokes the client defined validator function for net.IPMask type.
//
func (validator IPMaskOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// IPNetValidatorFn defines the validator function for net.IPNet type.
//
type IPNetValidatorFn func(net.IPNet, *pflag.Flag) error

// IPNetOptionValidator defines the struct that wraps the client defined validator function
// IPNetValidatorFn for net.IPNet type. This is the instance that is returned by
// validated binder function BindValidatedIPNet.
//
type IPNetOptionValidator GenericOptionValidatorWrapper[net.IPNet]

// Validate invokes the client defined validator function for net.IPNet type.
//
func (validator IPNetOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// StringValidatorFn defines the validator function for string type.
//
type StringValidatorFn func(string, *pflag.Flag) error

// StringOptionValidator defines the struct that wraps the client defined validator function
// StringValidatorFn for string type. This is the instance that is returned by
// validated binder function BindValidatedString.
//
type StringOptionValidator GenericOptionValidatorWrapper[string]

// Validate invokes the client defined validator function for string type.
//
func (validator StringOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// StringSliceOptionValidator defines the validator function for StringSlice type.
//
type StringSliceValidatorFn func([]string, *pflag.Flag) error

// StringSliceOptionValidator wraps the client defined validator function for type []string.
//
type StringSliceOptionValidator GenericOptionValidatorWrapper[[]string]

// Validate invokes the client defined validator function for []string type.
//
func (validator StringSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Uint16ValidatorFn defines the validator function for uint16 type.
//
type Uint16ValidatorFn func(uint16, *pflag.Flag) error

// Uint16OptionValidator defines the struct that wraps the client defined validator function
// Uint16ValidatorFn for uint16 type. This is the instance that is returned by
// validated binder function BindValidatedUint16.
//
type Uint16OptionValidator GenericOptionValidatorWrapper[uint16]

// Validate invokes the client defined validator function for uint16 type.
//
func (validator Uint16OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Uint32ValidatorFn defines the validator function for uint32 type.
//
type Uint32ValidatorFn func(uint32, *pflag.Flag) error

// Uint32OptionValidator defines the struct that wraps the client defined validator function
// Uint32ValidatorFn for uint32 type. This is the instance that is returned by
// validated binder function BindValidatedUint32.
//
type Uint32OptionValidator GenericOptionValidatorWrapper[uint32]

// Validate invokes the client defined validator function for uint32 type.
//
func (validator Uint32OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Uint64ValidatorFn defines the validator function for uint64 type.
//
type Uint64ValidatorFn func(uint64, *pflag.Flag) error

// Uint64OptionValidator defines the struct that wraps the client defined validator function
// Uint64ValidatorFn for uint64 type. This is the instance that is returned by
// validated binder function BindValidatedUint64.
//
type Uint64OptionValidator GenericOptionValidatorWrapper[uint64]

// Validate invokes the client defined validator function for uint64 type.
//
func (validator Uint64OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// Uint8ValidatorFn defines the validator function for uint8 type.
//
type Uint8ValidatorFn func(uint8, *pflag.Flag) error

// Uint8OptionValidator defines the struct that wraps the client defined validator function
// Uint8ValidatorFn for uint8 type. This is the instance that is returned by
// validated binder function BindValidatedUint8.
//
type Uint8OptionValidator GenericOptionValidatorWrapper[uint8]

// Validate invokes the client defined validator function for uint8 type.
//
func (validator Uint8OptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// UintValidatorFn defines the validator function for uint type.
//
type UintValidatorFn func(uint, *pflag.Flag) error

// UintOptionValidator defines the struct that wraps the client defined validator function
// UintValidatorFn for uint type. This is the instance that is returned by
// validated binder function BindValidatedUint.
//
type UintOptionValidator GenericOptionValidatorWrapper[uint]

// Validate invokes the client defined validator function for uint type.
//
func (validator UintOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// UintSliceOptionValidator defines the validator function for UintSlice type.
//
type UintSliceValidatorFn func([]uint, *pflag.Flag) error

// UintSliceOptionValidator wraps the client defined validator function for type []uint.
//
type UintSliceOptionValidator GenericOptionValidatorWrapper[[]uint]

// Validate invokes the client defined validator function for []uint type.
//
func (validator UintSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value, validator.Flag)
}

// <---- end of auto generated
