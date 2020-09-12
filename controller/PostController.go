package controller

import (
	"blogBack/common"
	"blogBack/model"
	"blogBack/response"
	"blogBack/utils"
	"blogBack/validatorData"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
	NewPostList(ctx *gin.Context)
	MonthPostList(ctx *gin.Context)
	YearsFiles(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost validatorData.CreatePostRequetst
	trans := common.GetVal()
	if err := ctx.ShouldBind(&requestPost); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
		return
	}
	// 获取用户
	user, _ := ctx.Get("user")
	// 创建文章
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		HeadImg:    requestPost.HeadImg,
		Title:      requestPost.Title,
		Content:    requestPost.Content,
		Preface:    requestPost.Preface,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(ctx, nil, "创建文章成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost validatorData.CreatePostRequetst
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, "文章填写有问题", nil)
		return
	}
	// 校验文章是否存在
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	// 校验文章归属
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, "文章不属于你", nil)
		return
	}
	// 更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(ctx, "更新失败", nil)
		return
	}
	response.Success(ctx, nil, "更新文章成功")
}

func (p PostController) Show(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	post.Views++
	if err := p.DB.Model(&post).Update(post).Error; err != nil {
		response.Fail(ctx, "查询该文章失败", nil)
		return
	}
	response.Success(ctx, gin.H{"post": post}, "查询文章成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	// 判断文章是否被删除
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	// 判断删除文章是否用户自己本人
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, "文章不属于你，非法操作", nil)
		return
	}
	// 执行删除操作
	p.DB.Delete(&post)
	response.Success(ctx, nil, "删除成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	tagId, _ := strconv.Atoi(ctx.DefaultQuery("tagId", "0"))
	yearMonth := ctx.DefaultQuery("yearMonth", "")
	var firstmonth int64
	var lastmonth int64
	if yearMonth == "" {
		firstmonth = 0
		lastmonth = 0
	} else {
		arr := strings.Split(yearMonth, "-")
		year_str, month_str := arr[0], arr[1]
		year, err := strconv.Atoi(year_str)
		if err != nil {
			response.Fail(ctx, "年份输入格式有误", nil)
			return
		}
		month, err := strconv.Atoi(month_str)
		if err != nil {
			response.Fail(ctx, "月份输入格式有误", nil)
			return
		}

		now := time.Now()
		currentLocation := now.Location()

		firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		lastmonth = lastOfMonth.Unix()
		firstmonth = firstOfMonth.Unix()
	}
	// 月份时间戳:
	fmt.Println(firstmonth, lastmonth)
	// 分页
	var posts []*model.Post
	// 查询总条数
	var total int
	// 初始化查询
	tx := p.DB.Preload("Category").Order("created_at desc")
	tx_count := p.DB.Model(model.Post{})
	if firstmonth != 0 && lastmonth != 0 {
		tx = tx.Where("created_at BETWEEN ? AND ?", time.Unix(firstmonth, 0).Format("2006-01-02 15:04:05"), time.Unix(lastmonth, 0).Format("2006-01-02 15:04:05"))
		tx_count = tx_count.Where("created_at BETWEEN ? AND ?", time.Unix(firstmonth, 0).Format("2006-01-02 15:04:05"), time.Unix(lastmonth, 0).Format("2006-01-02 15:04:05"))
	}
	if tagId != 0 {
		tx = tx.Where("category_id = ?", tagId)
		tx_count = tx_count.Where("category_id = ?", tagId)
	}
	tx_count.Count(&total)
	tx.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)
	if total == 0 {
		response.Response(ctx, http.StatusNotFound, 404, nil, "没有文章")
		return
	}
	host := ctx.ClientIP()
	rdb := common.GetRedisDB()
	for _, value := range posts {
		articleId := fmt.Sprintf("%s", value.ID)
		ok, _ := rdb.SIsMember(articleId, host).Result()
		if ok == true {
			value.Iszan = ok
		}
	}
	response.Success(ctx, gin.H{"data": posts, "total": total}, "获取文章成功")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}

// 最新更新
func (p PostController) NewPostList(ctx *gin.Context) {
	//var newposts []newpost
	var newposts []model.Post
	p.DB.Select([]string{"id", "created_at", "title"}).Order("created_at desc").Limit(10).Find(&newposts)
	//p.DB.Order("created_at desc").Limit(5).Find(&newposts)
	response.Success(ctx, gin.H{"data": newposts}, "获取文章成功")
}

func (p PostController) MonthPostList(ctx *gin.Context) {
	archives, err := utils.ListPostsMonth()
	if err != nil {
		response.Fail(ctx, "按月获取文章失败", nil)
		return
	}
	response.Success(ctx, gin.H{"data": archives}, "按月获取文章成功")
}

func (p PostController) YearsFiles(ctx *gin.Context) {
	yearPosts, err := utils.ListPostsYear()
	if err != nil {
		response.Fail(ctx, "按年获取文章失败", nil)
		return
	}
	response.Success(ctx, gin.H{"data": yearPosts}, "按年获取文章成功")
}
