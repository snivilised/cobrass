package clif

import (
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var booleans = []string{"true", "false"}

type (
	tokenInput struct {
		token        string
		lead         string
		bare         string
		optionValue  string
		existingCL   ThirdPartyCommandLine
		presentFlags ChangedFlagsMap
		knownBy      KnownByCollection
	}

	handleTokenResult struct {
		doConcatenate bool
	}

	concatIfFunc      func(input *tokenInput) *handleTokenResult
	concatenateResult struct {
		commandLine  ThirdPartyCommandLine
		handleResult handleTokenResult
	}
)

func (i *tokenInput) consume(nextIndex int, secondaryCL ThirdPartyCommandLine) int {
	const (
		unaryIncrement = 1
		pairIncrement  = 2
	)

	handleAsPair := false

	if nextIndex < len(secondaryCL) {
		next := secondaryCL[nextIndex]
		nextLead, nextBare := split(next)

		if strings.HasPrefix(i.lead, "-") && !strings.HasPrefix(nextLead, "-") {
			i.optionValue = nextBare
			handleAsPair = true
		}
	}

	return lo.Ternary(handleAsPair, pairIncrement, unaryIncrement)
}

func (i *tokenInput) concatIf(concatFunc concatIfFunc) *concatenateResult {
	handleResult := concatFunc(i)

	if handleResult.doConcatenate {
		i.existingCL = append(i.existingCL, i.token)

		if i.optionValue != "" {
			i.existingCL = append(i.existingCL, i.optionValue)
		}
	}

	return &concatenateResult{
		commandLine:  i.existingCL,
		handleResult: *handleResult,
	}
}

func notInPresent(input *tokenInput) *handleTokenResult {
	var (
		result = &handleTokenResult{}
	)

	if input.lead == "" {
		result.doConcatenate = true

		return result
	}

	if _, found := input.presentFlags[input.bare]; found {
		return result
	}

	aka := input.knownBy[input.bare]
	_, found := input.presentFlags[aka]
	result.doConcatenate = !found

	return result
}

// Evaluate merges the secondary command line with the specified flags.
// The flags that occur in specified take precedence over those in
// secondary. There is a slight complication caused by the fact that
// a flag in the specified set may be in the secondary set but in the opposite
// form; eg a flag may be in its short from in specified but in long form
// in secondary. This is resolved by the knownBy set. The specified set
// contains flags in their bare long form.
func Evaluate(presentFlags ChangedFlagsMap,
	knownBy KnownByCollection,
	secondaryCL ThirdPartyCommandLine,
) ThirdPartyCommandLine {
	result := &concatenateResult{}
	bilateralKnownBy := composeBilateral(knownBy)

	result.commandLine = spreadFlags(presentFlags)

	if len(secondaryCL) == 0 {
		return result.commandLine
	}

	if len(secondaryCL) == 1 {
		token := secondaryCL[0]
		lead, bare := split(token)

		input := &tokenInput{
			token:        token,
			lead:         lead,
			bare:         bare,
			existingCL:   result.commandLine,
			presentFlags: presentFlags,
			knownBy:      bilateralKnownBy,
		}
		result = input.concatIf(notInPresent)

		return result.commandLine
	}

	for t, n := 0, 1; t < len(secondaryCL); {
		token := secondaryCL[t]
		lead, bare := split(token)

		input := &tokenInput{
			token:        token,
			lead:         lead,
			bare:         bare,
			existingCL:   result.commandLine,
			presentFlags: presentFlags,
			knownBy:      bilateralKnownBy,
		}
		increment := input.consume(n, secondaryCL)
		result = input.concatIf(notInPresent)

		t += increment
		n += increment
	}

	return result.commandLine
}

func split(token string) (string, string) { //nolint:gocritic // pedant
	var (
		lead string
		bare = token
	)

	if strings.HasPrefix(token, "--") {
		lead = "--"
		bare = token[2:]
	} else if strings.HasPrefix(token, "-") {
		lead = "-"
		bare = token[1:]
	}

	return lead, bare
}

func spreadFlags(presentFlags ChangedFlagsMap) ThirdPartyCommandLine {
	commandLine := ThirdPartyCommandLine{}

	for _, flag := range presentFlags.Keys() {
		option := presentFlags[flag]
		dash := lo.Ternary(len(flag) == 1, "-", "--")
		prefixed := dash + flag
		withOption := !slices.Contains(booleans, option)

		commandLine = append(commandLine, prefixed)

		if withOption {
			commandLine = append(commandLine, option)
		}
	}

	return commandLine
}

func composeBilateral(knownBy KnownByCollection) KnownByCollection {
	const twice = 2
	bilateral := make(KnownByCollection, len(knownBy)*twice)

	maps.Copy(bilateral, knownBy)

	for long, short := range knownBy {
		bilateral[short] = long
	}

	return bilateral
}
