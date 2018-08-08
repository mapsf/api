package app

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"log"
	"os"
	"strconv"
)

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	return port
}

func InitServer() {
	log.Printf(`RUNNING API SERVER IN "%s" mode`, os.Getenv("APP_ENV"))
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(getPort()), getRouter()))
}
