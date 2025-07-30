package models

import "gorm.io/gorm"

type GitRepo struct {
	gorm.Model
	Alias       string `json:"alias" gorm:"not null;unique"`
	URL         string `json:"url" gorm:"not null;unique"`
	Branch      string `json:"branch" gorm:"not null;default:'main'"`
	Description string `json:"description" gorm:"type:text"`
}
