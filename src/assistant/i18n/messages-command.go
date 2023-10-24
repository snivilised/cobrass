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

// WorkerPoolCPUParamUsageTemplData
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

// WorkerPoolCPUParamUsageTemplData
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