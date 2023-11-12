package main

import (
	"easy-goland/demo_bubble/dao"
	"easy-goland/demo_bubble/models"
	"easy-goland/demo_bubble/routers"
)

func main() {
	dao.DB = dao.Init()
	err := dao.DB.Debug().AutoMigrate(&models.Tudo{})
	if err != nil {
		return
	}
	r := routers.SetupRouter()
	r.Run(":9090")
}
