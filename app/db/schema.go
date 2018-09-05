package db

import (
	"time"
	"github.com/mapsf/api/app/models"
	"github.com/nferruzzi/gormGIS"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Player{},
		&models.Battle{},
	}
}

func createSchema() {

	//Conn.Exec("DROP EXTENSION IF EXISTS postgis_topology CASCADE")
	//Conn.Exec("DROP EXTENSION IF EXISTS postgis CASCADE")

	Conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis")
	Conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis_topology")

	panicIfErr(Conn.DropTableIfExists(getModels()...).Error)
	panicIfErr(Conn.AutoMigrate(getModels()...).Error)

	chars := []models.Player{
		{
			Level:        0,
			Login:        "ALEXANDR",
			Sex:          "M",
			Online:       true,
			Bot:          false,
			Email:        "jilexandr@gmail.com",
			Experience:   0,
			Password:     "1",
			Location:     gormGIS.GeoPoint{Lat: 0, Lng: 0},
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
			Location:     gormGIS.GeoPoint{Lat: 10, Lng: 10},
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
			Location:     gormGIS.GeoPoint{Lat: 20, Lng: 20},
			RegisteredAt: time.Now(),
			Reputation:   0,
		},
	}

	for _, char := range chars {
		panicIfErr(Conn.Create(&char).Error)
	}
}
