package models

import "gorm.io/gorm"

//Store real extracted struct from git repo
//used for syncing from git repo -> db

type Struct struct {
	gorm.Model
	Name    string  `gorm:"not null;unique" json:"name"`
	Comment string  `gorm:"type:text" json:"comment"`
	Fields  []Field `gorm:"foreignKey:StructID" json:"fields"`
	RepoID  uint    `json:"repo_id" gorm:"not null;index"`
}
