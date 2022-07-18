package initialize

import (
	"system/rpc/internal/global"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGORM() *gorm.DB {
	switch global.GVA_CONFIG.DB.Type {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.DB
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.MysqlDSN(),
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
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db
	}
}

func GormPgSql() *gorm.DB {
	p := global.GVA_CONFIG.DB
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.PostgresDSN(),
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConn)
		sqlDB.SetMaxOpenConns(p.MaxOpenConn)
		return db
	}
}
