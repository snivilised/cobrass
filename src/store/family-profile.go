package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
)

type ProfileParameterSet struct {
	Profile string
}

func (f *ProfileParameterSet) BindAll(self *assistant.ParamSet[ProfileParameterSet]) {
	const (
		defaultProfile = ""
	)

	self.BindValidatedStringIsMatch(
		newFlagInfo(
			xi18n.Text(i18n.ProfileParamUsageTemplData{}),
			defaultProfile,
		),
		&self.Native.Profile,
		`^[\w-]+$`,
	)
}
