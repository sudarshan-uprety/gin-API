package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID    uint   `json:"user" gorm:"foreignKey:ID"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Like      int    `json:"like" binding:"required"`
	Visiblity bool   `json:"visiblity" binding:"required"`
	User      User   `json:"user_data"`
}
