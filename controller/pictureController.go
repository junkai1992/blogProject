package controller

import (
	"blogBack/response"
	"blogBack/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func PictureUpload(ctx *gin.Context) {
	files, _ := ctx.FormFile("file")
	filename := files.Filename
	urlPath := fmt.Sprintf("http://%s:%s/export/title/picture/?filename=%s", viper.GetString("server.host"), viper.GetString("server.port"), filename)
	ok, _ := utils.PathExists(urlPath)
	if ok {
		response.Success(ctx, gin.H{"filename": filename, "urlpath": urlPath}, "上传文章标题图片成功")
		return
	}
	if files != nil {
		post_title_path := viper.GetString("imgespath.titlepath")
		if err := ctx.SaveUploadedFile(files, post_title_path+filename); err != nil {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "上传失败")
			return
		}
	}
	response.Success(ctx, gin.H{"filename": filename, "urlpath": urlPath}, "上传文章图片成功")
}

// 文章图片上传
func PictureContentUpload(ctx *gin.Context) {
	files, _ := ctx.FormFile("image")
	filename := files.Filename
	urlPath := fmt.Sprintf("http://%s:%s/export/content/picture/?filename=%s", viper.GetString("server.host"), viper.GetString("server.port"), filename)
	ok, _ := utils.PathExists(urlPath)
	if ok {
		response.Success(ctx, gin.H{"filename": filename, "urlpath": urlPath}, "上传文章图片成功")
		return
	}
	if files != nil {
		// static/titlePost/
		post_title_path := viper.GetString("imgespath.contentpath")
		if err := ctx.SaveUploadedFile(files, post_title_path+filename); err != nil {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "上传失败")
			return
		}
	}
	response.Success(ctx, gin.H{"filename": filename, "urlpath": urlPath}, "上传文章图片成功")
	return

}

//文章图片展示
func PictureContentExport(ctx *gin.Context) {
	filename, err := ctx.GetQuery("filename")
	if !err {
		response.Response(ctx, http.StatusNotFound, 404, nil, "没有找到要展示的照片")
		return
	}
	post_title_path := viper.GetString("imgespath.contentpath")
	path := post_title_path + filename
	ok, _ := utils.PathExists(path)
	if !ok {
		response.Response(ctx, http.StatusNotFound, 404, nil, "没有找到要展示的照片")
		return
	}
	ctx.File(path)
}

func PictureExport(ctx *gin.Context) {
	filename, err := ctx.GetQuery("filename")
	if !err {
		response.Response(ctx, http.StatusNotFound, 404, nil, "没有找到要展示的照片")
		return
	}
	post_title_path := viper.GetString("imgespath.titlepath")
	path := post_title_path + filename
	ok, _ := utils.PathExists(path)
	if !ok {
		response.Response(ctx, http.StatusNotFound, 404, nil, "该图片或是已被删除")
		return
	}
	ctx.File(path)
}

func RemoveTitlePicture(ctx *gin.Context) {
	filename := ctx.Params.ByName("filename")
	pictureTitlePath := viper.GetString("imgespath.titlepath") + filename
	if _, err := os.Stat(pictureTitlePath); err != nil {
		if os.IsNotExist(err) {
			response.Response(ctx, http.StatusNotFound, 404, nil, "要删除文件不存在")
			return
		} else {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "访问错误")
			return
		}
	} else {
		err := os.Remove(pictureTitlePath)
		if err != nil {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "删除图片失败")
			return
		}
	}
	response.Success(ctx, nil, "删除图片成功")
}
