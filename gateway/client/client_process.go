package client

import (
	"github.com/TCP404/OneTiny-core/config"
	"github.com/TCP404/OneTiny-core/engine"
	"github.com/TCP404/OneTiny-core/gateway/vo"
	"github.com/TCP404/OneTiny-core/verify"
)

type processClient struct{}

// verifier
func (p *processClient) NewVerifyChain() vo.Verifier {
	return verify.NewVerifyChain()
}
func (p *processClient) NewPortHandler(port int) vo.Handler {
	return verify.NewPortHandler(port)
}
func (p *processClient) NewPathHandler(path string) vo.Handler {
	return verify.NewPathHandler(path)
}
func (p *processClient) NewUPSHandler(weight uint8) vo.Handler {
	return verify.NewUPSHandler(weight)
}

// kernel
func (p *processClient) New(mode vo.Mode) vo.Engine {
	return engine.New(engine.Mode(mode))
}
func (p *processClient) NewWithConfig(mode vo.Mode, c vo.Config) vo.Engine {
	return engine.NewWithConfig(engine.Mode(mode), c)
}
func (p *processClient) RunCore(ip string, port int, path string, opt ...vo.ConfigOpt) error {
	return engine.RunCore(ip, port, path, opt...)
}
func (p *processClient) Run(e vo.Engine, port int) error {
	return e.Run(port)
}
func (p *processClient) Default() vo.Engine {
	return engine.Default()
}

// configure
func (p *processClient) NewConfig(ip string, port int, path string, opt ...vo.ConfigOpt) vo.Config {
	return config.NewConfig(ip, port, path, opt...)
}
func (p *processClient) IsAllowUpload(allow bool) vo.ConfigOpt {
	return config.OptIsAllowUpload(allow)
}
func (p *processClient) RootPath(RootPath string) vo.ConfigOpt {
	return config.OptRootPath(RootPath)
}
func (p *processClient) MaxLevel(MaxLevel uint) vo.ConfigOpt {
	return config.OptMaxLevel(MaxLevel)
}
func (p *processClient) IsSecure(IsSecure bool) vo.ConfigOpt {
	return config.OptIsSecure(IsSecure)
}
