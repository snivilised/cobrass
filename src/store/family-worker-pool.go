package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
)

type WorkerPoolParameterSet struct {
	CPU       bool
	NoWorkers int
}

func (f *WorkerPoolParameterSet) BindAll(self *assistant.ParamSet[WorkerPoolParameterSet]) {
	// --cpu(C)
	//
	const (
		defaultCPU = false
	)

	self.BindBool(
		newFlagInfo(
			xi18n.Text(i18n.WorkerPoolCPUParamUsageTemplData{}),
			defaultCPU,
		),
		&self.Native.CPU,
	)

	// --now(N)
	//
	const (
		defaultNoW = -1
		minNow     = -1
		maxNow     = 100
	)

	self.BindValidatedIntWithin(
		newFlagInfo(
			xi18n.Text(i18n.WorkerPoolNoWParamUsageTemplData{}),
			defaultNoW,
		),
		&self.Native.NoWorkers,
		minNow,
		maxNow,
	)

	self.Command.MarkFlagsMutuallyExclusive("cpu", "now")
}
