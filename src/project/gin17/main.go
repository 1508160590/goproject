package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	//Any:请求方法的大集合
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"msg": "POST"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"msg": "Put"})
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"msg": "GET"})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{"msg": "DELETE"})
		}
	})
	//NoRoute:返回未定义路由的页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404未定义",
		})
	})
	//路由组,Group().用于区分不同的业务线
	//把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/aa")
		videoGroup.GET("/bb")
		videoGroup.GET("/cc")
		videoGroup.GET("/dd")
		//嵌套路由组
		shopGroup := videoGroup.Group("/shop")
		xx := shopGroup.Group("xx")
		xx.GET("/ee", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "GET /video/xx/ee"})
		})
	}

	r.Run(":8080")
}
