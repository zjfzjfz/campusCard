package model

import (
	"campusCard/dao"
)

func ChangeBalance(money float64, id string) error {
	// 查询账户信息
	var accountInfo dao.AccountInfo
	err := dao.Db.Where("id = ?", id).First(&accountInfo).Error
	if err != nil {
		return err
	}

	// 更新余额
	newBalance := accountInfo.Balance + money

	// 更新数据库中的余额
	err = dao.Db.Model(&accountInfo).Where("id = ?", id).Update("balance", newBalance).Error
	if err != nil {
		return err
	}

	return nil
}
