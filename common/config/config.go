package config

type Config struct {
	baseDir            string
	ServerUrl          string
	DebuggerHtmlReport string
}

func (c *Config) SetBaseDir(d string) {
	c.baseDir = d
}

func (c *Config) DebuggerHtmlReportFullPath() string {
	return c.baseDir + c.DebuggerHtmlReport
}
