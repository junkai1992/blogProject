package repository

import (
	"blogBack/common"
	"blogBack/model"
	"github.com/jinzhu/gorm"
)

type VisitRepository struct {
	DB *gorm.DB
}

func NewVisitRepository() VisitRepository {
	return VisitRepository{DB: common.GetDB()}
}

func (v VisitRepository) Create(visitCount int) (*model.Visit, error) {
	visit := model.Visit{
		Visitcount: visitCount,
	}
	if err := v.DB.Create(&visit).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (v VisitRepository) Update(visit model.Visit, up int) (*model.Visit, error) {
	if err := v.DB.Model(&visit).UpdateColumn("visitcount", gorm.Expr("visitcount + ?", up)).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (v VisitRepository) SelectVisit() (*model.Visit, error) {
	var visit model.Visit
	if err := v.DB.First(&visit).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}
