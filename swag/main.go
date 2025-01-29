package main

import (
	"net/http"
	_ "page/swag/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// Request 请求参数
type Request struct {
	Code  int    `form:"code" json:"code"`   // 状态码
	Msg   string `form:"msg" json:"msg"`     // 响应信息
	Limit int    `form:"limit" json:"limit"` // 返回条数
	Data  string `form:"data" json:"data"`   // 返回数据
}

// Index 首页
// @Tags 用户管理
// @Summary 用户列表
// @Description 返回一个用户列表，可根据查询参数指定
// @Param limit query int false "返回多少条"
// @Param data body Request true "表示多个参数"
// @Router /api/index [get]
// @Produce json
// @Success 200
func index(c *gin.Context) {
	var request Request
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": request,
	})
}

// @title 调试swagger文档
// @version 1.0
// @description api接口测试文档
// @host 127.0.0.1:8080
func main() {
	r := gin.Default()
	r.GET("/api/index", index) // 确保路由路径与注释中的路径一致
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
