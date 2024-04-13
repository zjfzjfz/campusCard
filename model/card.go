package model

import (
	"fmt"
	_ "fmt"
	"strconv"
	_ "time"

	"github.com/pkg/errors"

	"campusCard/dao"
)

type Card struct {
	Name string
	Id   string
	Status int
	Balance float32
	Validation string
}

type AccountInfo struct {
	Cid   string
    ID    string   
	Status int  
	Balance float32
	Validation string
    Limit float32  
}

func (Card) TableName() string { return "card" }

func GetCard(id string) (Card, error) {
	var card Card
	if err := dao.Db.Table("account_info").
        Joins("JOIN student_info ON account_info.id = student_info.id").
        Where("student_info.id = ?", id).
        Select("student_info.name, student_info.id, account_info.status, account_info.balance, SUBSTRING_INDEX(account_info.validation, ' ', 1) as validation").
        Scan(&card).Error; err != nil {
        return Card{}, err
    }
	return card, nil
}

func ChangeLimit(id string, limit string) (float32,error){
	// 将 limit 转换为 float32 类型
    newLimit64, err := strconv.ParseFloat(limit, 32)
    if err != nil {
        return 0, errors.Wrap(err, "failed to convert limit to float32")
    }
    newLimit := float32(newLimit64)
	fmt.Println(123)
    // 假设 db 是你的数据库连接对象
    result := dao.Db.Table("account_info").Where("id = ?", id).Update("Limit", newLimit)
    if result.Error != nil {
        return 0, errors.Wrap(result.Error, "failed to update limit")
    }

    // 查询更新后的记录
    var updatedAccount AccountInfo
    queryResult := dao.Db.Table("account_info").Where("id = ?", id).First(&updatedAccount)
    if queryResult.Error != nil {
        return 0, errors.Wrap(queryResult.Error, "failed to query updated account info")
    }

    // 返回修改后的 limit 值
    return updatedAccount.Limit, nil
}
