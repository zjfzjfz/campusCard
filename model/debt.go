package model

import (
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
