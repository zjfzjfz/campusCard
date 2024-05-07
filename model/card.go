package model

import (
	"math/rand"
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

func GetLimit(id string) (float64, error) {
	var accountInfo dao.AccountInfo
	err := dao.Db.Where("id = ?", id).First(&accountInfo).Error
	if err != nil {
		return 0, err
	}
	return accountInfo.Limit, nil
}

func GetAccountStatus(id string) (int, error) {
	var accountInfo dao.AccountInfo
	err := dao.Db.Where("id = ?", id).First(&accountInfo).Error
	if err != nil {
		return 0, err
	}
	return accountInfo.Status, nil
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

func CreateAndDeleteAccount(id string) (interface{}, error) { // 首先获取原始账户信息
	tx := dao.Db.Begin()
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	newAccount := dao.AccountInfo{
		CID:        strconv.Itoa(rand.Intn(90000000) + 10000000),
		ID:         accountInfo.ID,
		Status:     2,
		Balance:    accountInfo.Balance,
		Validation: accountInfo.Validation,
		Limit:      accountInfo.Limit,
	}

	//创建新账户
	tx.Table("account_info")
	result := tx.Create(&newAccount)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// 删除原始账户记录
	result = tx.Where("c_id = ?", accountInfo.CID).Delete(&dao.AccountInfo{})
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	return newAccount, nil
}
