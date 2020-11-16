package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB map[string]*gorm.DB

type Config struct {
	Addr       string
	Database   string
	User       string
	Password   string
	PoolSwitch int
	MaxIdle    int
	MaxOpen    int
}

// initializeDB 初始化数据库配置
func initializeDB(name string, cfg Config) {
	fmt.Println("mysql初始化...")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", cfg.User, cfg.Password, cfg.Addr, cfg.Database)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	DB[name], err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic("mysql连接失败...")
	}
	sqlDB, _ := DB[name].DB()

	if cfg.PoolSwitch == 1 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
		sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	}

	fmt.Println("mysql连接成功...")
}
