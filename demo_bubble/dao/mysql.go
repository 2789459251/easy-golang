package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() (DB *gorm.DB) {
	username := "root"
	passwd := "123"
	host := "127.0.0.1"
	port := 3306
	Dbname := "bubble"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, passwd, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //开启跳过默认事物
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 单数表名，不开启单数
			NoLowerCase:   false, // 关闭小写转换
		},
	})
	DB = db
	return DB
}
