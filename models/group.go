package models

import (
	"gorm.io/gorm"
)

var SuperGroup Group

type Group struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;uniqueIndex;size:200"`
}

func (g *Group) Save() error {
	if g.ID == 0 {
		return DB().Create(g).Error
	}
	return DB().Save(g).Error
}

func init() {
	SuperGroup.ID = 1
	SuperGroup.Name = "超级管理员"
}
