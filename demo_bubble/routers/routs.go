package routers

import (
	"easy-goland/demo_bubble/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static") //url,目录
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/", controller.IndexHandler)
	//api:v1
	v1Group := r.Group("v1")
	{
		//待办事项==数据
		//添加，更改，查看，删除
		v1Group.POST("/todo", controller.CreateAtodo)
		v1Group.GET("/todo", controller.GetAllTodo)
		v1Group.GET("/todo:id", controller.GetATodo)
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "小主，功能不存在~",
		})
	})
	return r
}
