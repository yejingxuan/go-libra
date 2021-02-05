package gdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

type ConfigDB struct {
	DSN                string
	Dialect            string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    int
	LogMode            bool
}

//标准配置
func StdConfig() ConfigDB {
	return rawConfig(fmt.Sprintf("system.db"))
}

func rawConfig(name string) ConfigDB {
	config := ConfigDB{
		DSN:                viper.GetString(fmt.Sprintf("%s.dsn", name)),
		Dialect:            viper.GetString(fmt.Sprintf("%s.dialect", name)),
		MaxOpenConnections: viper.GetInt(fmt.Sprintf("%s.maxopenconnections", name)),
		MaxIdleConnections: viper.GetInt(fmt.Sprintf("%s.maxidleconnections", name)),
		ConnMaxLifetime:    viper.GetInt(fmt.Sprintf("%s.connmaxlifetime", name)),
		LogMode:            viper.GetBool(fmt.Sprintf("%s.logmode", name)),
	}
	return config
}

//创建 mq 连接
func (stdConfig ConfigDB) Build() (*gorm.DB, error) {
	db, err := gorm.Open(stdConfig.Dialect, stdConfig.DSN)
	//可以同时打开的连接数
	db.DB().SetMaxOpenConns(stdConfig.MaxOpenConnections)
	//允许在连接池中最多保留空闲连接
	db.DB().SetMaxIdleConns(stdConfig.MaxIdleConnections)
	//设置了连接可重用的最大时间长度
	db.DB().SetConnMaxLifetime(time.Duration(stdConfig.ConnMaxLifetime) * time.Hour)
	//日志输出
	db.LogMode(stdConfig.LogMode)
	return db, err
}
