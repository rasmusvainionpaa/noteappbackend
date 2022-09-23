package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model // creates a gorm model that adds id and creation time
	// id
	// created at
	// updated at
	// deleted at
	UserID    float32   `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Alert     time.Time `json:"InsideHumidity"`
	Important bool      `json:"isImportant"`
}
