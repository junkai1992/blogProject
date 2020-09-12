package main

import (
	"blogBack/common"
	"blogBack/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func InitLogger() {

	hook := logger.LogConfig{
		Level:      viper.GetString("loggersource.level"),    //日志级别
		Filename:   viper.GetString("loggersource.filename"), //日志文件路径
		MaxSize:    viper.GetInt("loggersource.maxsize"),     //每个日志保存最大尺寸
		MaxAge:     viper.GetInt("loggersource.maxage"),      //日志文件最多保存多少天
		MaxBackups: viper.GetInt("loggersource.maxbackups"),  //日志文件最多保存多个个备份
	}
	if err := logger.InitLogger(&hook); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
}

func main() {
	// 项目启动加载配置
	InitConfig()
	// 初始化日志
	InitLogger()
	// 初始化数据库
	db := common.InitDB()
	defer db.Close()
	err := common.InitRedisClient()
	if err != nil {
		panic(err)
	}
	// 中文表单校验初始化
	if err := common.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	// 返回默认引擎
	//r := gin.Default()
	r := gin.New()
	// 注册路由
	fmt.Println("-->", viper.GetString("loggersource.filename"))
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
		fmt.Println("init config error!")
	}
}
