package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	ctx context.Context
}

func (repo *Repository) GetDb() *gorm.DB {
	return repo.db
}

func InitiateRepo(db *gorm.DB, ctx context.Context) *Repository {
	return &Repository{
		db:  db,
		ctx: ctx,
	}
}
