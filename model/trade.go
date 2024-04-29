package model

import (
	"time"
	_ "time"

	"campusCard/dao"
	"github.com/pkg/errors"
)

type TransactionResult struct {
	NewBalance interface{}
	Err        error
}


type Transaction struct {
	TType     int
	TLocation string
	TTime     string
	TAmount   float64
}

func GetTrade(id string) ([]dao.TransactionRecord, error) {
	var records []dao.TransactionRecord
	result := dao.Db.Where("id = ?", id).Find(&records)
	if result.Error != nil {
		return records, errors.Wrap(result.Error, "failed to Get TradeInfo")
	}
	return records, nil
}

func InsertTransaction(id string, transaction Transaction) (interface{}, error) {
	// 开启一个事务
	tx := dao.Db.Begin()

	// 检索账户余额
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}

	newBalance := accountInfo.Balance + transaction.TAmount
	if newBalance < 0 {
		tx.Rollback()
		return nil, errors.New("余额不足")
	}

	record := dao.TransactionRecord{
		ID:        id, // 使用传入的 id
		TType:     transaction.TType,
		TLocation: transaction.TLocation,
		TTime:     transaction.TTime,
		TAmount:   transaction.TAmount,
	}
	// 插入数据到 "transaction_records" 表
	tx.Table("transaction_records")
	result := tx.Create(&record)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	return newBalance, nil
}

func ChangeBalance(money float64, id string) error {
	if money <= 0 {
		return errors.New("金额必须为正数")
	}

	tx := dao.Db.Begin()

	// 检索账户余额
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新余额
	newBalance := accountInfo.Balance + money

	// 更新数据库中的余额
	if err := tx.Model(&accountInfo).Where("id = ?", id).Update("balance", newBalance).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 创建交易记录
	record := dao.TransactionRecord{
		ID:        id,
		TType:     1,
		TLocation: "一卡通",
		TTime:     time.Now().Format("2006-01-02 15:04:05"),
		TAmount:   money,
	}

	// 插入交易记录
	tx.Table("transaction_records")
	result := tx.Create(&record)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}
	return nil

}

func ChangeDebt(id string) error {

	tx := dao.Db.Begin()

	// 检索账户余额
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
	}

	var debtInfo dao.DebtRepayment
	if err := tx.Where("id = ?", id).First(&debtInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新余额
	newBalance := accountInfo.Balance + debtInfo.BDebt
	if newBalance < 0 {
		tx.Rollback()
		return errors.New("余额不足")
	}

	// 更新数据库中的余额
	if err := tx.Model(&accountInfo).Where("id = ?", id).Update("balance", newBalance).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
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
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}
	return nil
}

func ChangeLibrary(id string) error {

	tx := dao.Db.Begin()

	// 检索账户余额
	var accountInfo dao.AccountInfo
	if err := tx.Where("id = ?", id).First(&accountInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
	}

	var debtInfo dao.DebtRepayment
	if err := tx.Where("id = ?", id).First(&debtInfo).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
	}

	// 更新余额
	newBalance := accountInfo.Balance + debtInfo.LDebt
	if newBalance < 0 {
		tx.Rollback()
		return errors.New("余额不足")
	}

	// 更新数据库中的余额
	if err := tx.Model(&accountInfo).Where("id = ?", id).Update("balance", newBalance).Error; err != nil {
		// 回滚事务并返回错误
		tx.Rollback()
		return err
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
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return err
	}
	return nil
}
