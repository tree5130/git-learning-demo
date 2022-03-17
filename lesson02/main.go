package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.tmpl") //模板解析
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", gin.H{ // 模板渲染
			"title": "XXX.com",
		})
	})
	r.Run(":8088") //启动server
}
