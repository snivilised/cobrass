package adapters

// OptionValidator wraps the user defined option validator function.
// This is the instance that is returned from the validated binder
// methods on the ParamSet.
//
type OptionValidator interface {
	Validate() error
}

// Needed because its not possible to create a type safe hetrogenious collection
// of objects that would be required for the ValidatorContainer
//
type GenericOptionValidatorWrapper[T any] struct {
	Fn    func(value T) error
	Value *T
}

func (validator GenericOptionValidatorWrapper[T]) Validate() error {
	// This method mysteriously doesn't satisfy the OptionValidator interface
	// resulting in every type having to define it with same implementation
	// itself, defeating the point of generics!
	//
	return validator.Fn(*validator.Value)
}

// CrossFieldValidator is a client function that is the callback passed into
// ParamSet.CrossValidate. Should be done after all parsed values have been bound
// and individually validated
//
type CrossFieldValidator[N any] func(native *N) error
