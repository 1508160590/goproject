package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//// 加载静态文件
	// r.Static("/xxx", "./statics") //找到xxx下面的statics下的文件
	// r.LoadHTMLGlob("templates/*")

	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 加载 HTML 模板文件
	r.LoadHTMLFiles("./posts/f1.tmpl", "./users/f2.tmpl")

	r.GET("/posts/f", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/f1.tmpl", gin.H{
			"title": "baidu.com",
		})
	})
	r.GET("/users/f", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/f2.tmpl", gin.H{
			"title": "<a href='http://www.baidu.com'>百度</a>",
		})
	})
	r.Run(":9090") //默认启动8080端口
}
