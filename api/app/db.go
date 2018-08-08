package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func InitDb() {

	log.Printf("STARTING DB CONNECTION...")

	var err error
	var conn = fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Panicf("Failed to connect database becauase in reason %s", err.Error())
	}

	if !DB.HasTable(&Post{}) {
		DB.CreateTable(&Post{})
	}

	testPost := Post{Author: "Dorper", Message: "GoDoRP is Dope"}
	DB.Create(&testPost)
}
