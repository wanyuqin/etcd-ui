package db

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var localDns = "root:root@tcp(127.0.0.1:3306)/etcd_ui?charset=utf8mb4&parseTime=True&loc=Local"

var Mysql *gorm.DB

func InitMysql() error {

	sqlDB, err := sql.Open("mysql", localDns)
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:            logger.Default.LogMode(logger.Info),
		AllowGlobalUpdate: true,
	})
	if err != nil {
		return err
	}

	Mysql = db

	return nil
}
