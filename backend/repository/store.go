package repository

import (
	"gorm.io/gorm"
)

type DBStore struct {
}

func NewDBStore(ConDb *gorm.DB) *DBStore {
	return &DBStore{}
}
