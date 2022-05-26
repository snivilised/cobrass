package utils

/*
TernaryIf, implements the missing ternary operator in go in the form
	cond ? vtrue : vfalse
*/
func TernaryIf[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

/*
TernaryIfIf, implements a double layer ternary operator of the form:
	condA ? atrue : (condB? vtrue : vfalse)
*/
func TernaryIfIf[T any](condA bool, atrue T, condB bool, vtrue, vfalse T) T {
	if condA {
		return atrue
	} else if condB {
		return vtrue
	}
	return vfalse
}
