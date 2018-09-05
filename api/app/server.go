package app

import (
	"net/http"
	"log"
	"os"
	"strconv"
	"github.com/gorilla/handlers"
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
	log.Printf(`ENTRY HOST IS http://%s`, os.Getenv("API_HOST"))

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)

	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(getPort()), cors(getRouter())))
}
