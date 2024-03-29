package models

import (
  "time"
  "fmt"
//  "gorm.io/gorm"
)

type LocalTime time.Time
func (t *LocalTime) MarshalJSON() ([]byte, error) {
    tTime := time.Time(*t)
    return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02"))), nil
}
func (t *LocalTime) Scan(v interface{}) error {
    if value, ok := v.(time.Time); ok {
        *t = LocalTime(value)
        return nil
    }
    return fmt.Errorf("can not convert %v to timestamp", v)
}

type User struct {
	Name string
	FirstName string
  RegisteredAt *time.Time `json:"registered_at"`
  //RegisteredAt time.Time `gorm:"registered_at;type:date"`
}

