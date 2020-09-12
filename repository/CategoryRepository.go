package repository

import (
	"blogBack/common"
	"blogBack/model"
	"github.com/jinzhu/gorm"
)

// Category 重复查询，合并
type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	// 工厂函数
	return CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	// 标签的创建方法的实现
	category := model.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// 标签的更新方法实现
func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// 查询单个标签方法
func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// 删除标签方法
func (c CategoryRepository) DeleteById(id int) error {
	var category model.Category
	if err := c.DB.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepository) SelectByLists() ([]model.Category, error) {
	var categorys []model.Category
	if err := c.DB.Find(&categorys).Error; err != nil {
		return nil, err
	}
	return categorys, nil
}
