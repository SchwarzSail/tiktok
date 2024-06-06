package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"tiktok/cmd/video/config"
)

var DB *gorm.DB

func InitMySQL() {
	conf := config.Config.Mysql
	dsn := conf.UserName + ":" + conf.MysqlPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	DB = db
}
