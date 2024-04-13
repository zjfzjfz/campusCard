package model

import (
	"campusCard/dao"
)

type Card struct {
	ID   int
	Name string
}

func (Card) TableName() string { return "card" }

func GetCard(id int) (Card, error) {
	var card Card
	err := dao.Db.Where("id = ?", id).First(&card).Error
	return card, err
}
