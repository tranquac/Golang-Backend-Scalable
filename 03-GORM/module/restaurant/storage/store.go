package restaurantstorage

import "gorm.io/gorm"

type sqlStorage struct {
	db *gorm.DB
}

func NewSQLStorage(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}
