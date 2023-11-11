package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问index的GET请求走这个函数
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "get--获得",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "put--更新",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "post--创建",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"methed": "delete--删除",
		})
	})

	//视频的首页详情页
	r.GET("/video/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "video/index",
		})
	})
	//r.GET("/shop/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "shop/index",
	//	})
	//
	//})
	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"me": "小王子",
	//	})
	//})

	//把公用前缀提取，多用于区分不同业务线或API版本
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "shop/index",
			})
		})
	}

	r.Run(":9090")

}
