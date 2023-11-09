package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/JSON", func(c *gin.Context) {
		//1.map->.gin.H:快捷方式
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello word",
		//	"age":     18,
		//}
		data := gin.H{"name": "小王子",
			"message": "hello word?",
			"age":     18,
		}

		c.JSON(http.StatusOK, data)
	})
	r.GET("/msg", func(c *gin.Context) {
		//2.结构体-->大写才可以序列化S
		type msg struct {
			Name    string `json:"name"`
			Age     int
			Message string
		}
		m := msg{
			Name:    "玫瑰",
			Age:     11,
			Message: "you,too",
		}
		c.JSON(http.StatusOK, m) //json的序列化
	})
	r.Run(":9090")
}
