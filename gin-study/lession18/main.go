package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index进来")
	name := c.MustGet("name")
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义中间件
func m1(c *gin.Context) {
	fmt.Println("m1进来了")
	//计时
	start := time.Now()
	c.Next() //调用后续处理函数
	//c.Abort() //阻止调用
	cost := time.Since(start)
	fmt.Printf("花费时间：%v", cost)
	fmt.Println("m1走了")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in")
	c.Set("name", "zy")
	c.Next()
	return
	fmt.Println("m2 out")
}

// 中间件	钩子函数	公共逻辑
func main() {
	r := gin.Default()         //默认使用Logger和Recovery中间件，不要可以用gin.new
	r.Use(m1, m2, auth(false)) //全局注册m1
	r.GET("/index", indexHandler)
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	//路由组注册中间件1.
	shopGroup := r.Group("/shop", auth(true))
	{
		shopGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	//2.
	xxGroup := r.Group("/xx", auth(true))
	xxGroup.Use(auth(true))
	{
		xxGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(":9090")
}

// 有开关的控制的包装中间件
func auth(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			//具体逻辑
			c.Next()
		} else {
			c.Next()
		}

	}
}
