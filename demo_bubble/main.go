package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
)

var DB *gorm.DB

// model
func init() {
	username := "root"
	passwd := "123"
	host := "127.0.0.1"
	port := "3306"
	Dbname := "bubble"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, passwd, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //开启跳过默认事物
		NamingStrategy: schema.NamingStrategy{
			// 表名前缀
			SingularTable: false, // 单数表名，不开启单数
			NoLowerCase:   false, // 关闭小写转换
		},
	})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	DB = db
}

type Tudo struct {
	ID     int    `json:"id" gorm:"primaryKey;unique;comment:主键"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func check(c *gin.Context) {
	var TudoList []Tudo
	DB.Take(&TudoList)
	//查询数据库
	c.HTML(http.StatusOK, "index.html", TudoList)
}
func main() {
	//创建数据库
	//sql :CREATE table bubble
	//连接数据库
	DB = DB.Session(&gorm.Session{})
	DB.Debug().AutoMigrate(&Tudo{})

	r := gin.Default()
	r.Static("/static", "static") //url,目录
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/bubble", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//api:v1
	v1Group := r.Group("/v1")
	{
		//待办事项==数据
		//添加，更改，查看，删除
		v1Group.POST("todo", func(c *gin.Context) {
			var t Tudo
			c.ShouldBind(&t)
			err := DB.Create(&t).Error
			if err != nil {
				fmt.Println("添加数据库失败")
			}
		}, check)
		//查看所有/某
		//钩子函数
		v1Group.GET("/todo", check)

		v1Group.GET("/todo:id", func(c *gin.Context) {
			var t Tudo
			id := c.Param("id")
			DB.Take(&t, "id = ?", id)
			c.HTML(http.StatusOK, "index.html", t)
		})

		v1Group.PUT("/tudo/id", func(c *gin.Context) {
			var t Tudo
			err := c.ShouldBind(&t)
			if err != nil {
				panic("绑定表单数据错误" + err.Error())
			}
			DB.Save(&t)
			c.HTML(http.StatusOK, "index.html", t)
		}, check)
		v1Group.DELETE("/tudo/id", func(c *gin.Context) {
			var t Tudo
			err := c.ShouldBind(&t)
			if err != nil {
				panic("数据删除获取失误" + err.Error())
			}
			DB.Delete(&t, t.ID)
			c.HTML(http.StatusOK, "index.html", t)
		}, check)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "小主，功能不存在~",
		})
	})

	r.Run(":9090")
}
