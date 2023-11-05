package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var mysqlLogger logger.Interface
var DB *gorm.DB

type Student struct {
	Name      string
	Age       int
	MyStudent string
}

func init() {
	username := "root"  //账号
	password := "123"   //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "gorm"    //数据库名
	timeout := "10s"    //连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info) //全局日志===显示日志等级
	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //开启跳过默认事物
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "f_",  // 表名前缀
			SingularTable: false, // 单数表名，不开启单数
			NoLowerCase:   false, // 关闭小写转换
		},
		//Logger: mysqlLogger,//全局日志
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	DB = db
}
func main() {
	//DB = DB.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	DB.Debug().AutoMigrate(&Student{})
}
