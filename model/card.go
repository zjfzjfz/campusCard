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
	Balance    float64
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

func ChangeLimit(id string, limit string) (float64, error) {
	newLimit, err := strconv.ParseFloat(limit, 32)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert limit to float32")
	}

	result := dao.Db.Table("account_info").Where("id = ?", id).Update("Limit", newLimit)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to update limit")
	}
	return newLimit, nil
}

func UpdateAccountStatus(id string, status int) error {
	var accountInfo dao.AccountInfo
	err := dao.Db.Where("id = ?", id).First(&accountInfo).Error
	if err != nil {
		return err
	}
	accountInfo.Status = status

	err = dao.Db.Save(&accountInfo).Error
	return err
}

func DeleteAndCreateAccount(id string) (interface{}, error) { // 首先获取原始账户信息
	tx := dao.Db.Begin()
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 删除原始账户
	if err := tx.Where("id = ?", id).Delete(&accountInfo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 创建新的账户，保留原始信息，只修改状态
	newAccount := dao.AccountInfo{
		ID:         accountInfo.ID,
		Status:     2,
		Balance:    accountInfo.Balance,
		Validation: accountInfo.Validation,
		Limit:      accountInfo.Limit,
	}
	if err := tx.Create(&newAccount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return newAccount, nil
}
