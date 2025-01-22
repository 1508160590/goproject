package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello Golang",
	})
}

func main() {
	//创建一个路由引擎
	r := gin.Default()

	//第一种使用匿名函数定义
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{ //返回json格式数据
			"message": "hello GET",
		})
	})
	//第二种单独定义
	r.GET("/gets", sayHello)

	r.POST("/posts", func(c *gin.Context) {
		c.JSON(200, gin.H{ //返回json格式数据
			"message": "hello POST",
		})
	})
	r.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{ //返回json格式数据
			"message": "hello PUT",
		})
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{ //返回json格式数据
			"message": "hello DELETE",
		})
	})
	r.Run() //启动服务
}
