package models

import "time"

type Character struct {
	Model
	Login    string `gorm:"not null;unique" json:"login"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Bot      bool   `gorm:"not null" json:"bot"`
	Sex      string `gorm:"not null" json:"sex"`
	Level    uint   `gorm:"not null" json:"level"`

	Experience uint `gorm:"not null" json:"experience"`
	Reputation int  `gorm:"not null" json:"reputation"`

	Online bool `gorm:"not null" json:"online"`

	// координаты
	Position string `gorm:"not null" json:"position"`

	// статы
	// сила
	Power uint `gorm:"not null" json:"power"`
	// удача
	Fortune uint `gorm:"not null" json:"fortune"`
	// ярость
	Rage uint `gorm:"not null" json:"rage"`
	// жизнеспособность
	Vitality uint `gorm:"not null" json:"vitality"`

	// статистика боев
	Wins    uint `gorm:"not null" json:"wins"`
	Defeats uint `gorm:"not null" json:"defeats"`

	RegisteredAt time.Time `gorm:"not null" json:"registered_at"`
}
