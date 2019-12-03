package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsActive  bool
	Type      string
	EmailConfirmed bool
}
 
