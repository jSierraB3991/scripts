package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (repo *Repository) GetDb() *gorm.DB {
	return repo.db
}

func InitiateRepo(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
