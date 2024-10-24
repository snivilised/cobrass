package locale

import (
	"fmt"
)

// Non user facing internal error messages
//
// (These are messages that are due to programming errors,
// tat will mean nothing to the end user and are thus not
// translated).
//

// ❌ EnumValueValueAlreadyExists

// NewEnumValueValueAlreadyExistsNativeError enum already exists, invalid enum info specified
func NewEnumValueValueAlreadyExistsNativeError(value string, number int) error {
	return fmt.Errorf(
		"'%v (%v)' already exists, invalid enum info specified", value, number,
	)
}

// ❌ NewIsNotValidEnumValueNativeError

// NewIsNotValidEnumValueNativeError, is not a valid enum value
func NewIsNotValidEnumValueNativeError(value string) error {
	return fmt.Errorf(
		"'%v' is not a valid enum value", value,
	)
}

// ❌ failed to add validator

// failed to add validator for flag, because it already exists.
func NewFailedToAddValidatorAlreadyExistsNativeError(flag string) error {
	return fmt.Errorf(
		"failed to add validator for flag: '%v', because it already exists", flag,
	)
}

// ❌ NewCommandAlreadyRegisteredNativeError

// NewCommandAlreadyRegisteredNativeError, command already registered
func NewCommandAlreadyRegisteredNativeError(name string) error {
	return fmt.Errorf(
		"cobra container: command '%v' already registered", name,
	)
}

// ❌ NewParentCommandNotRegisteredNativeError

// NewParentCommandNotRegisteredNativeError, parent command not registered
func NewParentCommandNotRegisteredNativeError(parent string) error {
	return fmt.Errorf(
		"cobra container: parent command '%v' not registered", parent,
	)
}

// ❌ NewParamSetAlreadyRegisteredNativeError

// NewParamSetAlreadyRegisteredNativeError, param set already registered.
func NewParamSetAlreadyRegisteredNativeError(name string) error {
	return fmt.Errorf(
		"parameter set '%v' already registered", name,
	)
}

// ❌ NewParamSetObjectMustBeStructNativeError

// NewParamSetObjectMustBeStructNativeError, param set must be struct.
func NewParamSetObjectMustBeStructNativeError(name, typ string) error {
	return fmt.Errorf(
		"the native param set object ('%v') must be a struct, actual type: '%v'",
		name, typ,
	)
}

// ❌ NewParamSetObjectMustBePointerNativeError

// NewParamSetObjectMustBePointerNativeError, param set must be pointer.
func NewParamSetObjectMustBePointerNativeError(name, typ string) error {
	return fmt.Errorf(
		"the native param set object ('%v') must be a pointer, actual type: '%v'",
		name, typ,
	)
}

// ❌ NewParamSetNotFoundNativeError

// NewParamSetNotFoundNativeError, param set not found.
func NewParamSetNotFoundNativeError(name string) error {
	return fmt.Errorf(
		"parameter set '%v' not found", name,
	)
}
