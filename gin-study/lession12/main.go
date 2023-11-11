package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin与form
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//1.PostForm

		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//
		//2.DefaultPostForm

		//username := c.DefaultPostForm("username", "莎莎")
		//password := c.DefaultPostForm("password", "***")

		//3.GetPostForm 两个返回值
		username, _ := c.GetPostForm("username")
		password, _ := c.GetPostForm("password")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})

	})
	r.Run(":9090")
}
