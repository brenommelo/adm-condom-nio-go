package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	FirstName string
	LastName  string
}

func (User) TableName() string {
	return "admin.users"
}
