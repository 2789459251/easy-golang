package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	uaername string
	password string
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	uaername: username,
		//	password: password,
		//}
		//fmt.Println(u)

		//2.数据较多，绑定
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"massage": "ok",
			})
		}
	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"massage": "ok",
			})
		}
	})
	r.Run(":9090")
}
