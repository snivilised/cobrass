package locale

// FilesGlobParamUsageTemplData
// 🧊
type FilesGlobParamUsageTemplData struct {
	CobrassTemplData
}

func (td FilesGlobParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "files-glob-filter.param-usage",
		Description: "files glob filter (negate-able with leading !)",
		Other:       "files-gb files glob filter (negate-able with leading !)",
	}
}

// FilesRegExParamUsageTemplData
// 🧊
type FilesRegExParamUsageTemplData struct {
	CobrassTemplData
}

func (td FilesRegExParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "files-regex-filter.param-usage",
		Description: "files regex filter (negate-able with leading !)",
		Other:       "files-rx folder regular expression filter (negate-able with leading !)",
	}
}

// FilesExGlobParamUsageTemplData
// 🧊
type FilesExGlobParamUsageTemplData struct {
	CobrassTemplData
}

func (td FilesExGlobParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "files-ex-glob-filter.param-usage",
		Description: "files extended glob filter (negate-able with leading !)",
		Other:       "files extended glob filter: <glob>|<suffixes csv> (negate-able with leading !)",
	}
}

// FoldersExGlobParamUsageTemplData
// 🧊
type FoldersExGlobParamUsageTemplData struct {
	CobrassTemplData
}

func (td FoldersExGlobParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "folders-ex-glob-filter.param-usage",
		Description: "folders extended glob filter (negate-able with leading !)",
		Other:       "folders extended glob filter: <glob> (negate-able with leading !)",
	}
}

// FoldersGlobParamUsageTemplData
// 🧊
type FoldersGlobParamUsageTemplData struct {
	CobrassTemplData
}

func (td FoldersGlobParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "folders-glob-filter.param-usage",
		Description: "folders glob (negate-able with leading !)",
		Other:       "folders-gb folder glob filter (negate-able with leading !)",
	}
}

// FoldersRexExParamUsageTemplData
// 🧊
type FoldersRexExParamUsageTemplData struct {
	CobrassTemplData
}

func (td FoldersRexExParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "folders-regex-filter.param-usage",
		Description: "folders regex filter (negate-able with leading !)",
		Other:       "folders-rx folder regular expression filter (negate-able with leading !)",
	}
}

// WorkerPoolCPUParamUsageTemplData
// 🧊
type WorkerPoolCPUParamUsageTemplData struct {
	CobrassTemplData
}

func (td WorkerPoolCPUParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "worker-pool-cpu.param-usage",
		Description: "run with the number of workers in pool set to number of CPUs available",
		Other:       "cpu denotes parallel execution with all available processors",
	}
}

// WorkerPoolCPUParamUsageTemplData
// 🧊
type WorkerPoolNoWParamUsageTemplData struct {
	CobrassTemplData
}

func (td WorkerPoolNoWParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "worker-pool-cpu.param-usage",
		Description: "run with the number of workers in pool set to this number",
		Other:       "now denotes parallel execution with this number of workers in pool",
	}
}

// ProfileParamUsageTemplData
// 🧊
type ProfileParamUsageTemplData struct {
	CobrassTemplData
}

func (td ProfileParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "profile.param-usage",
		Description: "pre-defined flag/option list in config file",
		Other:       "profile specifies which set of flags/options to load from config",
	}
}

// SchemeParamUsageTemplData
// 🧊
type SchemeParamUsageTemplData struct {
	CobrassTemplData
}

func (td SchemeParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "scheme.param-usage",
		Description: "scheme is a collection of profiles, typically to repeat an operation over",
		Other:       "scheme is a collection of profiles, typically to repeat an operation over",
	}
}

// DryRunParamUsageTemplData
// 🧊
type DryRunParamUsageTemplData struct {
	CobrassTemplData
}

func (td DryRunParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "dry-run.param-usage",
		Description: "allows the user to preview the effects of a command without running it",
		Other:       "dry-run allows the user to see the effects of a command without running it",
	}
}

// LanguageParamUsageTemplData
// 🧊
type LanguageParamUsageTemplData struct {
	CobrassTemplData
}

func (td LanguageParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "language.param-usage",
		Description: "language allows the user to override the language the app runs in",
		Other:       "language allows the user to override the language the app runs in",
	}
}

// CascadeDepthParamUsageTemplData
// 🧊
type CascadeDepthParamUsageTemplData struct {
	CobrassTemplData
}

func (td CascadeDepthParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "cascade-depth.param-usage",
		Description: "limits the number of sub directories navigated",
		Other:       "depth denotes the number of sub directories to navigate",
	}
}

// CascadeNoRecurseParamUsageTemplData
// 🧊
type CascadeNoRecurseParamUsageTemplData struct {
	CobrassTemplData
}

func (td CascadeNoRecurseParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "cascade-no-recurse.param-usage",
		Description: "sets the navigator to not descend into sub directories",
		Other:       "no-recurse sets the navigator to not descend into sub directories",
	}
}

// SamplingSampleUsageTemplData
// 🧊
type SamplingSampleUsageTemplData struct {
	CobrassTemplData
}

func (td SamplingSampleUsageTemplData) Message() *Message {
	return &Message{
		ID:          "sampling-sample.param-usage",
		Description: "sampling sample usage; activates sampling",
		Other:       "sample is a flag that activates sampling",
	}
}

// SamplingNoFilesUsageTemplData
// 🧊
type SamplingNoFilesUsageTemplData struct {
	CobrassTemplData
}

func (td SamplingNoFilesUsageTemplData) Message() *Message {
	return &Message{
		ID:          "sampling-no-files.param-usage",
		Description: "sampling files usage; no of files in sample set",
		Other:       "no-files specifies the number of files to sample",
	}
}

// SamplingNoFoldersUsageTemplData
// 🧊
type SamplingNoFoldersUsageTemplData struct {
	CobrassTemplData
}

func (td SamplingNoFoldersUsageTemplData) Message() *Message {
	return &Message{
		ID:          "sampling-no-folders.param-usage",
		Description: "sampling folders usage; no of folders in sample set",
		Other:       "no-folders specifies the number of folders to sample",
	}
}

// SamplingLastUsageTemplData
// 🧊
type SamplingLastUsageTemplData struct {
	CobrassTemplData
}

func (td SamplingLastUsageTemplData) Message() *Message {
	return &Message{
		ID:          "sampling-last.param-usage",
		Description: "sampling last usage; indicates which n items are to be sampled",
		Other:       "last is a flag that indicates last n items are to be sampled instead of the first",
	}
}

// TextualInteractionIsNoTUIUsageTemplData
// 🧊
type TextualInteractionIsNoTUIUsageTemplData struct {
	CobrassTemplData
}

func (td TextualInteractionIsNoTUIUsageTemplData) Message() *Message {
	return &Message{
		ID:          "textual-interaction-is-no-tui.param-usage",
		Description: "textual interaction is no-tui usage; deactivates tui mode",
		Other:       "no-tui is a flag that turns off tui mode",
	}
}

// CliInteractionIsTUIUsageTemplData
// 🧊
type CliInteractionIsTUIUsageTemplData struct {
	CobrassTemplData
}

func (td CliInteractionIsTUIUsageTemplData) Message() *Message {
	return &Message{
		ID:          "cli-interaction-is-tui.param-usage",
		Description: "tui interaction is tui usage; activates tui mode",
		Other:       "tui is a flag that enables tui mode",
	}
}
