package main

import (
	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
	email := "test.qq.com"
	user.Email = &email
	return nil

}
func main() {
	//DB.AutoMigrate(&Student{})
	//email := "1826664598@qq.com"
	//添加
	//s1 := Student{
	//	Name:   "小赵爱吃菜",
	//	Age:    8,
	//	Gender: true,
	//	Email:  &email,
	//}
	//err := DB.Create(&s1).Error
	//if err != nil {
	//	fmt.Println("添加失败")
	//}
	//fmt.Println(err)
	//fmt.Println(s1)

	//批量插入
	//var studentList []Student
	//for i := 0; i < 10; i++ {
	//	studentList = append(studentList, Student{
	//		Name:   fmt.Sprintf("小赵爱数%d", i+1),
	//		Age:    9 + i,
	//		Gender: true,
	//		Email:  &email,
	//	})
	//}
	//err1 := DB.Create(&studentList).Error
	//fmt.Println(err1)

	//单条查询
	//var s1 Student
	//DB = DB.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	//DB.First(&s1)
	//DB.Last(&s1)
	//DB.Take(&s1, "name = ?", "小赵爱数5")
	//转意造成的sql注入
	//DB.Take(&s1, fmt.Sprintf("name = '%s'", "小赵爱吃菜'or 1 = 1#"))

	//s1.ID = 5
	//DB.Take(&s1)
	//fmt.Println(s1)

	//查询多条记录

	//var studentList []Student
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	//cnt := DB.Find(&studentList).RowsAffected
	//fmt.Println(cnt)
	////for _, student := range studentList {
	////	fmt.Println(student)
	////}
	//data, _ := json.Marshal(studentList)
	//fmt.Println(string(data))

	//DB.Find(&studentList, []int{4, 7, 9})
	//fmt.Println(studentList)

	//DB.Find(&studentList, "name in (?)", []string{"小赵爱数6", "小赵爱数8"})
	//fmt.Println(studentList)
	//var student Student
	//DB.Take(&student, "name = ?", "错误")
	//student.Name = "正确"
	//DB.Select("name").Save(&student)
	DB.Create(&Student{
		Name: "一休哥",
		Age:  5,
	})
}
