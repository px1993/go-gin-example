package models

import (
	"fmt"
	setting "go-gin-example/pkg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	TablePrefix string
)

//base model
type Model struct {
	ID         int   `gorm:"primary_key" json:"id"`
	CreatedOn  int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
}

//注册数据库
func init() {
	var (
		err                          error
		dbName, user, password, host string
	)
	dbconf, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("获取database配置错误，%v", err)
	}
	dbName = dbconf.Key("NAME").String()
	user = dbconf.Key("USER").String()
	password = dbconf.Key("PASSWORD").String()
	host = dbconf.Key("HOST").String()
	TablePrefix = dbconf.Key("TABLE_PREFIX").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库链接失败！%v", err)
		panic("failed to connect database")
	}
}

//关闭
func closeDB() {
	// defer db.DB.Close()
}
