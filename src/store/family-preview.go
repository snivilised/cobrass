package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
)

type PreviewParameterSet struct {
	DryRun bool
}

func (f *PreviewParameterSet) BindAll(self *assistant.ParamSet[PreviewParameterSet]) {
	// --dry-run(D)
	//
	const (
		defaultDryRun = false
	)

	self.BindBool(
		newFlagInfo(
			xi18n.Text(i18n.DryRunParamUsageTemplData{}),
			defaultDryRun,
		),
		&self.Native.DryRun,
	)
}
