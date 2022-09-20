package usercontrollers

import "gorm.io/gorm"

type UserHandler struct {
	DB *gorm.DB
}
