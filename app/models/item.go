package models

import "github.com/jinzhu/gorm"

// игровые предметы
type Item struct {
	gorm.Model
	Owner *Player
}
