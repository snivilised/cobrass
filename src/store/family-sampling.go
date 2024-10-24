package store

import (
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/locale"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

type SamplingParameterSet struct {
	IsSampling bool
	NoFiles    uint
	NoFolders  uint
	Last       bool
}

func (f *SamplingParameterSet) BindAll(
	parent *assistant.ParamSet[SamplingParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	const (
		defIsSampling = false
	)

	// --sample
	//
	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(locale.SamplingSampleUsageTemplData{}),
			defIsSampling,
			flagSet...,
		),
		&parent.Native.IsSampling,
	)

	const (
		defFSItems = uint(3)
		minFSItems = uint(1)
		maxFSItems = uint(128)
	)

	// --no-files
	//
	parent.BindValidatedUintWithin(
		resolveNewFlagInfo(
			xi18n.Text(locale.SamplingNoFilesUsageTemplData{}),
			defFSItems,
			flagSet...,
		),
		&parent.Native.NoFiles,
		minFSItems,
		maxFSItems,
	)

	// --no-folders
	//
	parent.BindValidatedUintWithin(
		resolveNewFlagInfo(
			xi18n.Text(locale.SamplingNoFoldersUsageTemplData{}),
			defFSItems,
			flagSet...,
		),
		&parent.Native.NoFolders,
		minFSItems,
		maxFSItems,
	)

	const (
		defIsLast = false
	)

	// --last
	//
	parent.BindBool(
		resolveNewFlagInfo(
			xi18n.Text(locale.SamplingLastUsageTemplData{}),
			defIsLast,
			flagSet...,
		),
		&parent.Native.Last,
	)
}
