package i18n

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

// FolderGlobParamUsageTemplData
// 🧊
type FolderGlobParamUsageTemplData struct {
	CobrassTemplData
}

func (td FolderGlobParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "folders-glob-filter.param-usage",
		Description: "folders glob (negate-able with leading !)",
		Other:       "folders-gb folder glob filter (negate-able with leading !)",
	}
}

// FolderRexExParamUsageTemplData
// 🧊
type FolderRexExParamUsageTemplData struct {
	CobrassTemplData
}

func (td FolderRexExParamUsageTemplData) Message() *Message {
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

// CascadeSkimParamUsageTemplData
// 🧊
type CascadeSkimParamUsageTemplData struct {
	CobrassTemplData
}

func (td CascadeSkimParamUsageTemplData) Message() *Message {
	return &Message{
		ID:          "cascade-skim.param-usage",
		Description: "sets the navigator to not descend into sub directories",
		Other:       "skim sets the navigator to not descend into sub directories",
	}
}
