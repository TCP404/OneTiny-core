package engine

import (
	"github.com/TCP404/OneTiny-core/api"
	"github.com/gin-gonic/gin"
)

func loadLoginRoute(r *Engine) {
	r.GET("/login", api.LoginGet)
	r.POST("/login", api.LoginPost)
}

func loadCoreRoute(r *Engine) *gin.RouterGroup {
	fileG := r.Group("/file", r.CheckLogin(), r.CheckLevel())
	{
		fileG.GET("/*filename", api.Downloader)
		fileG.POST("/upload", api.Uploader)
	}
	return fileG
}

func load404Route(app *Engine) {
	app.NoRoute(api.NotFound)
}

func loadIndexRoute(app *Engine) {
	app.GET("/", api.Index)
}

func loadICORoute(app *Engine) {
	app.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(200)
		c.Writer.Write([]byte(`<svg t="1622861750046" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5806" width="108" height="108"><path d="M277.76 322.56h468.48v65.28H277.76zM277.76 570.24h468.48v65.28H277.76zM277.76 817.28h468.48v78.08H277.76z" fill="#69C6DC" p-id="5807"></path><path d="M264.96 323.2h26.24v65.28h-26.24zM733.44 323.2h26.24v65.28h-26.24zM264.96 570.88h26.24v65.28h-26.24zM733.44 570.88h26.24v65.28h-26.24zM264.96 817.92h26.24V896h-26.24zM733.44 817.92h26.24V896h-26.24zM343.04 869.76h338.56V896H343.04z" fill="#34303D" p-id="5808"></path><path d="M811.52 336H212.48V128h599.04z m-572.8-25.6h546.56V154.24H238.72z" fill="#34303D" p-id="5809"></path><path d="M720 232.32m-39.04 0a39.04 39.04 0 1 0 78.08 0 39.04 39.04 0 1 0-78.08 0Z" fill="#EA5E5A" p-id="5810"></path><path d="M707.2 218.88h26.24v26.24h-26.24zM629.12 193.28h26.24v26.24h-26.24zM577.28 193.28h26.24v26.24h-26.24zM524.8 193.28h26.24v26.24H524.8zM629.12 245.12h26.24v26.24h-26.24zM577.28 245.12h26.24v26.24h-26.24zM524.8 245.12h26.24v26.24H524.8zM316.8 193.28h26.24v78.08h-26.24zM264.96 193.28h26.24v78.08h-26.24zM368.64 193.28h26.24v78.08h-26.24zM811.52 583.68H212.48V375.04h599.04z m-572.8-26.24h546.56V401.28H238.72z" fill="#34303D" p-id="5811"></path><path d="M720 479.36m-39.04 0a39.04 39.04 0 1 0 78.08 0 39.04 39.04 0 1 0-78.08 0Z" fill="#EA5E5A" p-id="5812"></path><path d="M707.2 465.92h26.24v26.24h-26.24zM629.12 440.32h26.24v26.24h-26.24zM577.28 440.32h26.24v26.24h-26.24zM524.8 440.32h26.24v26.24H524.8zM629.12 492.16h26.24v26.24h-26.24zM577.28 492.16h26.24v26.24h-26.24zM524.8 492.16h26.24v26.24H524.8zM316.8 440.32h26.24V518.4h-26.24zM264.96 440.32h26.24V518.4h-26.24zM368.64 440.32h26.24V518.4h-26.24zM811.52 832H212.48V622.72h599.04z m-572.8-26.24h546.56v-156.8H238.72z" fill="#34303D" p-id="5813"></path><path d="M720 727.04m-39.04 0a39.04 39.04 0 1 0 78.08 0 39.04 39.04 0 1 0-78.08 0Z" fill="#EA5E5A" p-id="5814"></path><path d="M707.2 713.6h26.24v26.24h-26.24zM629.12 688h26.24v26.24h-26.24zM577.28 688h26.24v26.24h-26.24zM524.8 688h26.24v26.24H524.8zM629.12 739.84h26.24v26.24h-26.24zM577.28 739.84h26.24v26.24h-26.24zM524.8 739.84h26.24v26.24H524.8zM316.8 688h26.24v78.08h-26.24zM264.96 688h26.24v78.08h-26.24zM368.64 688h26.24v78.08h-26.24z" fill="#34303D" p-id="5815"></path></svg>`))
	})
}
