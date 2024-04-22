package model

import (
	_ "time"

	"github.com/pkg/errors"
	"campusCard/dao"

)

type Transaction struct {
    TType      int
    TLocation  string
    TTime      string
    TAmount    float64
}


func GetTrade(id string) ([]dao.TransactionRecord, error) {
	var records []dao.TransactionRecord
	result := dao.Db.Where("id = ?", id).Find(&records)
	if result.Error != nil {
		return records, errors.Wrap(result.Error, "failed to Get TradeInfo")
	}
	return records, nil
}

func InsertTransaction(id string, transaction Transaction) (error) {
	record := dao.TransactionRecord{
        ID:         id, // 使用传入的 id
        TType:      transaction.TType,
        TLocation:  transaction.TLocation,
        TTime:      transaction.TTime,
        TAmount:    transaction.TAmount,
    }
	// 插入数据到 "transaction_records" 表
	dao.Db.Table("transaction_records")
    result := dao.Db.Create(&record)
    if result.Error != nil {
        return result.Error
    }
	return  nil
}