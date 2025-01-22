package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//通过Query获取请求中携带的usernamestring参数
		//username :=c.Query("username")
		address := c.Query("address")

		//或者query字段的参数，没有query字段，默认值为somibody
		//name := c.DefaultOuery("query","somibody")

		//name, ok :=c.GetQuery("query")// 取不到第二个参数就返回false
		//if !ok {
		// 取不到
		//name = "somebody"
		//}

		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
