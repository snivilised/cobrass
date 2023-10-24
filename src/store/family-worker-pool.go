package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type WorkerPoolParameterSet struct {
	CPU       bool
	NoWorkers int
}

func (f *WorkerPoolParameterSet) BindAll(
	parent *assistant.ParamSet[WorkerPoolParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --cpu(C)
	//
	const (
		defaultCPU = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(i18n.WorkerPoolCPUParamUsageTemplData{}),
			defaultCPU,
			flagSet...,
		),
		&parent.Native.CPU,
	)

	// --now(N)
	//
	const (
		defaultNoW = -1
		minNow     = -1
		maxNow     = 100
	)

	parent.BindValidatedIntWithin(
		resolveNewFlagInfo(
			xi18n.Text(i18n.WorkerPoolNoWParamUsageTemplData{}),
			defaultNoW,
			flagSet...,
		),
		&parent.Native.NoWorkers,
		minNow,
		maxNow,
	)

	parent.Command.MarkFlagsMutuallyExclusive("cpu", "now")
}
