package dao

import (
	"campusCard/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	dsn := config.Mysqldb
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

}
