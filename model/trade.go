package model

import (
	"fmt"
	"time"
	"context"
	"campusCard/dao"
	"campusCard/cache"
	"github.com/pkg/errors"
)

var ctx = context.Background()


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
	// 查询账户余额
    var accountInfo dao.AccountInfo
    if err := dao.Db.Where("id = ?", id).First(&accountInfo).Error; err != nil {
        return nil, err
    }


	// 查询 Redis 中 ID 对应的 Limit
    limit, err := cache.Rdb.HGet(ctx, "tradeLimit", id).Float64()
    if err != nil {
        return nil, err
    }

	// 比较 Limit 和交易金额
    if limit + transaction.TAmount < 0 {
        return nil, errors.New("超过限额")
    }

    // 验证交易是否有效
    if err := validTransaction(accountInfo, transaction.TTime, accountInfo.Validation, transaction.TAmount); err != nil {
        return nil, err
    }
	
	// 开启一个事务
	tx := dao.Db.Begin()

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
	
	err = cache.Rdb.HSet(ctx, "tradeLimit", id, limit+transaction.TAmount).Err()
    if err != nil {
		tx.Rollback()
        return nil, err
    }
	
	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		cache.Rdb.HSet(ctx, "tradeLimit", id, limit)
		tx.Rollback()
		return nil, err
	}
	return accountInfo.Balance + transaction.TAmount, nil
}

func validTransaction(accountInfo dao.AccountInfo, transactionTime string, validationTime string, transactionAmount float64) error {
    // 检查账户状态
    if accountInfo.Status != 0 {
        return errors.New("交易失败：账户处于非正常状态")
    }

    // 如果交易时间大于 Validation 时间，更新 Status 为 4
    if transactionTime > validationTime {
        accountInfo.Status = 4
        if err := dao.Db.Save(&accountInfo).Error; err != nil {
            return err
        }
        return errors.New("已过期")
    }

    // 计算新余额
    newBalance := accountInfo.Balance + transactionAmount
    if newBalance < -10 {
        return errors.New("余额不足")
    }

    // 交易有效
    return nil
}

func ChangeBalance(money float64, id string) (interface{}, error) {
	if money <= 0 {
		return nil, errors.New("金额必须为正数")
	}

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

	newBalance := accountInfo.Balance + money
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
		return nil, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		// 如果提交时发生错误，回滚事务并返回错误
		tx.Rollback()
		return nil, err
	}
	return newBalance, nil

}
