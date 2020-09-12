package common

import (
	"blogBack/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

var DB *gorm.DB

// 开启数据库连接池
func InitDB() *gorm.DB {
	// 初始化mysql配置
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	// 拼接mysql配置
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	// 连接mysql
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed connect db... err:" + err.Error())
	}
	// 自动创建(migrate)数据表user
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	// 返回模型对象实例，用于视图函数使用CURD
	return DB
}
