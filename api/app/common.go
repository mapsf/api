package app

import (
	"log"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func HandlePanic() {
	if err := recover(); err != nil {
		log.Println("UNEXPECTED END OF THE MAIN FUNCTION")
		log.Fatalf(`Panic! "%s"`, err)
		DB.Close()
	}
}

func LoadEnvFile() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
}
