package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	//重定向到百度
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	// 跳转到 /b 对应的路由处理函数
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/d" // 把请求的URI修改
		r.HandleContext(c)        // 继续后续的处理
	})
	r.GET("/d", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})
	r.Run(":8080")
}
