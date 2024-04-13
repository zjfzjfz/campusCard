package model

import (
	"campusCard/dao"
	_"time"
)

type Card struct {
	Name string
	Id   string
	Status int
	Balance float32
	Validation string
}

func (Card) TableName() string { return "card" }

func GetCard(id string) (Card, error) {
	var card Card
	if err := dao.Db.Table("account_info").
        Joins("JOIN student_info ON account_info.id = student_info.id").
        Where("student_info.id = ?", id).
        Select("student_info.name, student_info.id, account_info.status, account_info.balance, SUBSTRING_INDEX(account_info.validation, ' ', 1) as validation").
        Scan(&card).Error; err != nil {
        return Card{}, err
    }
	return card, nil
}
