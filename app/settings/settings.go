package settings

type Settings struct {
	DebugMode bool

	ThreadCount int

	ConfigPath string
	ReportPath string

	Stage string
}
