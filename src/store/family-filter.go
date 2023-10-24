package store

import (
	"regexp"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/i18n"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/spf13/pflag"
)

const (
	defaultFilterValue = ""
)

// NB: We don't want to use up too many of the letters of the alphabet
// on short flag names, because it leaves less for the client to use.
// Therefore, for compound filters typically used when we want to filter
// file system nodes by file name and directory name, we forego the
// ability to specify compound file names (when using the navigator
// in extendio with the FoldersWithFiles subscription type) with a
// short code, as this is seen as a niche feature. The more common
// scenarios would be to either filter files, directories or both
// by using an 'any' scope. With this compromise, the user would
// always have to spell the compound file filter in it full form:
// --files-rx or --files-gb. When using extendio nav, the folders
// with files subscription would have to be used, ie there is no
// standalone file file, so --files-rx and --files-gb are both free
// to use without ambiguity.
// For a regular files scenario, we would need to use the files
// subscription type and in this case, --files-rx(x) and --files-gb(g)
// are still free to be used without ambiguity.

type FilesFilterParameterSet struct {
	FilesGlob  string
	FilesRexEx string
}

func (f *FilesFilterParameterSet) BindAll(
	parent *assistant.ParamSet[FilesFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --files-gb(G)
	//
	parent.BindString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FilesGlobParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FilesGlob,
	)

	// --files-rx(X)
	//
	parent.BindValidatedString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FilesRegExParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FilesRexEx,
		func(value string, _ *pflag.Flag) error {
			_, err := regexp.Compile(value)
			return err
		},
	)

	parent.Command.MarkFlagsMutuallyExclusive("files-gb", "files-rx")
}

type FoldersFilterParameterSet struct {
	FoldersGlob  string
	FoldersRexEx string
}

func (f *FoldersFilterParameterSet) BindAll(
	parent *assistant.ParamSet[FoldersFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --folders-gb(Z)
	//
	parent.BindString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FolderGlobParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FoldersGlob,
	)

	// --folders-rx(y)
	//
	parent.BindValidatedString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FolderRexExParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FoldersRexEx,
		func(value string, _ *pflag.Flag) error {
			_, err := regexp.Compile(value)
			return err
		},
	)

	parent.Command.MarkFlagsMutuallyExclusive("folders-gb", "folders-rx")
}

type PolyFilterParameterSet struct {
	FilesGlob    string
	FilesRexEx   string
	FoldersGlob  string
	FoldersRexEx string
}

func (f *PolyFilterParameterSet) BindAll(
	parent *assistant.ParamSet[PolyFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// argh, code smell here, because we're duplicating the functionality
	// in FileFilterParameterSet and FoldersFilterParameterSet, but that can't
	// be helped because of the paramSet instance is type specific.
	//
	// --files-gb(G)
	//
	parent.BindString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FilesGlobParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FilesGlob,
	)

	// --files-rx(X)
	//
	parent.BindValidatedString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FilesRegExParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FilesRexEx,
		func(value string, _ *pflag.Flag) error {
			_, err := regexp.Compile(value)
			return err
		},
	)

	// --folders-gb(Z)
	//
	parent.BindString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FolderGlobParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FoldersGlob,
	)

	// --folders-rx(y)
	//
	parent.BindValidatedString(
		resolveNewFlagInfo(
			xi18n.Text(i18n.FolderRexExParamUsageTemplData{}),
			defaultFilterValue,
			flagSet...,
		),
		&parent.Native.FoldersRexEx,
		func(value string, _ *pflag.Flag) error {
			_, err := regexp.Compile(value)
			return err
		},
	)

	parent.Command.MarkFlagsMutuallyExclusive("files-gb", "files-rx")
	parent.Command.MarkFlagsMutuallyExclusive("folders-gb", "folders-rx")
}
