package app

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Author  string
	Message string
}
