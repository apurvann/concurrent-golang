package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Number interface {
	int | float32
}

type Movie[NUM Number] interface {
	Insert(*gorm.DB, NUM) uuid.UUID
	Update(*gorm.DB, NUM)
	Delete(*gorm.DB, uuid.UUID)
}
