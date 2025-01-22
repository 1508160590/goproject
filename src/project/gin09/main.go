package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H 是màp[string]interface{}的缩写
	r.GET("/someJS0N", func(c *gin.Context) {
		//方式一:自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	//方法二:使用map
	r.GET("/moreJSON", func(c *gin.Context) {
		msg := map[string]interface{}{
			"name":    "猪猪侠",
			"message": "hello,feifei",
			"age":     18,
		}
		c.JSON(http.StatusOK, msg)
	})

	type msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}
	r.GET("/soreJSON", func(c *gin.Context) {
		//方法三:使用结构体
		data := msg{
			Name:    "猪猪侠",
			Message: "hello,feifei",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")
}
