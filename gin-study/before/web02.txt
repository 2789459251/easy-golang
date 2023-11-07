package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayhello(c *gin.Context) {
	c.JSONP(200, gin.H{
		"message": "hello golang",
	})
}
func main() {
	//创建默认路由引擎
	r := gin.Default()
	//指定地址和函数->用户使用get访问

	r.GET("/hello", sayhello) //http请求方法

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "get",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "post",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "put",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "delete",
		})
	})
	//启动服务
	r.Run(":9090")
}
