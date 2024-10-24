package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
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
			xi18n.Text(locale.WorkerPoolCPUParamUsageTemplData{}),
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
			xi18n.Text(locale.WorkerPoolNoWParamUsageTemplData{}),
			defaultNoW,
			flagSet...,
		),
		&parent.Native.NoWorkers,
		minNow,
		maxNow,
	)

	parent.Command.MarkFlagsMutuallyExclusive("cpu", "now")
}
