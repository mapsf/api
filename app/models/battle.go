package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Battle struct {
	gorm.Model
	Characters []*Player
	Winner     *Player
	StartedAt  time.Time
	FinishedAt time.Time
}
