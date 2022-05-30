package adapters

import "fmt"

// InsertSafePositional inserts positional parameter into genenric set and
// ensures that a positional does clash with an extsing named flag
// Panics, if the required parameter already exists.
//
func InsertSafePositional(generic GenericParameterSet, key, value string) {
	if existing, ok := generic[key]; ok {
		panic(fmt.Sprintf("failed to insert '%v' at '%v', already exists as: '%v'",
			value, key, existing))
	}
	generic[key] = value
}
