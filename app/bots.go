package app

import (
	"github.com/mapsf/api/app/models"
	"github.com/mapsf/api/app/db"
	"time"
	"log"
)

type Position struct {
	Lat float64
	Lng float64
}

func RunBots() {
	bots := []models.Player{}

	err := db.Conn.Find(&bots, &models.Player{Bot: true}).Error
	if err != nil {
		panic(err)
	}

	for {
		for _, bot := range bots {
			position := Position{Lat: bot.Location.Lat, Lng: bot.Location.Lng}
			log.Println(position)
			time.Sleep(1 * time.Second)
		}
	}
}
