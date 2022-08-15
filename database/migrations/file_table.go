package migrations

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	FileName string
	Path     string
	MimeType string
	Size     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
