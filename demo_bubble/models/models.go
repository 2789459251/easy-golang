package models

import (
	"easy-goland/demo_bubble/dao"
)

type Tudo struct {
	ID     int    `json:"id" gorm:"primaryKey;unique;comment:主键"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateAtodo(t *Tudo) (err error) {
	//存数据库
	err = dao.DB.Create(&t).Error
	return err

}
func GetAllTodo() (TudoList *[]Tudo, err error) {
	err = dao.DB.Find(&TudoList).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetATodo(id string) (t *Tudo, err error) {
	err = dao.DB.Take(&t, "id = ?", id).Error
	return t, err

}
func UpdateATodo(id string) (t *Tudo, err error) {
	err = dao.DB.Where("id =?", id).First(&t).Error
	if err != nil {
		return t, err
	}
	err = dao.DB.Save(&t).Error
	return t, err
}
func DeleteATodo(id string) (err error) {
	var t Tudo
	err = dao.DB.Where("id=?", id).First(&Tudo{}).Error
	if err != nil {
		return err
	}
	err = dao.DB.Delete(&t).Error
	return err
}
