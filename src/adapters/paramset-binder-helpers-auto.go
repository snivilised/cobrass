package adapters

import (
	"fmt"
	"regexp"
	"time"

	"github.com/samber/lo"
)

/*
not-able?

pre-define validators

-- intX, floatX, uintX, duration (make generic on comparable types)
within (ie range)
greaterThan
atLeast
lessThan
atMost

-- string
oneOf
minLen
maxLen
match

==> lanaguage dependent (need to accept a language setting)
isMonth
isDay
-- enum (generic)
*/

// ----> auto generated(Build-Predefined/gen-help)

// BindValidatedEnumWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumWithin(info *FlagInfo, to *string, lo, hi string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedEnumNotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumNotWithin(info *FlagInfo, to *string, lo, hi string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsEnum (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsEnum(info *FlagInfo, to *string, collection []string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedEnumNotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumNotContains(info *FlagInfo, to *string, collection []string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedEnumGreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumGreaterThan(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedEnumAtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumAtLeast(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedEnumLessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumLessThan(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedEnumAtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedEnumAtMost(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindEnum(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringWithin(info *FlagInfo, to *string, lo, hi string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringNotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringNotWithin(info *FlagInfo, to *string, lo, hi string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsString (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsString(info *FlagInfo, to *string, collection []string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringNotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringNotContains(info *FlagInfo, to *string, collection []string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringIsMatch (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringIsMatch(info *FlagInfo, to *string, pattern string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if regexp.MustCompile(pattern).Match([]byte(value)) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', does not match: [%v]",
				info.FlagName(), value, pattern,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringIsNotMatch (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringIsNotMatch(info *FlagInfo, to *string, pattern string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if !(regexp.MustCompile(pattern).Match([]byte(value))) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', matches: [%v]",
				info.FlagName(), value, pattern,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringGreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringGreaterThan(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringAtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringAtLeast(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringLessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringLessThan(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedStringAtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedStringAtMost(info *FlagInfo, to *string, threshold string) OptionValidator {

	params.BindString(info, to)
	wrapper := GenericOptionValidatorWrapper[string]{
		Fn: func(value string) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntWithin(info *FlagInfo, to *int, lo, hi int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntNotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntNotWithin(info *FlagInfo, to *int, lo, hi int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsInt (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsInt(info *FlagInfo, to *int, collection []int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntNotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntNotContains(info *FlagInfo, to *int, collection []int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntGreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntGreaterThan(info *FlagInfo, to *int, threshold int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntAtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntAtLeast(info *FlagInfo, to *int, threshold int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntLessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntLessThan(info *FlagInfo, to *int, threshold int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedIntAtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedIntAtMost(info *FlagInfo, to *int, threshold int) OptionValidator {

	params.BindInt(info, to)
	wrapper := GenericOptionValidatorWrapper[int]{
		Fn: func(value int) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8Within(info *FlagInfo, to *int8, lo, hi int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8NotWithin(info *FlagInfo, to *int8, lo, hi int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsInt8 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsInt8(info *FlagInfo, to *int8, collection []int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8NotContains(info *FlagInfo, to *int8, collection []int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8GreaterThan(info *FlagInfo, to *int8, threshold int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8AtLeast(info *FlagInfo, to *int8, threshold int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8LessThan(info *FlagInfo, to *int8, threshold int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt8AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt8AtMost(info *FlagInfo, to *int8, threshold int8) OptionValidator {

	params.BindInt8(info, to)
	wrapper := GenericOptionValidatorWrapper[int8]{
		Fn: func(value int8) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16Within(info *FlagInfo, to *int16, lo, hi int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16NotWithin(info *FlagInfo, to *int16, lo, hi int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsInt16 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsInt16(info *FlagInfo, to *int16, collection []int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16NotContains(info *FlagInfo, to *int16, collection []int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16GreaterThan(info *FlagInfo, to *int16, threshold int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16AtLeast(info *FlagInfo, to *int16, threshold int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16LessThan(info *FlagInfo, to *int16, threshold int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt16AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt16AtMost(info *FlagInfo, to *int16, threshold int16) OptionValidator {

	params.BindInt16(info, to)
	wrapper := GenericOptionValidatorWrapper[int16]{
		Fn: func(value int16) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32Within(info *FlagInfo, to *int32, lo, hi int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32NotWithin(info *FlagInfo, to *int32, lo, hi int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsInt32 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsInt32(info *FlagInfo, to *int32, collection []int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32NotContains(info *FlagInfo, to *int32, collection []int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32GreaterThan(info *FlagInfo, to *int32, threshold int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32AtLeast(info *FlagInfo, to *int32, threshold int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32LessThan(info *FlagInfo, to *int32, threshold int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt32AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt32AtMost(info *FlagInfo, to *int32, threshold int32) OptionValidator {

	params.BindInt32(info, to)
	wrapper := GenericOptionValidatorWrapper[int32]{
		Fn: func(value int32) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64Within(info *FlagInfo, to *int64, lo, hi int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64NotWithin(info *FlagInfo, to *int64, lo, hi int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsInt64 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsInt64(info *FlagInfo, to *int64, collection []int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64NotContains(info *FlagInfo, to *int64, collection []int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64GreaterThan(info *FlagInfo, to *int64, threshold int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64AtLeast(info *FlagInfo, to *int64, threshold int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64LessThan(info *FlagInfo, to *int64, threshold int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedInt64AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedInt64AtMost(info *FlagInfo, to *int64, threshold int64) OptionValidator {

	params.BindInt64(info, to)
	wrapper := GenericOptionValidatorWrapper[int64]{
		Fn: func(value int64) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintWithin(info *FlagInfo, to *uint, lo, hi uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintNotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintNotWithin(info *FlagInfo, to *uint, lo, hi uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsUint (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsUint(info *FlagInfo, to *uint, collection []uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintNotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintNotContains(info *FlagInfo, to *uint, collection []uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintGreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintGreaterThan(info *FlagInfo, to *uint, threshold uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintAtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintAtLeast(info *FlagInfo, to *uint, threshold uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintLessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintLessThan(info *FlagInfo, to *uint, threshold uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUintAtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUintAtMost(info *FlagInfo, to *uint, threshold uint) OptionValidator {

	params.BindUint(info, to)
	wrapper := GenericOptionValidatorWrapper[uint]{
		Fn: func(value uint) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8Within(info *FlagInfo, to *uint8, lo, hi uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8NotWithin(info *FlagInfo, to *uint8, lo, hi uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsUint8 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsUint8(info *FlagInfo, to *uint8, collection []uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8NotContains(info *FlagInfo, to *uint8, collection []uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8GreaterThan(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8AtLeast(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8LessThan(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint8AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint8AtMost(info *FlagInfo, to *uint8, threshold uint8) OptionValidator {

	params.BindUint8(info, to)
	wrapper := GenericOptionValidatorWrapper[uint8]{
		Fn: func(value uint8) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16Within(info *FlagInfo, to *uint16, lo, hi uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16NotWithin(info *FlagInfo, to *uint16, lo, hi uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsUint16 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsUint16(info *FlagInfo, to *uint16, collection []uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16NotContains(info *FlagInfo, to *uint16, collection []uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16GreaterThan(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16AtLeast(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16LessThan(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint16AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint16AtMost(info *FlagInfo, to *uint16, threshold uint16) OptionValidator {

	params.BindUint16(info, to)
	wrapper := GenericOptionValidatorWrapper[uint16]{
		Fn: func(value uint16) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32Within(info *FlagInfo, to *uint32, lo, hi uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32NotWithin(info *FlagInfo, to *uint32, lo, hi uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsUint32 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsUint32(info *FlagInfo, to *uint32, collection []uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32NotContains(info *FlagInfo, to *uint32, collection []uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32GreaterThan(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32AtLeast(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32LessThan(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint32AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint32AtMost(info *FlagInfo, to *uint32, threshold uint32) OptionValidator {

	params.BindUint32(info, to)
	wrapper := GenericOptionValidatorWrapper[uint32]{
		Fn: func(value uint32) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64Within(info *FlagInfo, to *uint64, lo, hi uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64NotWithin(info *FlagInfo, to *uint64, lo, hi uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsUint64 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsUint64(info *FlagInfo, to *uint64, collection []uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64NotContains(info *FlagInfo, to *uint64, collection []uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64GreaterThan(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64AtLeast(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64LessThan(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedUint64AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedUint64AtMost(info *FlagInfo, to *uint64, threshold uint64) OptionValidator {

	params.BindUint64(info, to)
	wrapper := GenericOptionValidatorWrapper[uint64]{
		Fn: func(value uint64) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32Within(info *FlagInfo, to *float32, lo, hi float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32NotWithin(info *FlagInfo, to *float32, lo, hi float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsFloat32 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsFloat32(info *FlagInfo, to *float32, collection []float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32NotContains(info *FlagInfo, to *float32, collection []float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32GreaterThan(info *FlagInfo, to *float32, threshold float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32AtLeast(info *FlagInfo, to *float32, threshold float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32LessThan(info *FlagInfo, to *float32, threshold float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat32AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat32AtMost(info *FlagInfo, to *float32, threshold float32) OptionValidator {

	params.BindFloat32(info, to)
	wrapper := GenericOptionValidatorWrapper[float32]{
		Fn: func(value float32) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64Within (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64Within(info *FlagInfo, to *float64, lo, hi float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64NotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64NotWithin(info *FlagInfo, to *float64, lo, hi float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsFloat64 (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsFloat64(info *FlagInfo, to *float64, collection []float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64NotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64NotContains(info *FlagInfo, to *float64, collection []float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64GreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64GreaterThan(info *FlagInfo, to *float64, threshold float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64AtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64AtLeast(info *FlagInfo, to *float64, threshold float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64LessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64LessThan(info *FlagInfo, to *float64, threshold float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedFloat64AtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedFloat64AtMost(info *FlagInfo, to *float64, threshold float64) OptionValidator {

	params.BindFloat64(info, to)
	wrapper := GenericOptionValidatorWrapper[float64]{
		Fn: func(value float64) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationWithin(info *FlagInfo, to *time.Duration, lo, hi time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if value >= lo && value <= hi {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', out of range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationNotWithin (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationNotWithin(info *FlagInfo, to *time.Duration, lo, hi time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if !(value >= lo && value <= hi) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is within range: [%v]..[%v]",
				info.FlagName(), value, lo, hi,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedContainsDuration (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedContainsDuration(info *FlagInfo, to *time.Duration, collection []time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if lo.IndexOf(collection, value) >= 0 {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationNotContains (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationNotContains(info *FlagInfo, to *time.Duration, collection []time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if !(lo.IndexOf(collection, value) >= 0) {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', is a member of: [%v]",
				info.FlagName(), value, collection,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationGreaterThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationGreaterThan(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if value > threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not greater than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationAtLeast (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationAtLeast(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if value >= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at least: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationLessThan (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationLessThan(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if value < threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not less than: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// BindValidatedDurationAtMost (documentation comment pending ...)
//
func (params *ParamSet[N]) BindValidatedDurationAtMost(info *FlagInfo, to *time.Duration, threshold time.Duration) OptionValidator {

	params.BindDuration(info, to)
	wrapper := GenericOptionValidatorWrapper[time.Duration]{
		Fn: func(value time.Duration) error {
			if value <= threshold {
				return nil
			}
			return fmt.Errorf("(%v): option validation failed, '%v', not at most: [%v]",
				info.FlagName(), value, threshold,
			)
		},
		Value: to,
	}
	params.validatorGroup.Add(info.FlagName(), wrapper)
	return wrapper
}

// <---- end of auto generated
