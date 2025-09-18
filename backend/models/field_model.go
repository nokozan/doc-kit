package models

import "gorm.io/gorm"

type Field struct {
	gorm.Model
	StructID uint   `json:"struct_id" gorm:"not null;index"`
	Name     string `json:"name" gorm:"not null"`
	Type     string `json:"type" gorm:"not null"`
	Tag      string `json:"tag"`
	Comment  string `json:"comment" gorm:"type:text"`
}
