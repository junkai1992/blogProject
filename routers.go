package main

import (
	"blogBack/controller"
	"blogBack/logger"
	"blogBack/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware(), logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	// 注册
	r.POST("/api/auth/register", controller.Register)
	// 登陆
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	// 点赞
	r.POST("/api/zan", controller.ZanControl)
	// 图片上传和获取
	r.POST("/upload/title/picture", controller.PictureUpload)
	r.GET("/export/title/picture", controller.PictureExport)
	r.DELETE("/remove/title/picture/:filename", controller.RemoveTitlePicture)
	// 	文章图片  PictureContentUpload
	r.POST("/upload/content/picture", controller.PictureContentUpload)
	r.GET("/export/content/picture", controller.PictureContentExport)
	// 访问者统计
	visitRouter := r.Group("visit")
	visitController := controller.NewVisitController()
	visitRouter.GET("/getvistor", visitController.Show)
	visitRouter.PUT("/putvistor", visitController.Update)
	visitRouter.POST("", visitController.Create)
	// 标签操作
	categoryRouter := r.Group("categories")
	categoryController := controller.NewCategoryController()
	categoryRouter.POST("", categoryController.Create)
	categoryRouter.PUT("/:id", categoryController.Update)
	categoryRouter.GET("/:query", categoryController.Show)
	categoryRouter.DELETE("/:id", categoryController.Delete)
	// 路由分组
	postRoutes := r.Group("/posts")
	// 文章操作:POST,DELETE,PUT登陆验证
	postController := controller.NewPostController()
	postRoutes.POST("", middleware.AuthMiddleware(), postController.Create)
	postRoutes.PUT("/:id", middleware.AuthMiddleware(), postController.Update)
	postRoutes.DELETE("/:id", middleware.AuthMiddleware(), postController.Delete)
	postRoutes.GET("/read/:id", postController.Show)
	postRoutes.GET("/page/list", postController.PageList)
	postRoutes.GET("/new/list", postController.NewPostList)
	// 文章按月分组
	postRoutes.GET("/month/list", postController.MonthPostList)
	// 文章存档
	postRoutes.GET("/files/list", postController.YearsFiles)
	return r
}
