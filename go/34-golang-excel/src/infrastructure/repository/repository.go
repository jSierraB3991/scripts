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

func (repo *Repository) existsSisProByCode(value interface{}, code string) (interface{}, error) {
	err := repo.db.Model(value).Where("code = ?", code).Find(&value).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}
func (repo *Repository) existsSisProByCodeGe(value interface{}, code string) error {
	err := repo.db.Model(value).Where("code = ?", code).Find(&value).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetCodesForData(value interface{}, result *[]string) error {
	err := repo.db.Select("code").Model(value).Find(&result).Error
	if err != nil {
		return err
	}
	return nil
}
