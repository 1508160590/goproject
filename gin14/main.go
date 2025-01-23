package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./file.html")
	r.GET("/file", func(c *gin.Context) {
		c.HTML(http.StatusOK, "file.html", nil)
	})
	//处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	//router.MaxMultipartMemory=8<<20 //8 MiB
	r.POST("/file", func(c *gin.Context) {
		//从请求中读取文件
		file, err := c.FormFile("f1") //从请求中获取携带的参数一样的
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})
		} else {
			//file.Size获取上传的文件大小
			//file.Header.Get("Content-Type")获取上传的文件类型
			//file.Filename获取上传的文件名
			//./将上传的文件保存到本地
			dst := fmt.Sprintf("./%s", file.Filename)
			//将读取的文件保存在本地
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		}
	})
	/*
		r.POST("/file",func(c *gin.Context){
			//从请求中读取文件
			file,_ := c.MultipartForm()
			files := form.File["f1"]

			for index,file := range files{
				log.Println(file.Filename)
				dst := fmt.Sprintf("xxx/",file.Filename,index)
				c.SaveUploadedFile(file,dst)
			}
			c.JSON(http.StatusOK,gin.H{
				"msg": fmt.Sprintf("%d file uploaded!", len(files)),
			})
		})
	*/
	r.Run()
}
