package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// 文章表
type Post struct {
	// 文章id
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	// 关联用户id
	UserId uint `json:"user_id" gorm:"not null"`
	// 类别id
	CategoryId uint `json:"category_id" gorm:"not null"`
	// 关联类型
	Preface  string `json:"preface" gorm:"not null"`
	Category *Category
	// 文章标题
	Title string `json:"title" gorm:"type:varchar(50);not null"`
	// 头部图片
	HeadImg string `json:"head_img"`
	// 文章内容
	Content string `json:"content" gorm:"type:text;not null"`
	// 创建时间
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"`
	// 更新时间
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"`
	// 阅读数
	Views int `json:"views" gorm:"type:int;default:0"`
	// 是否赞
	Iszan bool `json:"iszan" gorm:"type:bool;default:false"`
	// 赞数量
	ZanCount int `json:"zan_count" gorm:"type:int;default:0"`
}

// 在创建之前给ID赋值uuid
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
