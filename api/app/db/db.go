package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"github.com/mapsf/api/api/app/models"
	"time"
)

var (
	Conn *gorm.DB
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDb() {

	log.Printf("STARTING Conn CONNECTION...")

	var err error
	var conn = fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	Conn, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Panicf("Failed to connect database becauase in reason %s", err.Error())
	}

	Conn.LogMode(true)

	allModels := []interface{}{&models.User{}, &models.Character{}, &models.Battle{}}

	panicIfErr(Conn.DropTableIfExists(allModels...).Error)
	panicIfErr(Conn.AutoMigrate(allModels...).Error)

	chars := []models.Character{
		{
			Level:        0,
			Login:        "ALEXANDR",
			Sex:          "M",
			Online:       true,
			Bot:          false,
			Email:        "jilexandr@gmail.com",
			Experience:   0,
			Password:     "1",
			Position:     "49.444433,32.059766999999965",
			RegisteredAt: time.Now(),
			Reputation:   0,
		},
		{
			Level:        0,
			Login:        "bot1",
			Sex:          "M",
			Online:       true,
			Bot:          true,
			Email:        "bot1@gmail.com",
			Experience:   0,
			Password:     "1",
			Position:     "49.444433,32.050766999999965",
			RegisteredAt: time.Now(),
			Reputation:   0,
		},
		{
			Level:        0,
			Login:        "bot2",
			Sex:          "M",
			Online:       true,
			Bot:          true,
			Email:        "bot2@gmail.com",
			Experience:   0,
			Password:     "1",
			Position:     "49.424433,32.050766999999965",
			RegisteredAt: time.Now(),
			Reputation:   0,
		},
	}

	for _, char := range chars {
		err = Conn.Create(&char).Error
		panicIfErr(err)
	}
}
