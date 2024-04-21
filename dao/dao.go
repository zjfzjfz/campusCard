package dao

import (
	"campusCard/config"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type StudentInfo struct {
    ID   string `gorm:"primary_key"`
    Pwd  string
    Name string
    IId  string
}

type AccountInfo struct {
    CID        string `gorm:"primary_key"`
    ID         string
    Status     int
    Balance    float64
    Validation string
    Limit      float64
}

type TransactionRecords struct {
    ID         string
    TType      int
    TLocation  string
    TTime      string
    TAmount    float64
}

type DebtRepayment struct {
    ID     string
    BDebt  float64
    LDebt  float64
}

func (StudentInfo) TableName() string {
    return "student_info" // 指定自定义表名
}

func (AccountInfo) TableName() string {
    return "account_info" // 指定自定义表名
}

func (TransactionRecords) TableName() string {
    return "transaction_records" // 指定自定义表名
}

func (DebtRepayment) TableName() string {
    return "debt_repayment" // 指定自定义表名
}

var (
	Db  *gorm.DB
	err error
)

func init() {
	//dsn := config.Mysqldb
	//Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db, err = gorm.Open("mysql", config.Localdb)
	if err != nil {
        panic("连接数据库失败")
    }
	Db.AutoMigrate(&StudentInfo{}, &AccountInfo{}, &TransactionRecords{}, &DebtRepayment{})
}
