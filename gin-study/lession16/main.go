package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://liwenzhou.com/posts/Go/gin/#c-0-6-0")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c)        //让后续处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run(":9090")
}
