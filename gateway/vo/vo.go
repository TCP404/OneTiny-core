package vo

import (
	"github.com/TCP404/OneTiny-core/config"
	"github.com/TCP404/OneTiny-core/engine"
	"github.com/TCP404/OneTiny-core/verify"
)

type (
	Handler   = verify.Handler
	Verifier  = verify.Verifier
	Engine    = *engine.Engine
	Config    = *config.Config
	ConfigOpt = config.Opt
	Mode      = engine.Mode
)

const (
	Debug   Mode = engine.Debug
	Release Mode = engine.Release
	Test    Mode = engine.Test
)

var (
	PAddr   = config.PrintAddr
	PPath   = config.PrintPath
	PAllow  = config.PrintAllow
	PLevel  = config.PrintLevel
	PSecure = config.PrintSecure

	RootPath      = config.OptRootPath
	IsAllowUpload = config.OptIsAllowUpload
	MaxLevel      = config.OptMaxLevel
	IsSecure      = config.OptIsSecure
	Output        = config.OptOutput
)
