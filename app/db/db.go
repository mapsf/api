package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var (
	Conn *gorm.DB
)

func InitDb() {

	log.Printf("STARTING DB CONNECTION...")

	var err error
	var conn = fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	Conn, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Panicf("Failed to connect database becauase in reason %s", err.Error())
	}

	Conn.LogMode(true)

	createSchema()
}
