package models

import (
  "log"
  "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
  dbURL := "postgres://postgres:postgres@localhost:5432/crud"

  db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
  })

  if err != nil {
      log.Fatalln(err)
  }

  //bogdan := User { Name: "Bogdan", registered_at: "1992-09-23" }
  bogdan := User { Name: "Bogdan" }
  jorg := User { Name: "Jorg" }

	db.Create(&bogdan)
	db.Create(&jorg)

	DB = db
}
