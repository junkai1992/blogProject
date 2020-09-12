package controller

import (
	"blogBack/common"
	"blogBack/dto"
	"blogBack/model"
	"blogBack/response"
	"blogBack/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	trans := common.GetVal()
	var u LoginUser
	if err := ctx.ShouldBind(&u); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
		return
	}
	email := u.Email
	var user model.User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
		//ctx.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
	}
	// 发送token jwt
	token, err := common.ReleaseToken(user)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	response.Success(ctx, gin.H{"token": token, "user": user.Name, "user_id": user.ID}, "登陆成功")
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "登陆成功",
	//	"result": gin.H{"token": token},
	//})
}

type registerUser struct {
	LoginUser
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Telephone  string `json:"telephone" binding:"required,gte=11,lte=11"`
}

// 查询手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 查询邮箱是否存在
func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	trans := common.GetVal()
	var requestUser registerUser
	if err := ctx.ShouldBind(&requestUser); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
			return
		}
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, common.RemoveTopStruct(errs.Translate(trans)))
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
	telephone := requestUser.Telephone
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "当前当前手机号已经绑定")
		return
	}
	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code":500,"msg":"加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Email:     email,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	// 返回结果
	response.Success(ctx, nil, "注册成功")
}

func Info(ctx *gin.Context) {
	// 上下文取出user对象
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取个人信息")
}
