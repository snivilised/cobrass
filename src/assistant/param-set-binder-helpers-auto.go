//nolint:gocritic // regexp.MustCompile: solution unknown
package assistant

import (
	"regexp"
	"time"

	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/cobrass/src/internal/third/lo"

	"github.com/spf13/pflag"
)

// ----> auto generated(Build-Predefined/gen-help)

// BindValidatedDurationWithin is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedDurationWithin(info *FlagInfo, to *time.Duration, low, high time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedDurationNotWithin is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedDurationWithin'.
func (params *ParamSet[N]) BindValidatedDurationNotWithin(info *FlagInfo, to *time.Duration, low, high time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsDuration is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsDuration(info *FlagInfo, to *time.Duration, collection []time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsDuration is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsDuration'.
func (params *ParamSet[N]) BindValidatedNotContainsDuration(info *FlagInfo, to *time.Duration, collection []time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedDurationGreaterThan is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedDurationGreaterThan(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedDurationAtLeast is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedDurationAtLeast(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedDurationLessThan is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedDurationLessThan(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedDurationAtMost is an alternative to using BindValidatedDuration. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedDurationAtMost(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {
	params.BindDuration(info, to)

	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsEnum is an alternative to using BindValidatedEnum. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsEnum(info *FlagInfo, to *string, collection []string) OptionValidator {
	params.BindEnum(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsEnum is an alternative to using BindValidatedEnum. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsEnum'.
func (params *ParamSet[N]) BindValidatedNotContainsEnum(info *FlagInfo, to *string, collection []string) OptionValidator {
	params.BindEnum(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat32Within is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedFloat32Within(info *FlagInfo, to *float32, low, high float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat32NotWithin is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedFloat32Within'.
func (params *ParamSet[N]) BindValidatedFloat32NotWithin(info *FlagInfo, to *float32, low, high float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsFloat32 is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsFloat32(info *FlagInfo, to *float32, collection []float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsFloat32 is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsFloat32'.
func (params *ParamSet[N]) BindValidatedNotContainsFloat32(info *FlagInfo, to *float32, collection []float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat32GreaterThan is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat32GreaterThan(info *FlagInfo, to *float32, threshold float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat32AtLeast is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat32AtLeast(info *FlagInfo, to *float32, threshold float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat32LessThan is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat32LessThan(info *FlagInfo, to *float32, threshold float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat32AtMost is an alternative to using BindValidatedFloat32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat32AtMost(info *FlagInfo, to *float32, threshold float32) OptionValidator {
	params.BindFloat32(info, to)

	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat64Within is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedFloat64Within(info *FlagInfo, to *float64, low, high float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat64NotWithin is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedFloat64Within'.
func (params *ParamSet[N]) BindValidatedFloat64NotWithin(info *FlagInfo, to *float64, low, high float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsFloat64 is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsFloat64(info *FlagInfo, to *float64, collection []float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsFloat64 is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsFloat64'.
func (params *ParamSet[N]) BindValidatedNotContainsFloat64(info *FlagInfo, to *float64, collection []float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat64GreaterThan is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat64GreaterThan(info *FlagInfo, to *float64, threshold float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat64AtLeast is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat64AtLeast(info *FlagInfo, to *float64, threshold float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat64LessThan is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat64LessThan(info *FlagInfo, to *float64, threshold float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedFloat64AtMost is an alternative to using BindValidatedFloat64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedFloat64AtMost(info *FlagInfo, to *float64, threshold float64) OptionValidator {
	params.BindFloat64(info, to)

	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedIntWithin is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedIntWithin(info *FlagInfo, to *int, low, high int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedIntNotWithin is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedIntWithin'.
func (params *ParamSet[N]) BindValidatedIntNotWithin(info *FlagInfo, to *int, low, high int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsInt is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsInt(info *FlagInfo, to *int, collection []int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsInt is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsInt'.
func (params *ParamSet[N]) BindValidatedNotContainsInt(info *FlagInfo, to *int, collection []int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedIntGreaterThan is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedIntGreaterThan(info *FlagInfo, to *int, threshold int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedIntAtLeast is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedIntAtLeast(info *FlagInfo, to *int, threshold int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedIntLessThan is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedIntLessThan(info *FlagInfo, to *int, threshold int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedIntAtMost is an alternative to using BindValidatedInt. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedIntAtMost(info *FlagInfo, to *int, threshold int) OptionValidator {
	params.BindInt(info, to)

	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt16Within is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedInt16Within(info *FlagInfo, to *int16, low, high int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt16NotWithin is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedInt16Within'.
func (params *ParamSet[N]) BindValidatedInt16NotWithin(info *FlagInfo, to *int16, low, high int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsInt16 is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsInt16(info *FlagInfo, to *int16, collection []int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsInt16 is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsInt16'.
func (params *ParamSet[N]) BindValidatedNotContainsInt16(info *FlagInfo, to *int16, collection []int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt16GreaterThan is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt16GreaterThan(info *FlagInfo, to *int16, threshold int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt16AtLeast is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt16AtLeast(info *FlagInfo, to *int16, threshold int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt16LessThan is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt16LessThan(info *FlagInfo, to *int16, threshold int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt16AtMost is an alternative to using BindValidatedInt16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt16AtMost(info *FlagInfo, to *int16, threshold int16) OptionValidator {
	params.BindInt16(info, to)

	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt32Within is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedInt32Within(info *FlagInfo, to *int32, low, high int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt32NotWithin is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedInt32Within'.
func (params *ParamSet[N]) BindValidatedInt32NotWithin(info *FlagInfo, to *int32, low, high int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsInt32 is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsInt32(info *FlagInfo, to *int32, collection []int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsInt32 is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsInt32'.
func (params *ParamSet[N]) BindValidatedNotContainsInt32(info *FlagInfo, to *int32, collection []int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt32GreaterThan is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt32GreaterThan(info *FlagInfo, to *int32, threshold int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt32AtLeast is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt32AtLeast(info *FlagInfo, to *int32, threshold int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt32LessThan is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt32LessThan(info *FlagInfo, to *int32, threshold int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt32AtMost is an alternative to using BindValidatedInt32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt32AtMost(info *FlagInfo, to *int32, threshold int32) OptionValidator {
	params.BindInt32(info, to)

	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt64Within is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedInt64Within(info *FlagInfo, to *int64, low, high int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt64NotWithin is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedInt64Within'.
func (params *ParamSet[N]) BindValidatedInt64NotWithin(info *FlagInfo, to *int64, low, high int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsInt64 is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsInt64(info *FlagInfo, to *int64, collection []int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsInt64 is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsInt64'.
func (params *ParamSet[N]) BindValidatedNotContainsInt64(info *FlagInfo, to *int64, collection []int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt64GreaterThan is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt64GreaterThan(info *FlagInfo, to *int64, threshold int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt64AtLeast is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt64AtLeast(info *FlagInfo, to *int64, threshold int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt64LessThan is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt64LessThan(info *FlagInfo, to *int64, threshold int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt64AtMost is an alternative to using BindValidatedInt64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt64AtMost(info *FlagInfo, to *int64, threshold int64) OptionValidator {
	params.BindInt64(info, to)

	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt8Within is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedInt8Within(info *FlagInfo, to *int8, low, high int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt8NotWithin is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedInt8Within'.
func (params *ParamSet[N]) BindValidatedInt8NotWithin(info *FlagInfo, to *int8, low, high int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsInt8 is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsInt8(info *FlagInfo, to *int8, collection []int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsInt8 is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsInt8'.
func (params *ParamSet[N]) BindValidatedNotContainsInt8(info *FlagInfo, to *int8, collection []int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt8GreaterThan is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt8GreaterThan(info *FlagInfo, to *int8, threshold int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt8AtLeast is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt8AtLeast(info *FlagInfo, to *int8, threshold int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt8LessThan is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedInt8LessThan(info *FlagInfo, to *int8, threshold int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedInt8AtMost is an alternative to using BindValidatedInt8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedInt8AtMost(info *FlagInfo, to *int8, threshold int8) OptionValidator {
	params.BindInt8(info, to)

	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringWithin is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedStringWithin(info *FlagInfo, to *string, low, high string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringNotWithin is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedStringWithin'.
func (params *ParamSet[N]) BindValidatedStringNotWithin(info *FlagInfo, to *string, low, high string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsString is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsString(info *FlagInfo, to *string, collection []string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsString is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsString'.
func (params *ParamSet[N]) BindValidatedNotContainsString(info *FlagInfo, to *string, collection []string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringIsMatch is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'pattern' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not match the regular expression denoted by 'pattern'.
func (params *ParamSet[N]) BindValidatedStringIsMatch(info *FlagInfo, to *string, pattern string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if regexp.MustCompile(pattern).Match([]byte(value)) {
				return nil
			}

			return locale.NewMatchOptValidationError(info.FlagName(), value, pattern)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringIsNotMatch is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'pattern' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedStringIsMatch'.
func (params *ParamSet[N]) BindValidatedStringIsNotMatch(info *FlagInfo, to *string, pattern string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(regexp.MustCompile(pattern).Match([]byte(value))) {
				return nil
			}

			return locale.NewNotMatchOptValidationError(info.FlagName(), value, pattern)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringGreaterThan is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedStringGreaterThan(info *FlagInfo, to *string, threshold string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringAtLeast is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedStringAtLeast(info *FlagInfo, to *string, threshold string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringLessThan is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedStringLessThan(info *FlagInfo, to *string, threshold string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedStringAtMost is an alternative to using BindValidatedString. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedStringAtMost(info *FlagInfo, to *string, threshold string) OptionValidator {
	params.BindString(info, to)

	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint16Within is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedUint16Within(info *FlagInfo, to *uint16, low, high uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint16NotWithin is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedUint16Within'.
func (params *ParamSet[N]) BindValidatedUint16NotWithin(info *FlagInfo, to *uint16, low, high uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsUint16 is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsUint16(info *FlagInfo, to *uint16, collection []uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsUint16 is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsUint16'.
func (params *ParamSet[N]) BindValidatedNotContainsUint16(info *FlagInfo, to *uint16, collection []uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint16GreaterThan is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint16GreaterThan(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint16AtLeast is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint16AtLeast(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint16LessThan is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint16LessThan(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint16AtMost is an alternative to using BindValidatedUint16. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint16AtMost(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {
	params.BindUint16(info, to)

	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint32Within is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedUint32Within(info *FlagInfo, to *uint32, low, high uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint32NotWithin is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedUint32Within'.
func (params *ParamSet[N]) BindValidatedUint32NotWithin(info *FlagInfo, to *uint32, low, high uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsUint32 is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsUint32(info *FlagInfo, to *uint32, collection []uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsUint32 is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsUint32'.
func (params *ParamSet[N]) BindValidatedNotContainsUint32(info *FlagInfo, to *uint32, collection []uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint32GreaterThan is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint32GreaterThan(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint32AtLeast is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint32AtLeast(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint32LessThan is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint32LessThan(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint32AtMost is an alternative to using BindValidatedUint32. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint32AtMost(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {
	params.BindUint32(info, to)

	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint64Within is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedUint64Within(info *FlagInfo, to *uint64, low, high uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint64NotWithin is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedUint64Within'.
func (params *ParamSet[N]) BindValidatedUint64NotWithin(info *FlagInfo, to *uint64, low, high uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsUint64 is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsUint64(info *FlagInfo, to *uint64, collection []uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsUint64 is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsUint64'.
func (params *ParamSet[N]) BindValidatedNotContainsUint64(info *FlagInfo, to *uint64, collection []uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint64GreaterThan is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint64GreaterThan(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint64AtLeast is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint64AtLeast(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint64LessThan is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint64LessThan(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint64AtMost is an alternative to using BindValidatedUint64. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint64AtMost(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {
	params.BindUint64(info, to)

	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint8Within is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedUint8Within(info *FlagInfo, to *uint8, low, high uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint8NotWithin is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedUint8Within'.
func (params *ParamSet[N]) BindValidatedUint8NotWithin(info *FlagInfo, to *uint8, low, high uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsUint8 is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsUint8(info *FlagInfo, to *uint8, collection []uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsUint8 is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsUint8'.
func (params *ParamSet[N]) BindValidatedNotContainsUint8(info *FlagInfo, to *uint8, collection []uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint8GreaterThan is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint8GreaterThan(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint8AtLeast is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint8AtLeast(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint8LessThan is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedUint8LessThan(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUint8AtMost is an alternative to using BindValidatedUint8. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUint8AtMost(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {
	params.BindUint8(info, to)

	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUintWithin is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method fails validation if the option value does not lie within 'low' and 'high' (inclusive).
func (params *ParamSet[N]) BindValidatedUintWithin(info *FlagInfo, to *uint, low, high uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= low && value <= high {
				return nil
			}

			return locale.NewWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUintNotWithin is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'low, high' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedUintWithin'.
func (params *ParamSet[N]) BindValidatedUintNotWithin(info *FlagInfo, to *uint, low, high uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(value >= low && value <= high) {
				return nil
			}

			return locale.NewNotWithinOptValidationError(info.FlagName(), value, low, high)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedContainsUint is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not a member of the 'collection' slice.
func (params *ParamSet[N]) BindValidatedContainsUint(info *FlagInfo, to *uint, collection []uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}

			return locale.NewContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedNotContainsUint is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'collection' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidatedContainsUint'.
func (params *ParamSet[N]) BindValidatedNotContainsUint(info *FlagInfo, to *uint, collection []uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}

			return locale.NewNotContainsOptValidationError(info.FlagName(), value, collection)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUintGreaterThan is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than 'threshold'.
func (params *ParamSet[N]) BindValidatedUintGreaterThan(info *FlagInfo, to *uint, threshold uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value > threshold {
				return nil
			}

			return locale.NewGreaterThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUintAtLeast is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably greater than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUintAtLeast(info *FlagInfo, to *uint, threshold uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value >= threshold {
				return nil
			}

			return locale.NewAtLeastOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUintLessThan is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than 'threshold'.
func (params *ParamSet[N]) BindValidatedUintLessThan(info *FlagInfo, to *uint, threshold uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value < threshold {
				return nil
			}

			return locale.NewLessThanOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// BindValidatedUintAtMost is an alternative to using BindValidatedUint. Instead of providing
// a function, the client passes in argument(s): 'threshold' to utilise predefined functionality as a helper.
// This method fails validation if the option value is not comparably less than or equal to 'threshold'.
func (params *ParamSet[N]) BindValidatedUintAtMost(info *FlagInfo, to *uint, threshold uint) OptionValidator {
	params.BindUint(info, to)

	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint, flag *pflag.Flag) error {
			if !flag.Changed {
				return nil
			}
			if value <= threshold {
				return nil
			}

			return locale.NewAtMostOptValidationError(info.FlagName(), value, threshold)
		},
		Value: to,
		Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
	}
	params.validators.Add(info.FlagName(), wrapper)

	return wrapper
}

// <---- end of auto generated
