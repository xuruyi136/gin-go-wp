package common

import (
	"fmt"
	"gin-go-wp/model"

	"net/url"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "121.196.195.242"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err" + err.Error())
	}
	//自动创建数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
