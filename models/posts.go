package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID    uint `json:"index"`
	User      User
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Like      int    `json:"like" binding:"required"`
	Visiblity bool   `json:"visiblity" binding:"required"`
}
