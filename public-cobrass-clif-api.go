package cobrass

import (
	"github.com/snivilised/cobrass/src/clif"
)

type (
	// ThirdPartyFlagName raw name of a flag, ie without the leading --/-
	ThirdPartyFlagName = clif.ThirdPartyFlagName

	// ThirdPartyPositional represents a positional argument
	ThirdPartyPositional = clif.ThirdPartyPositional

	// ThirdPartyPositionalArgs represents a collection of third party
	// positional arguments.
	ThirdPartyPositionalArgs = clif.ThirdPartyPositionalArgs

	// ThirdPartyOptionValue the string value of an option. Since this option
	// is being delegated to a third party command, it does not have to be
	// of a particular native go type and can be composed from a go type
	// using the value's String() method.
	ThirdPartyOptionValue = clif.ThirdPartyOptionValue

	// SpecifiedFlagsCollection represents the set of third party flags
	// presented by the user on the command line.
	// (NB: Cobra does not currently have a mechanism to collect third
	// party flags, by convention, anything that follows " -- "), therefore
	// we need to collect and handle these flags/options explicitly,
	// which is less than ideal.
	// A difference between SpecifiedFlagsCollection and ThirdPartyCommandLine
	// is that switch flags have a true/false option value in SpecifiedFlagsCollection
	// but not in ThirdPartyCommandLine.
	SpecifiedFlagsCollection = clif.SpecifiedFlagsCollection

	// ThirdPartyPresentFlags (see SpecifiedFlagsCollection)
	ThirdPartyPresentFlags = clif.ThirdPartyPresentFlags

	// KnownByCollection collection maps a full flag name to the
	// short name it is also known by. If a flag does not
	// have a short name, it should be mapped to the empty
	// string.
	KnownByCollection = clif.KnownByCollection

	// ThirdPartyFlagKnownBy (see KnownByCollection).
	ThirdPartyFlagKnownBy = clif.ThirdPartyFlagsKnownBy

	// ThirdPartyCommandLine represents the collection of flags
	// used to invoke a third party command. This collection
	// represents the raw flags used for the invocation in
	// the order required by the third party command. It also means
	// that this collection contains the leading --/- not just
	// the names of the flags and options.
	// For example, to invoke the magick command we may want to
	// compose this collection with:
	// magick --strip --interlace plane --gaussian-blur 0.05
	// and in this case, the list would be defined as a string slice:
	// []string{"--strip", "--interlace", "plane", "--gaussian-blur", "0.05"}
	ThirdPartyCommandLine = clif.ThirdPartyCommandLine

	// ExternalThirdParty base struct for cli applications using the
	// entry paradigm that need to delegate an invocation to an
	// external third party command.
	ExternalThirdParty = clif.ExternalThirdParty
)

var (
	// Evaluate merges the secondary command line with the specified flags.
	// The flags that occur in specified take precedence over those in
	// secondary. There is a slight complication caused by the fact that
	// a flag in the specified set may be in the secondary set but in the opposite
	// form; eg a flag may be in its short from in specified but in long form
	// in secondary. This is resolved by the knownBy set. The specified set
	// contains flags in their bare long form (bare as in without dash prefix).
	Evaluate = clif.Evaluate

	// Expand returns a slice of strings representing the positional arguments and
	// flags/options to be executed by the third party program. before and flags
	// are represented as a ThirdPartyCommandLine. This means that they can be
	// represented by any slice of strings. However, since before represents
	// positional args, those args are not expected to include any flags. Those
	// flags would be specified in the flags parameter. after is optional and
	// again represents further positional arguments.
	Expand = clif.Expand
)
