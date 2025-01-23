// Gin中间件
// Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子(Hook)函数。
// 这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Gin中的中间件必须是一个 gin.HandlerFunc 类型
func m1(c *gin.Context) {
	fmt.Println("m1")
	start := time.Now()
	//go funcxx() //在funcxx中只能使用c的拷贝  c.copy()
	c.Next() //调用后续的处理函数
	//c.Abort()//阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost:", cost)
}

func m2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	}
}

func main() {
	r := gin.Default() // 默认使用了Logger()和Recovery()两个中间件
	//Logger中间件将日志写入 gin.DefaultWriter，即使配置了 GIN MODE=release.
	//Recovery中间件会recover任何 panic。如果有panic的话，会写入500应码。
	r.GET("/index", m1, m2())

	//路由组注册中间件方法1:
	xxGroup := r.Group("/xx", m1, m2())
	xxGroup.GET("/index1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
	})
	//路由组注册中间件方法2:
	xx2Group := r.Group("/xx2")
	xx2Group.Use(m2())
	xx2Group.GET("/index2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
	})

	r.Run()
}
