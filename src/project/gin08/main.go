package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//// 加载静态文件
	r.Static("/xxx", "./statics") //找到xxx下面的statics下的文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/product", func(c *gin.Context) {
		c.HTML(http.StatusOK, "product.html", nil)
	})
	r.GET("/category", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category.html", nil)
	})
	r.Run(":9090") //默认启动8080端口
}
