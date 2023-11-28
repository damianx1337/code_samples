package models

import (
  "log"
  "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
  dbURL := "postgres://postgres:postgres@localhost:5432/crud"

  db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

  if err != nil {
      log.Fatalln(err)
  }

  if err := db.SetupJoinTable(&Project{}, "Users", &ProjectUser{}); err != nil {
		println(err.Error())
		panic("Failed to setup join table")
	}

  db.AutoMigrate(&User{}, &Project{})

  bogdan := User { Name: "Bogdan" }
  jorg := User { Name: "Jorg" }

	db.Create(&bogdan)
	db.Create(&jorg)

	db.Create(&Project{Name: "Test1", Users: []User {{ID: bogdan.ID},{ID: jorg.ID}}})

  steffen := User { Name: "Steffen" }
	db.Create(&steffen)

  //var project Project
  //db.Model(&project).Association("Users")
  //db.Model(&project).Where("name = ?", "Test1").Association("Users").Append([]User{{ID: steffen.ID}})

	DB = db
}
