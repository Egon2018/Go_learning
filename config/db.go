package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	"exchangeapp/global"
)

func initDb() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("fail to initialize database", err)
	}
	SqlDb, err := db.DB()
	SqlDb.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	SqlDb.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	SqlDb.SetConnMaxLifetime(time.Hour)
	if err != nil {
		fmt.Println("fail to config database", err)
	}

	global.Db = db
}
