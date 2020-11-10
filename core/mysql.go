package core

import (
	"fiberWeb/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

// initializeDB 初始化数据库配置
func initializeDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_ADDR"), config.Config("DB_DATABASE"))
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic("mysql连接失败...")
	}
	sqlDB, _ := DB.DB()

	maxIdle, err := strconv.Atoi(config.Config("DB_MAX_IDLE"))
	maxOpen, err := strconv.Atoi(config.Config("DB_MAX_OPEN"))
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)

	fmt.Println("mysql连接成功...")
}
