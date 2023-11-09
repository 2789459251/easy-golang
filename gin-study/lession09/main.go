package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx", "./statics")
	//gin框架中给模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"save": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*") //模版解析

	r.GET("/post/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/home.html", gin.H{ //模版渲染
			"title": "<a: href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/home.html", gin.H{
			"title": "<a: href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	r.Run(":9090")
}
