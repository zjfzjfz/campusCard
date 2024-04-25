package model

import "campusCard/dao"

type User struct {
	Id   string `json:"id"`
	Pwd  string `json:"password"`
	Name string `json:"username"`
	IId  string `json:"iid"`
}

func (User) TableName() string { return "student_info" }

func GetUserInfoByUserId(id string) (User, error) {
	var user User
	err := dao.Db.Where("id = ? ", id).First(&user).Error
	return user, err
}

func AddUser(id string, password string, username string, iid string) (string, error) {
	user := User{Id: id, Pwd: password, Name: username, IId: iid}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
