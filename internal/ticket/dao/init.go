package dao

import (
	"TTMS/configs/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	//db, err := sql.Open("mysql", consts.MySQLDefaultDSN)
	//db.SetMaxOpenConns(10000)
	//db.SetMaxIdleConns(500)
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,                                //禁用默认事务操作
			Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
		},
	)

	if err != nil {
		panic(err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
}
