package clif

import (
	"github.com/snivilised/cobrass/src/collections"
)

type (
	// ThirdPartyFlagName raw name of a flag, ie without the leading --/-
	ThirdPartyFlagName = string

	// ThirdPartyPositional represents a positional argument.
	ThirdPartyPositional = string

	// ThirdPartyPositionalArgs represents a collection of third party
	// positional arguments.
	ThirdPartyPositionalArgs = []string

	// ThirdPartyOptionValue the string value of an option. Since this option
	// is being delegated to a third party command, it does not have to be
	// of a particular native go type and can be composed from a go type
	// using the value's String() method.
	ThirdPartyOptionValue = string

	// SpecifiedFlagsCollection represents the set of third party flags
	// presented by the user on the command line.
	// (NB: Cobra does not currently have a mechanism to collect third
	// party flags, by convention, anything that follows " -- "), therefore
	// we need to collect and handle these flags/options explicitly,
	// which is less than ideal.
	// A difference between SpecifiedFlagsCollection and ThirdPartyCommandLine
	// is that switch flags have a true/false option value in SpecifiedFlagsCollection
	// but not in ThirdPartyCommandLine.
	SpecifiedFlagsCollection = collections.OrderedKeysMap[ThirdPartyFlagName, ThirdPartyOptionValue]

	// ThirdPartyPresentFlags (see SpecifiedFlagsCollection).
	ThirdPartyPresentFlags SpecifiedFlagsCollection

	// KnownByCollection collection maps a full flag name to the
	// short name it is also known by. If a flag does not
	// have a short name, it should be mapped to the empty
	// string.
	KnownByCollection map[ThirdPartyFlagName]ThirdPartyFlagName

	// ThirdPartyFlagsKnownBy (see KnownByCollection).
	ThirdPartyFlagsKnownBy KnownByCollection

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
	ThirdPartyCommandLine []string

	// ExternalThirdParty base struct for cli applications using the
	// entry paradigm that need to delegate an invocation to an
	// external third party command.
	ExternalThirdParty struct {
		// KnownBy represents the collection of all possible flags that
		// can be specified in a particular invocation (see KnownByCollection)
		KnownBy KnownByCollection

		// ExecutionFlags represents a particular invocation of a third party
		// command that contains flags explicitly specified and ones loaded
		// from other sources, such as config. (see ThirdPartyCommandLine).
		ExecutionFlags ThirdPartyCommandLine
	}
)
