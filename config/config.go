package config

import "io"

type Config struct {
	SessionVal    string
	IP            string
	Port          int
	RootPath      string
	IsAllowUpload bool
	MaxLevel      uint
	IsSecure      bool
	Output        io.Writer
}

func NewConfig(ip string, port int, path string, opt ...Opt) *Config {
	c := &Config{IP: ip, Port: port, RootPath: path}
	for _, o := range opt {
		o(c)
	}
	return c
}

type Opt func(c *Config)

func OptIsAllowUpload(IsAllowUpload bool) Opt {
	return func(c *Config) { c.IsAllowUpload = IsAllowUpload }
}
func OptRootPath(rootPath string) Opt { return func(c *Config) { c.RootPath = rootPath } }
func OptMaxLevel(maxLevel uint) Opt   { return func(c *Config) { c.MaxLevel = maxLevel } }
func OptIsSecure(isSecure bool) Opt   { return func(c *Config) { c.IsSecure = isSecure } }
func OptOutput(output io.Writer) Opt  { return func(c *Config) { c.Output = output } }
