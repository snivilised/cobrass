package adapters

import "fmt"

type ValidatorCollection map[string]OptionValidator

// ValidatorGroup manages the collection of client defined option validator
// functions.
//
type ValidatorGroup struct {
	validators ValidatorCollection
}

// ValidatorGroupOptions creation options
//
type ValidatorGroupOptions struct {
	// Size internal collection is initialised to
	//
	Size uint
}

// NewValidatorGroup creates an initialised ValidatorGroup instance. To
// use defaults, pass in nil for options.
//
func NewValidatorGroup(options *ValidatorGroupOptions) *ValidatorGroup {
	size := uint(1)
	if options != nil && options.Size > 0 {
		size = options.Size
	}
	return &ValidatorGroup{
		validators: make(ValidatorCollection, size),
	}
}

func (group ValidatorGroup) Add(flag string, validator OptionValidator) {
	if _, found := group.validators[flag]; found {
		message := fmt.Errorf("failed to add validator for flag: '%v', because it already exists",
			flag)
		panic(message)
	}
	group.validators[flag] = validator
}

// Get returns the option validator for the specified flag, nil if
// not found.
//
func (group ValidatorGroup) Get(flag string) OptionValidator {
	if validator, found := group.validators[flag]; found {
		return validator
	}
	return nil
}

// Run invokes all validators registered by calling their Vaildate method, which
// in turn, invokes the client defined validator function.
//
func (group ValidatorGroup) Run() error {

	for _, validator := range group.validators {
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	return nil
}
