package controller

import (
	"blogBack/common"
	"blogBack/model"
	"blogBack/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ZanClick struct {
	Status string `json:"status" binding:"required"`
	Id     string `json:"id" binding:"required"`
}

func ZanControl(ctx *gin.Context) {
	trans := common.GetVal()
	var zanclick ZanClick
	if err := ctx.ShouldBind(&zanclick); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
		return
	}
	status := zanclick.Status
	article_id := zanclick.Id
	host := ctx.ClientIP()
	rdb := common.GetRedisDB()
	if status == "add" {
		ok, err := rdb.SAdd(article_id, host).Result()
		if err != nil {
			panic("点赞失败")
		}
		if ok != 1 {
			response.Fail(ctx, "点赞失败", nil)
			return
		}
	} else {
		ok, err := rdb.SRem(article_id, host).Result()
		if err != nil {
			panic("取消赞失败")
		}
		if ok != 1 {
			response.Fail(ctx, "点赞失败", nil)
			return
		}
	}
	// 文章数据查询
	var post model.Post
	DB := common.GetDB()
	if DB.Where("id = ?", article_id).First(&post).RecordNotFound() {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	if status == "add" {
		if err := DB.Model(&post).UpdateColumn("zan_count", gorm.Expr("zan_count + ?", 1)).Error; err != nil {
			response.Fail(ctx, "给文章点赞失败", nil)
			return
		}
	} else {
		if err := DB.Model(&post).UpdateColumn("zan_count", gorm.Expr("zan_count - ?", 1)).Error; err != nil {
			response.Fail(ctx, "给文章取消赞失败", nil)
			return
		}
	}
	response.Success(ctx, nil, "操作成功")
}
