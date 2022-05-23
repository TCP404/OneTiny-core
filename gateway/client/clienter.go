package client

import (
	"github.com/TCP404/OneTiny-core/gateway/vo"
)

type Client interface {
	verifier
	kernel
	configure
}
type verifier interface {
	NewVerifyChain() vo.Verifier
	NewPortHandler(port int) vo.Handler
	NewPathHandler(path string) vo.Handler
	NewUPSHandler(weight uint8) vo.Handler
}

type kernel interface {
	New(mode vo.Mode) vo.Engine
	NewWithConfig(mode vo.Mode, c vo.Config) vo.Engine
	Run(e vo.Engine, port int) error
	RunCore(ip string, port int, path string, opt ...vo.ConfigOpt) error
}

type configure interface {
	NewConfig(ip string, port int, path string, opt ...vo.ConfigOpt) vo.Config
	IsAllowUpload(allow bool) vo.ConfigOpt
	RootPath(path string) vo.ConfigOpt
	MaxLevel(level uint) vo.ConfigOpt
	IsSecure(secure bool) vo.ConfigOpt
}

var _ Client = (*processClient)(nil)
var ProcessClient = newProcessClient()

func newProcessClient() Client { return &processClient{} }
