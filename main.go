package main

import (
	"log"
	"github.com/mapsf/api/app"
	"github.com/mapsf/api/app/db"
)

import (
	"github.com/joho/godotenv"
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

	//go app.RunBots()

	app.InitServer()
}
