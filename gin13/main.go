package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/ff", func(c *gin.Context) {
		var u UserInfo
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})
		} else {
			fmt.Printf("%v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		}
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		var u UserInfo
		u.Username = c.PostForm("username")
		u.Password = c.PostForm("password")
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"message":  "ok",
				"username": u.Username,
				"password": u.Password,
			})
		}
	})
	r.Run(":8080")
}
