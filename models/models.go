package models

import (
	"gin-blog/pkg/setting"
	"gin-blog/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB

type Model struct {
	CreatedAt utils.JsonTime
	UpdatedAt utils.JsonTime
	DeletedAt *utils.JsonTime `sql:"index"`
}

func init() {
	dbPath, err := setting.Cfg.Section("sqlite3").GetKey("path")
	if err != nil {
		log.Fatalf("读取配置文件错误：%v", err)
	}
	db, err = gorm.Open("sqlite3", dbPath.Value())
	if err != nil {
		log.Fatalf("连接数据库错误:%v", err)
	}

	// 全局禁用表名复数形式
	db.SingularTable(true)
	// 开启打印sql
	db.LogMode(true)

	// 同步数据库表
	db.AutoMigrate(&Tag{})
}
