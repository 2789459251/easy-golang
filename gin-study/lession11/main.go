package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取query string
		//name, _ := c.GetQuery("query") //通过query获取请求url数据
		name := c.DefaultQuery("query", "sb") //通过query获取请求url数据
		age, _ := c.GetQuery("age")           //通过query获取请求url数据
		//if !ok {
		//	name = "一些未知的"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
