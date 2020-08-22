package main

import (
	"blogBack/common"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// 项目启动加载配置
	InitConfig()
	// 初始化数据库
	db := common.InitDB()
	defer db.Close()
	// 中文表单校验初始化
	if err := common.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 返回默认引擎
	r := gin.Default()
	// 注册路由
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		// 指定端口启动
		panic(r.Run(":" + port))
	}
	// 运行服务
	panic(r.Run())
}

func InitConfig() {
	// 获取当前目录
	workDir, _ := os.Getwd()
	// 设置当前读取配置文件名:config下application
	viper.SetConfigName("application")
	// 设置读取文件类型为yml格式
	viper.SetConfigType("yml")
	// 设置配置文件路径
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("ok")
	}
}
