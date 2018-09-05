package main

import (
	"log"
	"github.com/mapsf/api/api/app"
	"github.com/mapsf/api/api/app/db"
)

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Println("UNEXPECTED END OF THE MAIN FUNCTION")
		log.Fatalf(`Panic! "%s"`, err)
		db.Conn.Close()
	}
}

func loadEnvFile() {
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
