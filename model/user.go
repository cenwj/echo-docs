package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id       string `gorm:"primary_key:id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
