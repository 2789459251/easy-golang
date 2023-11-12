package controller

import (
	"easy-goland/demo_bubble/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var TudoList []models.Tudo
var t models.Tudo

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateAtodo(c *gin.Context) {
	c.BindJSON(&t)

	err := models.CreateAtodo(&t)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, t)
	}
}
func GetAllTodo(c *gin.Context) {
	TudoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		fmt.Println("kkk")
	} else {
		c.JSON(http.StatusOK, TudoList)
	}

}
func GetATodo(c *gin.Context) {

	id := c.Param("id")
	t, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, t)
	}
}
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "出错啦",
		})
	}
	t, err := models.UpdateATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, t)
	}

}
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error1": "数据不存在",
		})
	}
	err := models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error2": err,
		})
	} else {
		c.JSON(http.StatusOK, t)
	}
}
