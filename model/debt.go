package model

import (
	"fmt"
	"time"
	_ "time"

	"campusCard/dao"
	"github.com/pkg/errors"
)

type Debt struct {
	BDebt float64
	LDebt float64
}

func GetDebt(id string) ([]dao.DebtRepayment, error) {
	var records []dao.DebtRepayment
	result := dao.Db.Where("id = ?", id).Find(&records)
	if result.Error != nil {
		return records, errors.Wrap(result.Error, "failed to Get DebtInfo")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no matching records found")
	}
	return records, nil
}

func ChangeBath(id string) (interface{}, error) {

	tx := dao.Db.Begin()

	// 检索账户余额
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	if accountInfo.Status != 0 {
		tx.Rollback()
		return nil, fmt.Errorf("交易失败：账户处于非正常状态")
	}
	var debtInfo dao.DebtRepayment
	if err := tx.Where("id = ?", id).First(&debtInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}

	// 更新余额
	newBalance := accountInfo.Balance + debtInfo.BDebt
	if newBalance < 0 {
		tx.Rollback()
		return nil, errors.New("余额不足")
	}

	// 创建交易记录
	record := dao.TransactionRecord{
		ID:        id,
		TType:     3,
		TLocation: "一卡通",
		TTime:     time.Now().Format("2006-01-02 15:04:05"),
		TAmount:   debtInfo.BDebt,
	}

	// 插入交易记录
	tx.Table("transaction_records")
	result := tx.Create(&record)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	// 将当前行的 "b_debt" 列的值置为 0
	if err := tx.Model(&debtInfo).Where("id = ?", id).Update("b_debt", 0).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	return newBalance, nil
}

func ChangeLibrary(id string) (interface{}, error) {

	tx := dao.Db.Begin()

	// 检索账户余额
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	if accountInfo.Status != 0 {
		tx.Rollback()
		return nil, fmt.Errorf("交易失败：账户处于非正常状态")
	}
	var debtInfo dao.DebtRepayment
	if err := tx.Where("id = ?", id).First(&debtInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}

	// 更新余额
	newBalance := accountInfo.Balance + debtInfo.LDebt
	if newBalance < 0 {
		tx.Rollback()
		return nil, errors.New("余额不足")
	}

	// 创建交易记录
	record := dao.TransactionRecord{
		ID:        id,
		TType:     3,
		TLocation: "一卡通",
		TTime:     time.Now().Format("2006-01-02 15:04:05"),
		TAmount:   debtInfo.LDebt,
	}

	// 插入交易记录
	tx.Table("transaction_records")
	result := tx.Create(&record)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	// 将当前行的 "b_debt" 列的值置为 0
	if err := tx.Model(&debtInfo).Where("id = ?", id).Update("l_debt", 0).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	return newBalance, nil
}
