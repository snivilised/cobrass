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

type FilterParameterSet struct {
	FilesGlob    string
	FilesRexEx   string
	FoldersGlob  string
	FoldersRexEx string
}

func (f *FilterParameterSet) BindAll(self *assistant.ParamSet[FilterParameterSet]) {
	// --files-gb(G)
	//
	self.BindString(
		newFlagInfo(
			xi18n.Text(i18n.FilesGlobParamUsageTemplData{}),
			defaultFilterValue,
		),
		&self.Native.FilesGlob,
	)

	// --files-rx(X)
	//
	self.BindValidatedString(
		newFlagInfo(
			xi18n.Text(i18n.FilesRegExParamUsageTemplData{}),
			defaultFilterValue),
		&self.Native.FilesRexEx,
		func(value string, _ *pflag.Flag) error {
			_, err := regexp.Compile(value)
			return err
		},
	)

	// --folders-gb(z)
	//
	self.BindString(
		newFlagInfo(
			xi18n.Text(i18n.FolderGlobParamUsageTemplData{}),
			defaultFilterValue,
		),
		&self.Native.FoldersGlob,
	)

	// --folders-rx(y)
	//
	self.BindValidatedString(
		newFlagInfo(
			xi18n.Text(i18n.FolderRexExParamUsageTemplData{}),
			defaultFilterValue,
		),
		&self.Native.FoldersRexEx,
		func(value string, _ *pflag.Flag) error {
			_, err := regexp.Compile(value)
			return err
		},
	)

	self.Command.MarkFlagsMutuallyExclusive("files-gb", "files-rx")
	self.Command.MarkFlagsMutuallyExclusive("folders-gb", "folders-rx")
}
