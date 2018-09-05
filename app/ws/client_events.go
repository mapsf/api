package ws

import (
	"github.com/mapsf/api/app/models"
	"github.com/nferruzzi/gormGIS"
	"github.com/mapsf/api/app/db"
	"log"
	"github.com/mapsf/api/app/repositories"
)

type playerPositionChangedData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func onPlayerPositionChanged(player *models.Player, i interface{}) {

	data := i.(map[string]interface{})

	a := playerPositionChangedData{
		Latitude:  data["latitude"].(float64),
		Longitude: data["longitude"].(float64),
	}

	player.Location = gormGIS.GeoPoint{Lat: a.Latitude, Lng: a.Longitude}

	err := db.Conn.Save(player).Error
	if err != nil {
		logError(err)
		return
	}

	log.Println("Позиция была изменена.", data)

	players, _ := repositories.FindNearPlayers(player, 1000)
	log.Printf("Найдено %v игроков в радиусе 1км", len(players))
}

func processClientEvent(player *models.Player, event *ClientEvent) {
	debug(`[RECEIVED] Process client event "%v", with data %v`, event.Type, event.Data)
	switch event.Type {
	case "position":
		onPlayerPositionChanged(player, event.Data)
	default:
		debug("Undefined event %v", event.Type)
	}
}
