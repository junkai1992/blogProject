package controller

import (
	"blogBack/common"
	"blogBack/model"
	"blogBack/repository"
	"blogBack/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IVisitController interface {
	RestController
}

type VisitController struct {
	Repository repository.VisitRepository
}

func (v VisitController) Create(ctx *gin.Context) {
	response.Response(ctx, http.StatusBadGateway, 502, nil, "无法访问")
	//trans := common.GetVal()
	//var requestVisit validatorData.CreateVisitRequest
	//if err := ctx.ShouldBind(&requestVisit); err != nil {
	//	errs, ok := err.(validator.ValidationErrors)
	//	if !ok {
	//		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
	//		return
	//	}
	//	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
	//	return
	//}
	//visit, err := v.Repository.Create(requestVisit.Visitcount)
	//if err != nil {
	//	response.Fail(ctx, "没有权限增加来访者", nil)
	//	return
	//	//panic(err)
	//}
	//response.Success(ctx, gin.H{"visit": visit}, "创建来访者成功")
}

func (v VisitController) Update(ctx *gin.Context) {
	host := ctx.ClientIP()
	rdb := common.GetRedisDB()
	num, err := rdb.SAdd("vistors", host).Result()
	visitors, _ := v.Repository.SelectVisit()
	if err != nil {
		response.Success(ctx, gin.H{"visitcount": visitors}, "获取访问数量成功")
		return
	}
	if num == 1 {
		vistor, err := v.Repository.Update(*visitors, 1)
		if err == nil {
			response.Success(ctx, gin.H{"visitcount": vistor}, "获取访问数量成功")
			return
		}
	}
	response.Success(ctx, gin.H{"visitcount": visitors}, "获取访问数量成功")
	return
}

func (v VisitController) Show(ctx *gin.Context) {
	visitors, err := v.Repository.SelectVisit()
	if err != nil {
		response.Fail(ctx, "获取来访者失败", nil)
	}
	response.Success(ctx, gin.H{"visitcount": visitors}, "获取访问数量成功")
}

func (v VisitController) Delete(ctx *gin.Context) {
	response.Response(ctx, http.StatusBadGateway, 502, nil, "无法访问")
}

func NewVisitController() IVisitController {
	repository := repository.NewVisitRepository()
	repository.DB.AutoMigrate(model.Visit{})
	return VisitController{Repository: repository}
}
