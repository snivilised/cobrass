package adapters

import (
	"net"
	"time"
)

// ----> auto generated(Build-Validators/gen-ov)

// DurationValidatorFn defines the validator function for time.Duration type.
//
type DurationValidatorFn func(value time.Duration) error

// DurationOptionValidator defines the struct that wraps the client defined validator function
// DurationValidatorFn for time.Duration type. This is the instance that is returned by
// validated binder function BindValidatedDuration.
//
type DurationOptionValidator GenericOptionValidatorWrapper[time.Duration]

// Validate invokes the client defined validator function for time.Duration type.
//
func (validator DurationOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// DurationSliceOptionValidator defines the validator function for DurationSlice type.
//
type DurationSliceValidatorFn func(value []time.Duration) error

// DurationSliceOptionValidator wraps the client defined validator function for type []time.Duration.
//
type DurationSliceOptionValidator GenericOptionValidatorWrapper[[]time.Duration]

// Validate invokes the client defined validator function for []time.Duration type.
//
func (validator DurationSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// EnumValidatorFn defines the validator function for enum type.
//
type EnumValidatorFn func(value string) error

// EnumOptionValidator defines the struct that wraps the client defined validator function
// EnumValidatorFn for enum type. This is the instance that is returned by
// validated binder function BindValidatedEnum.
//
type EnumOptionValidator GenericOptionValidatorWrapper[string]

// Float32ValidatorFn defines the validator function for float32 type.
//
type Float32ValidatorFn func(value float32) error

// Float32OptionValidator defines the struct that wraps the client defined validator function
// Float32ValidatorFn for float32 type. This is the instance that is returned by
// validated binder function BindValidatedFloat32.
//
type Float32OptionValidator GenericOptionValidatorWrapper[float32]

// Validate invokes the client defined validator function for float32 type.
//
func (validator Float32OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Float32SliceOptionValidator defines the validator function for Float32Slice type.
//
type Float32SliceValidatorFn func(value []float32) error

// Float32SliceOptionValidator wraps the client defined validator function for type []float32.
//
type Float32SliceOptionValidator GenericOptionValidatorWrapper[[]float32]

// Validate invokes the client defined validator function for []float32 type.
//
func (validator Float32SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Float64ValidatorFn defines the validator function for float64 type.
//
type Float64ValidatorFn func(value float64) error

// Float64OptionValidator defines the struct that wraps the client defined validator function
// Float64ValidatorFn for float64 type. This is the instance that is returned by
// validated binder function BindValidatedFloat64.
//
type Float64OptionValidator GenericOptionValidatorWrapper[float64]

// Validate invokes the client defined validator function for float64 type.
//
func (validator Float64OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Float64SliceOptionValidator defines the validator function for Float64Slice type.
//
type Float64SliceValidatorFn func(value []float64) error

// Float64SliceOptionValidator wraps the client defined validator function for type []float64.
//
type Float64SliceOptionValidator GenericOptionValidatorWrapper[[]float64]

// Validate invokes the client defined validator function for []float64 type.
//
func (validator Float64SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// IntValidatorFn defines the validator function for int type.
//
type IntValidatorFn func(value int) error

// IntOptionValidator defines the struct that wraps the client defined validator function
// IntValidatorFn for int type. This is the instance that is returned by
// validated binder function BindValidatedInt.
//
type IntOptionValidator GenericOptionValidatorWrapper[int]

// Validate invokes the client defined validator function for int type.
//
func (validator IntOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// IntSliceOptionValidator defines the validator function for IntSlice type.
//
type IntSliceValidatorFn func(value []int) error

// IntSliceOptionValidator wraps the client defined validator function for type []int.
//
type IntSliceOptionValidator GenericOptionValidatorWrapper[[]int]

// Validate invokes the client defined validator function for []int type.
//
func (validator IntSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Int16ValidatorFn defines the validator function for int16 type.
//
type Int16ValidatorFn func(value int16) error

// Int16OptionValidator defines the struct that wraps the client defined validator function
// Int16ValidatorFn for int16 type. This is the instance that is returned by
// validated binder function BindValidatedInt16.
//
type Int16OptionValidator GenericOptionValidatorWrapper[int16]

// Validate invokes the client defined validator function for int16 type.
//
func (validator Int16OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Int32ValidatorFn defines the validator function for int32 type.
//
type Int32ValidatorFn func(value int32) error

// Int32OptionValidator defines the struct that wraps the client defined validator function
// Int32ValidatorFn for int32 type. This is the instance that is returned by
// validated binder function BindValidatedInt32.
//
type Int32OptionValidator GenericOptionValidatorWrapper[int32]

// Validate invokes the client defined validator function for int32 type.
//
func (validator Int32OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Int32SliceOptionValidator defines the validator function for Int32Slice type.
//
type Int32SliceValidatorFn func(value []int32) error

// Int32SliceOptionValidator wraps the client defined validator function for type []int32.
//
type Int32SliceOptionValidator GenericOptionValidatorWrapper[[]int32]

// Validate invokes the client defined validator function for []int32 type.
//
func (validator Int32SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Int64ValidatorFn defines the validator function for int64 type.
//
type Int64ValidatorFn func(value int64) error

// Int64OptionValidator defines the struct that wraps the client defined validator function
// Int64ValidatorFn for int64 type. This is the instance that is returned by
// validated binder function BindValidatedInt64.
//
type Int64OptionValidator GenericOptionValidatorWrapper[int64]

// Validate invokes the client defined validator function for int64 type.
//
func (validator Int64OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Int64SliceOptionValidator defines the validator function for Int64Slice type.
//
type Int64SliceValidatorFn func(value []int64) error

// Int64SliceOptionValidator wraps the client defined validator function for type []int64.
//
type Int64SliceOptionValidator GenericOptionValidatorWrapper[[]int64]

// Validate invokes the client defined validator function for []int64 type.
//
func (validator Int64SliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Int8ValidatorFn defines the validator function for int8 type.
//
type Int8ValidatorFn func(value int8) error

// Int8OptionValidator defines the struct that wraps the client defined validator function
// Int8ValidatorFn for int8 type. This is the instance that is returned by
// validated binder function BindValidatedInt8.
//
type Int8OptionValidator GenericOptionValidatorWrapper[int8]

// Validate invokes the client defined validator function for int8 type.
//
func (validator Int8OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// IPMaskValidatorFn defines the validator function for net.IPMask type.
//
type IPMaskValidatorFn func(value net.IPMask) error

// IPMaskOptionValidator defines the struct that wraps the client defined validator function
// IPMaskValidatorFn for net.IPMask type. This is the instance that is returned by
// validated binder function BindValidatedIPMask.
//
type IPMaskOptionValidator GenericOptionValidatorWrapper[net.IPMask]

// Validate invokes the client defined validator function for net.IPMask type.
//
func (validator IPMaskOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// IPNetValidatorFn defines the validator function for net.IPNet type.
//
type IPNetValidatorFn func(value net.IPNet) error

// IPNetOptionValidator defines the struct that wraps the client defined validator function
// IPNetValidatorFn for net.IPNet type. This is the instance that is returned by
// validated binder function BindValidatedIPNet.
//
type IPNetOptionValidator GenericOptionValidatorWrapper[net.IPNet]

// Validate invokes the client defined validator function for net.IPNet type.
//
func (validator IPNetOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// StringValidatorFn defines the validator function for string type.
//
type StringValidatorFn func(value string) error

// StringOptionValidator defines the struct that wraps the client defined validator function
// StringValidatorFn for string type. This is the instance that is returned by
// validated binder function BindValidatedString.
//
type StringOptionValidator GenericOptionValidatorWrapper[string]

// Validate invokes the client defined validator function for string type.
//
func (validator StringOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// StringSliceOptionValidator defines the validator function for StringSlice type.
//
type StringSliceValidatorFn func(value []string) error

// StringSliceOptionValidator wraps the client defined validator function for type []string.
//
type StringSliceOptionValidator GenericOptionValidatorWrapper[[]string]

// Validate invokes the client defined validator function for []string type.
//
func (validator StringSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Uint16ValidatorFn defines the validator function for uint16 type.
//
type Uint16ValidatorFn func(value uint16) error

// Uint16OptionValidator defines the struct that wraps the client defined validator function
// Uint16ValidatorFn for uint16 type. This is the instance that is returned by
// validated binder function BindValidatedUint16.
//
type Uint16OptionValidator GenericOptionValidatorWrapper[uint16]

// Validate invokes the client defined validator function for uint16 type.
//
func (validator Uint16OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Uint32ValidatorFn defines the validator function for uint32 type.
//
type Uint32ValidatorFn func(value uint32) error

// Uint32OptionValidator defines the struct that wraps the client defined validator function
// Uint32ValidatorFn for uint32 type. This is the instance that is returned by
// validated binder function BindValidatedUint32.
//
type Uint32OptionValidator GenericOptionValidatorWrapper[uint32]

// Validate invokes the client defined validator function for uint32 type.
//
func (validator Uint32OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Uint64ValidatorFn defines the validator function for uint64 type.
//
type Uint64ValidatorFn func(value uint64) error

// Uint64OptionValidator defines the struct that wraps the client defined validator function
// Uint64ValidatorFn for uint64 type. This is the instance that is returned by
// validated binder function BindValidatedUint64.
//
type Uint64OptionValidator GenericOptionValidatorWrapper[uint64]

// Validate invokes the client defined validator function for uint64 type.
//
func (validator Uint64OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// Uint8ValidatorFn defines the validator function for uint8 type.
//
type Uint8ValidatorFn func(value uint8) error

// Uint8OptionValidator defines the struct that wraps the client defined validator function
// Uint8ValidatorFn for uint8 type. This is the instance that is returned by
// validated binder function BindValidatedUint8.
//
type Uint8OptionValidator GenericOptionValidatorWrapper[uint8]

// Validate invokes the client defined validator function for uint8 type.
//
func (validator Uint8OptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// UintValidatorFn defines the validator function for uint type.
//
type UintValidatorFn func(value uint) error

// UintOptionValidator defines the struct that wraps the client defined validator function
// UintValidatorFn for uint type. This is the instance that is returned by
// validated binder function BindValidatedUint.
//
type UintOptionValidator GenericOptionValidatorWrapper[uint]

// Validate invokes the client defined validator function for uint type.
//
func (validator UintOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// UintSliceOptionValidator defines the validator function for UintSlice type.
//
type UintSliceValidatorFn func(value []uint) error

// UintSliceOptionValidator wraps the client defined validator function for type []uint.
//
type UintSliceOptionValidator GenericOptionValidatorWrapper[[]uint]

// Validate invokes the client defined validator function for []uint type.
//
func (validator UintSliceOptionValidator) Validate() error {
	return validator.Fn(*validator.Value)
}

// <---- end of auto generated
