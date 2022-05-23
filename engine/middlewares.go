package engine

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/TCP404/OneTiny-core/define"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InterceptICO 拦截浏览器默认请求 favicon.ico 的行为
func (e *Engine) InterceptICO() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Param("filename"), ".ico") {
			c.Status(http.StatusOK)
			c.Abort()
		}
	}
}

func (e *Engine) CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查 session，
		// 有则放行
		// 无则跳转登录页面
		if !e.Config.IsSecure {
			return
		}

		if session := sessions.Default(c); session.Get("login") == e.Config.SessionVal {
			c.Next()
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func (e *Engine) enableCookieSession() gin.HandlerFunc {
	s := cookie.NewStore([]byte("secret"))
	return sessions.Sessions("SESSIONID", s)
}

// CheckLevel 负责检查当前访问层级是否超出设定最大层级
// 例如：
// 		共享目录为 /a/b/ , 最大层级为 2
//		✓: /a/b/			根目录
//		✓: /a/b/file	    根目录下文件
// 		✓: /a/b/c/			根目录下第一层目录
// 		✓: /a/b/c/file		根目录下第一层目录下的文件
//		✓: /a/b/c/d/		根目录下第二层目录
// 		✓: /a/b/c/d/file	根目录下第二层目录下的文件
//		x: /a/b/c/d/e/		根目录下第三层目录
// 		x: /a/b/c/d/e/file	根目录下第三层目录下的文件
func (e *Engine) CheckLevel() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := strings.TrimPrefix(c.Param("filename"), "/file")

		c.Set("filename", filePath)

		isD := e.isDir(filePath)
		c.Set("isDirectory", isD)
		isFile := !isD
		if e.isOverLevel(filePath, isFile) {
			c.String(http.StatusNotFound, "访问超出允许范围，请返回！")
			c.Abort()
		}
	}
}

// 判断是否是目录
func (e *Engine) isDir(filePath string) bool {
	if filePath == define.ROOT {
		return true
	}
	finfo, _ := os.Stat(filepath.Join(e.Config.RootPath, filePath))
	return finfo.IsDir()
}

// 检查当前访问的路径是否超过限定层级
func (e *Engine) isOverLevel(relPath string, isFile bool) bool {
	rel, _ := filepath.Rel(e.Config.RootPath, filepath.Join(e.Config.RootPath, relPath))
	i := strings.Split(rel, define.SEPARATORS)
	level := len(i)
	if i[0] == "." {
		level = 0
	}
	if isFile {
		level--
	}
	return level > int(e.Config.MaxLevel)
}
