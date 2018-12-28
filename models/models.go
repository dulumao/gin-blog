package models

import (
	"fmt"
	"gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"
	"time"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"type:datetime"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"type:datetime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index" gorm:"type:datetime"`
}

func init() {
	mysqlConf := setting.Cfg.Section("mysql")
	user, _ := mysqlConf.GetKey("USER")
	pwd, _ := mysqlConf.GetKey("PASSWORD")
	host, _ := mysqlConf.GetKey("HOST")
	dbName, _ := mysqlConf.GetKey("DB_NAME")
	mysqlUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName)
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		log.Fatalf("连接数据库错误:%v", err)
	}

	// 全局禁用表名复数形式
	db.SingularTable(true)
	// 开启打印sql
	db.LogMode(true)

	// 迁移表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tag{})
}
