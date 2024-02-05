package database

import (
	"log"
	"bufio"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbDataContainer = os.Getenv("DB_DATA_CONTAINER")	// -> CREATE DATABASE <dbname>
	dbSchema = os.Getenv("DB_SCHEMA")

	dbUserName = ""
	dbPassWord = ""

	DB *gorm.DB
)

func Init() {
	// init the other vars

	dsn := "host=" + dbHost + " user=" + dbUserName + " password=" + dbPassWord + " dbname=" + dbDataContainer + " port=" + strconv.Itoa(dbPort) + " TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbSchema + ".",
		},
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	DB = db
}
