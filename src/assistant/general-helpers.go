package assistant

import (
	"fmt"

	"github.com/snivilised/cobrass"
	"github.com/spf13/pflag"
)

// GetThirdPartyCL creates a command line from the flag set. If there are any
// short form flags, they are resolved using the knownBy map, which the client
// provides, mapping long form flag names to their short form. The client can
// choose to compose a command line consisting of all available flags or just
// the ones changed by the user (ie, they are explicitly specified on the
// command line as opposed to be defaulted).
func GetThirdPartyCL(
	flagSet *pflag.FlagSet,
	knownBy cobrass.KnownByCollection,
) cobrass.ThirdPartyCommandLine {
	// move to cobrass/clif
	//
	cl := cobrass.ThirdPartyCommandLine{}

	flagSet.Visit(func(f *pflag.Flag) {
		if _, found := knownBy[f.Name]; found {
			cl = append(cl, fmt.Sprintf("--%v", f.Name))
			if f.Value.Type() != "bool" {
				cl = append(cl, f.Value.String())
			}
		}
	})

	return cl
}
