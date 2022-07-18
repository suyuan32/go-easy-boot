package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestInit(t *testing.T) {
	db := GormMysql()
	err := db.AutoMigrate(&Menu{})
	if err != nil {
		return
	}
	//err := db.AutoMigrate(&User{})
	//if err != nil {
	//	return
	//}
}

func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/go-easy-boot?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
