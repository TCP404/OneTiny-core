package config

import (
	"log"

	"github.com/fatih/color"
)

// PrintInfo 会在程序启动后打印本机 IP、共享目录、是否允许上传的信息
func (c *Config) PrintInfo(opt ...PrintOpt) {
	log.SetOutput(color.Output)
	for _, o := range opt {
		o(c)
	}
}

type PrintOpt func(c *Config)

func PrintAll() []PrintOpt {
	return []PrintOpt{
		PrintAddr,
		PrintPath,
		PrintLevel,
		PrintAllow,
		PrintSecure,
	}
}

func PrintAddr(c *Config) {
	if c.IP != "" {
		log.Printf("Run on   [ %s ]", color.BlueString("http://%s:%d", c.IP, c.Port))
		return
	}
	log.Printf("%s", color.YellowString("Warning: [ 暂时获取不到您的IP, 可以打开新的命令行窗口输入 ->  ipconfig , 查看您的IP。]"))
}

// Print RootPath infomation
func PrintPath(c *Config) {
	log.Printf("Run with [ %s ]", color.BlueString("%s", c.RootPath))
}

// Print Max allow access level
func PrintLevel(c *Config) {
	log.Printf("Allow access level: [ %s ]", color.BlueString("%d", c.MaxLevel))
}

// Print Allow upload Status
func PrintAllow(c *Config) {
	status := color.RedString("%t", c.IsAllowUpload)
	if c.IsAllowUpload {
		status = color.GreenString("%t", c.IsAllowUpload)
	}
	log.Printf("Allow upload: [ %s ]", status)
}

// Print Secure status
func PrintSecure(c *Config) {
	status := color.RedString("%t", c.IsSecure)
	if c.IsSecure {
		status = color.GreenString("%t", c.IsSecure)
	}
	log.Printf("Need Login: [ %s ]\n\n", status)
}
