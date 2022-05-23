package engine

import (
	"errors"
	"net"
	"os"
	"strconv"

	// "github.com/TCP404/OneTiny-cli/core/routes"
	"github.com/TCP404/OneTiny-core/config"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*config.Config
}

type Context struct {
	gin.Context
}

type Mode string

const (
	Debug   Mode = "debug"
	Release Mode = "release"
	Test    Mode = "test"
)

func New(mode Mode) *Engine {
	gin.SetMode(string(mode))
	e := &Engine{Engine: gin.New()}
	return e
}

func Default() *Engine {
	e := New(Debug)
	c := config.NewConfig("127.0.0.1", 9090, "/", config.OptOutput(os.Stderr))
	e.Config = c
	attachToConfig(c)

	return e
}

func NewWithConfig(mode Mode, c *config.Config) *Engine {
	e := New(mode)
	e.Config = c
	attachToConfig(c)
	return e
}

func attachToConfig(c *config.Config) {
	config.IP = c.IP
	config.Port = c.Port
	config.SessionVal = c.SessionVal
	config.RootPath = c.RootPath
	config.IsAllowUpload = c.IsAllowUpload
	config.MaxLevel = c.MaxLevel
	config.IsSecure = c.IsSecure
	config.Output = c.Output
}

// RunCore 函数负责启动 gin 实例，开始提供 HTTP 服务
func RunCore(ip string, port int, path string, opt ...config.Opt) error {
	e := NewWithConfig(Release, config.NewConfig(ip, port, path, opt...))
	e.PrintInfo(config.PrintAll()...)
	return e.Run(port)
}

func (e *Engine) SetConfig(ip string, port int, path string, opt ...config.Opt) *Engine {
	c := config.NewConfig(ip, port, path, opt...)
	e.Config = c
	attachToConfig(c)
	return e
}

func (e *Engine) PrintInfo(opt ...config.PrintOpt) *Engine {
	if e.Config == nil {
		return e
	}
	e.Config.PrintInfo(opt...)
	return e
}

func (e *Engine) Run(ports ...int) error {
	e.SetRoute()
	e.SetMiddleware()
	var port = 9090
	if len(ports) >= 1 {
		port = ports[0]
	} else if e.Config != nil {
		port = e.Config.Port
	}

	err := e.Engine.Run(":" + strconv.Itoa(port))
	switch err.(type) {
	case *net.OpError:
		return errors.New(color.RedString("指定的 %d 端口已被占用，请换一个端口号。", port))
	}
	return nil
}

func (e *Engine) SetRoute() *Engine {
	loadIndexRoute(e)
	loadCoreRoute(e)
	loadLoginRoute(e)
	load404Route(e)
	loadICORoute(e)
	return e
}

func (e *Engine) SetMiddleware() *Engine {
	e.Use(gin.LoggerWithWriter(e.Config.Output), gin.Recovery())
	e.Use(e.enableCookieSession())
	return e
}
