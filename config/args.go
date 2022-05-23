package config

import "io"

var (
	SessionVal    string
	IP            string
	Port          int
	RootPath      string
	IsAllowUpload bool
	MaxLevel      uint
	IsSecure      bool
	Output        io.Writer
)
