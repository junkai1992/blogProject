package controller

import (
	"blogBack/common"
	"blogBack/model"
	"blogBack/repository"
	"blogBack/response"
	"blogBack/validatorData"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	// 文章表自动迁移
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

/*

func Register(ctx *gin.Context) {
	DB:= common.GetDB()
	trans := common.GetVal()
	var requestUser registerUser
	if err := ctx.ShouldBind(&requestUser); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx,http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
		return
	}
	email := requestUser.Email
	if isEmailExist(DB, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "当前邮箱已经绑定了")
		return
	}
	name := requestUser.Name
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
*/

func isCategory(db *gorm.DB, name string) bool {
	var category model.Category
	db.Where("name = ?", name).First(&category)
	if category.ID != 0 {
		return true
	}
	return false
}

func (c CategoryController) Create(ctx *gin.Context) {
	trans := common.GetVal()
	var requestCategory validatorData.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
		return
	}
	DB := common.GetDB()
	if isCategory(DB, requestCategory.Name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "当前添加类型重复")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "创建标签成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	trans := common.GetVal()
	var requestCategory validatorData.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
		return
	}
	// 这里获取路由的ID参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}
	DB := common.GetDB()
	if isCategory(DB, requestCategory.Name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "当前修改类型重复")
		return
	}
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "修改分类成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
	}
	response.Success(ctx, gin.H{"category": category}, "获取分类成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, "删除类型失败", nil)
		return
	}
	response.Success(ctx, nil, "删除类型成功")
}
