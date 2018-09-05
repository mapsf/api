package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Battle struct {
	gorm.Model
	Characters []*Character
	Winner     *Character
	StartedAt  time.Time
	FinishedAt time.Time
}
