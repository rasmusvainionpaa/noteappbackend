package notecontrollers

import "gorm.io/gorm"

type NoteHandler struct {
	DB *gorm.DB
}
