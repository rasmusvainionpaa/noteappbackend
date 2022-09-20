package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model // creates a gorm model that adds id and creation time
	// id
	// created at
	// updated at
	// deleted at
	UserID         float32 `json:"UserID"`
	Title          string  `json:"Title"`
	Content        string  `json:"Content"`
	InsideHumidity float32 `json:"InsideHumidity"`
}
