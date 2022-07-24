package model

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	//var a Role
	//db.Where()
	//db.Create(&User{
	//	Username: "ad",
	//	Password: util.BcryptEncrypt("123"),
	//	Email:    "123456@qq.com",
	//	Role: Role{
	//		Model: gorm.Model{ID: 1},
	//	},
	//})

	// db.Create(&Role{
	// 	Model:         gorm.Model{},
	// 	RoleId:        1,
	// 	Name:          "Admin",
	// 	ParentId:      0,
	// 	Children:      nil,
	// 	Menu:          nil,
	// 	DefaultRouter: "dashboard",
	// })

	// db.Create(&Role{
	// 	Model:         gorm.Model{},
	// 	RoleId:        2,
	// 	Name:          "Stuff",
	// 	ParentId:      0,
	// 	Children:      nil,
	// 	Menu:          nil,
	// 	DefaultRouter: "dashboard",
	// })
	// result := db.Omit("Role").Create(&User{
	// 	UUID:     uuid.New(),
	// 	Username: "admin",
	// 	Nickname: "admin",
	// 	Password: util.BcryptEncrypt("123456"),
	// 	Email:    "admin@qq.com",
	// 	Role: Role{
	// 		Model: gorm.Model{ID: 2},
	// 	},
	// })

	// fmt.Println(result, result.Error.Error(), result.RowsAffected)

	// var u User
	// db.Where(&User{Username: "admin"}).First(&u)
	// t.Error(fmt.Sprintf("%+v", u))
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

func TestUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Error(uuid.New())
	}
}

func TestAdd(t *testing.T) {
	db := GormMysql()
	for i := 0; i < 10000; i++ {
		result := db.Create(&User{
			UUID:        uuid.NewString(),
			Username:    fmt.Sprintf("user_%d", i),
			Password:    "123",
			Nickname:    fmt.Sprintf("user_%d", i),
			SideMode:    "123",
			Avatar:      "",
			BaseColor:   "",
			ActiveColor: "",
			RoleId:      0,
			Mobile:      "",
			Email:       "",
			Status:      0,
		})
		if result.Error != nil {
			t.Error(result.Error)
		}
	}
}
