package assistant

import "fmt"

type ValidatorCollection map[string]OptionValidator

// ValidatorContainer manages the collection of client defined option validator
// functions.
//
type ValidatorContainer struct {
	validators ValidatorCollection
}

// ValidatorContainerOptions creation options
//
type ValidatorContainerOptions struct {
	// Size internal collection is initialised to
	//
	Size uint
}

// ValidatorContainerOptionFn definition ofa client defined function to
// set ValidatorContainer options.
//
type ValidatorContainerOptionFn func(o *ValidatorContainerOptions)

// NewValidatorContainer creates an initialised ValidatorContainer instance.
// To use default behaviour, invoke with no parameters.
//
func NewValidatorContainer(options ...ValidatorContainerOptionFn) *ValidatorContainer {

	option := ValidatorContainerOptions{
		Size: uint(1),
	}

	for _, functionalOption := range options {
		functionalOption(&option)
	}

	return &ValidatorContainer{
		validators: make(ValidatorCollection, option.Size),
	}
}

// Add adds the validator to the registered set of option validators. Only 1
// validator can be registered per flag, a panic will occur if the flag
// already has a validator registered for it.
//
func (container ValidatorContainer) Add(flag string, validator OptionValidator) {
	if _, found := container.validators[flag]; found {
		message := fmt.Errorf("failed to add validator for flag: '%v', because it already exists",
			flag)
		panic(message)
	}
	container.validators[flag] = validator
}

// Get returns the option validator for the specified flag, nil if
// not found.
//
func (container ValidatorContainer) Get(flag string) OptionValidator {
	if validator, found := container.validators[flag]; found {
		return validator
	}
	return nil
}

// run invokes all validators registered by calling their Validate method, which
// in turn, invokes the client defined validator function.
//
func (container ValidatorContainer) run() error {

	for _, validator := range container.validators {
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	return nil
}
