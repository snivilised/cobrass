package utils

func TernaryIf[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func TernaryIfIf[T any](condA bool, atrue T, condB bool, vtrue, vfalse T) T {
	if condA {
		return atrue
	} else if condB {
		return vtrue
	}
	return vfalse
}
