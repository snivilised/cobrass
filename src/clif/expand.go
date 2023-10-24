package clif

// Expand returns a slice of strings representing all the positional arguments
// and flags/options to be executed by the third party program. before and
// flags are represented as a ThirdPartyCommandLine. This means that they can
// be represented by any slice of strings. However, since before represents
// positional args, those args are not expected to include any flags. Those
// flags would be specified in the flags parameter. after is optional and
// again represents further positional arguments.
func Expand(before ThirdPartyPositionalArgs, flags ThirdPartyCommandLine,
	after ...ThirdPartyPositional,
) ThirdPartyCommandLine {
	// approxTokensPerFlag: this is an approximate value because switches
	// do not need an option value
	const approxTokensPerFlag = 2

	capacity := len(before) + (len(flags) * approxTokensPerFlag) + len(after)
	allFlags := make([]string, 0, capacity)
	allFlags = append(allFlags, before...)
	allFlags = append(allFlags, flags...)
	allFlags = append(allFlags, after...)

	return allFlags
}
