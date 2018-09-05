package main

import (
	"log"
	"github.com/mapsf/api/app"
	"github.com/mapsf/api/app/db"
	"github.com/joho/godotenv"
	"github.com/mapsf/api/app/repositories"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Println("UNEXPECTED END OF THE MAIN FUNCTION")
		log.Fatalf(`Panic! "%s"`, err)
		db.Conn.Close()
	}
}

func loadEnvFile() {
	log.Printf("LOAD ENV FILE...")
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
}

func main() {

	defer handlePanic()

	log.Printf("STARTING GOLANG APPLICATION...")

	loadEnvFile()

	db.InitDb()
	app.InitRedis()

	player1, _ := repositories.GetPlayerByID(1)
	player2, _ := repositories.GetPlayerByID(2)
	distance, err := repositories.GetDistanceBetweenTwoPlayers(player1, player2)
	if err != nil {
		log.Println(err)
	}
	log.Printf("distance is %v", distance)

	//go app.RunBots()

	app.InitServer()
}
