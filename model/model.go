package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Gender    string
	Email     string
	Qq        string
	Birthdate string
	Cookie    string
}
