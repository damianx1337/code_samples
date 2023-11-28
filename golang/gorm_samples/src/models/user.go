package models

import (
  "gorm.io/gorm"
)
type User struct {
  gorm.Model
	ID uint `gorm:"PrimaryKey"`
	Name string
	FirstName string
}

type ProjectUser struct {
  gorm.Model
	UserID uint
	ProjectID uint
	Roles string 
}

func (user *ProjectUser) BeforeSave(db *gorm.DB) error {
	user.Roles = "default_val"
	return nil
}

type Project struct {
  gorm.Model
	ID uint     `gorm:"PrimaryKey"`
	Name string `gorm:"unique; not_null"`
	Users []User `gorm:"many2many:foreignKey:ID;project_users;"`
}
