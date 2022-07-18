package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestInit(t *testing.T) {
	db := GormMysql()
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	//err := db.AutoMigrate(&User{})
	//if err != nil {
	//	return
	//}
	//var a Authority
	//db.Where()
	//db.Create(&User{
	//	Username: "ad",
	//	Password: util.BcryptEncrypt("123"),
	//	Email:    "123456@qq.com",
	//	Authority: Authority{
	//		Model: gorm.Model{ID: 1},
	//	},
	//})

	//db.Create(&Authority{
	//	Model:         gorm.Model{},
	//	AuthorityId:   1,
	//	Name:          "Admin",
	//	ParentId:      0,
	//	Children:      nil,
	//	Menu:          nil,
	//	DefaultRouter: "dashboard",
	//})

	//db.Create(&Authority{
	//	Model:         gorm.Model{},
	//	AuthorityId:   2,
	//	Name:          "Stuff",
	//	ParentId:      0,
	//	Children:      nil,
	//	Menu:          nil,
	//	DefaultRouter: "dashboard",
	//})
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
