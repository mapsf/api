package app

import (
	"github.com/mapsf/api/app/models"
	"github.com/mapsf/api/app/db"
	"time"
	"strings"
	"log"
	"strconv"
)

type Position struct {
	Lat float64
	Lng float64
}

func RunBots() {
	bots := []models.Character{}

	err := db.Conn.Find(&bots, &models.Character{Bot: true}).Error
	if err != nil {
		panic(err)
	}

	for {
		for _, bot := range bots {
			coords := strings.Split(bot.Position, ",")
			lat, _ := strconv.ParseFloat(coords[0], 64)
			lng, _ := strconv.ParseFloat(coords[1], 64)
			position := Position{Lat: lat, Lng: lng}
			log.Println(position)
			time.Sleep(1 * time.Second)
		}
	}
}
