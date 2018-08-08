package main

import (
	"log"
	"github.com/JILeXanDR/docker-compose-golang-postgre-govendor/api/app"
)

func main() {

	defer app.HandlePanic()

	log.Printf("STARTING GOLANG APPLICATION...")

	app.LoadEnvFile()
	app.InitDb()
	app.InitRedis()
	app.InitServer()
}
