package verify

import (
	"errors"
	"os"

	"github.com/fatih/color"
)

type pathHandler struct {
	rootPath string
}

// pathHandler.Handle 负责校验用户指定的共享目录的绝对路径
func (p *pathHandler) Handle() error {
	if _, err := os.Stat(p.rootPath); err != nil {
		if !os.IsExist(err) {
			return errors.New(color.RedString("无法设置您指定的共享路径, 请检查给出的路径是否有问题：%s", p.rootPath))
		}
	}
	return nil
}
