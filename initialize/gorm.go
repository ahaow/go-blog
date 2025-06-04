package initialize

import (
	"go-blog/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	mysqlConf := global.Config.Mysql
	db, err := gorm.Open(mysql.Open(mysqlConf.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlConf.LogLevel()),
	})
	if err != nil {
		global.Log.Error("Failed to connect to MySQL:", zap.Error(err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns)

	return db
}
