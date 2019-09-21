package store

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{gorm: db}
}
