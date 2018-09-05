package ws

import (
	"log"
	"github.com/mapsf/api/api/app/db"
	"strconv"
)

type playerPositionChangedData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func onPlayerPositionChanged(data *playerPositionChangedData) {
	//app.DB.Save()
	log.Println("Позиция была изменена.", data)
}

func processClientEvent(event *ClientEvent) {
	switch event.Type {
	case "position":
		data := event.Data.(map[string]interface{})
		lat := strconv.FormatFloat(data["latitude"].(float64), 'f', -1, 64)
		lng := strconv.FormatFloat(data["longitude"].(float64), 'f', -1, 64)
		event.User.Position = lat + "," + lng
		logError(db.Conn.Save(event.User).Error)
		//onPlayerPositionChanged(&playerPositionChangedData{
		//	Latitude:  data["latitude"].(float64),
		//	Longitude: data["longitude"].(float64),
		//})
	}
}
