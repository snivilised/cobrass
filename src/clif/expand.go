package clif

func Expand(before, flags ThirdPartyCommandLine, after ...ThirdPartyFlagName) []string {
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
