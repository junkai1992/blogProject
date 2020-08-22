package main

import (
	"blogBack/controller"
	"blogBack/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRouter := r.Group("categories")
	categoryController := controller.NewCategoryController()
	categoryRouter.POST("", categoryController.Create)
	categoryRouter.PUT("/:id", categoryController.Update)
	categoryRouter.GET("/rep/*id", categoryController.Show)
	categoryRouter.DELETE("/:id", categoryController.Delete)

	return r
}
