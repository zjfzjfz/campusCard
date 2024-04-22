package model

import (
	"strconv"
	_ "time"

	"github.com/pkg/errors"

	"campusCard/dao"
)

type Card struct {
	Name       string
	Id         string
	Status     int
	Balance    float32
	Validation string
}



func GetCard(id string) (Card, error) {
	var card Card
	if err := dao.Db.Table("account_info").
		Joins("JOIN student_info ON account_info.id = student_info.id").
		Where("student_info.id = ?", id).
		Select("student_info.name, student_info.id, account_info.status, account_info.balance, SUBSTRING_INDEX(account_info.validation, ' ', 1) as validation").
		Scan(&card).Error; err != nil {
		return Card{}, errors.Wrap(err, "failed to get card from database")
	}
	return card, nil
}

func ChangeLimit(id string, limit string) (float32, error) {
	newLimit64, err := strconv.ParseFloat(limit, 32)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert limit to float32")
	}
	newLimit := float32(newLimit64)

	result := dao.Db.Table("account_info").Where("id = ?", id).Update("Limit", newLimit)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to update limit")
	}
	return newLimit, nil
}
