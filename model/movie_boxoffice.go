package model

import "github.com/google/uuid"

// MovieBoxOffice maps to movieboxoffice table
type MovieBoxOffice struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string
	Collection float32
}
