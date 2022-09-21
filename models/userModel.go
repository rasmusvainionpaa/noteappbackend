package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// id
	// created at
	// updated at
	// deleted at
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
