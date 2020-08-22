package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(20);not null"`
	Telephone string `json:"telephone" gorm:"varchar(11);not null;unique"`
	Password  string `json:"password" gorm:"size:254;not null"`
	Email     string `json:"email" gorm:"varchar(20);not null;unique"`
	LockState bool   `json:"lock_state" gorm:"default:'1'"`
}
