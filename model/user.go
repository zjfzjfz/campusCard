package model

import "campusCard/dao"

type User struct {
	Id       int
	UserName string
}

func (User) TableName() string { return "user" }

func getUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ? ", id).First(&user).Error
	return user, err
}
