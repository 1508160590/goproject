package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username :=c.DefaultPostForm("username","小王子")

		//GetPostForm取不到返回bool类型值
		//username,ok := c.GetPostForm(username)
		//if !ok {
		//	username ="sb"
		//}
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})
	r.Run(":8080")
}
