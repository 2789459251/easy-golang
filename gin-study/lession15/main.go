package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//r.POST("/upload", func(c *gin.Context) {
	//	//从请求中读取文件
	//	f, err := c.FormFile("f1")
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//	} else {
	//		//将读取的文件保存到（服务端）
	//		//filepath := fmt.Sprintf("./%s", f.Filename)
	//		dst := path.Join("./", f.Filename)
	//		c.SaveUploadedFile(f, dst)
	//		c.JSON(http.StatusOK, gin.H{
	//			"status": "ok",
	//		})
	//	}
	//
	//})
	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["f1"]
		for index, file := range files {
			log.Printf(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d file uploaded!", len(files)),
		})
	})
	r.Run(":9090")
}
