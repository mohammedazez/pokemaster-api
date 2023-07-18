package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get() (result map[string]interface{}) {
	result = map[string]interface{}{"db_set": false, "ping": false}

	sqlDB, err := r.db.DB()
	if err != nil {
		return
	}
	result["db_set"] = true

	if err := sqlDB.Ping(); err != nil {
		return
	}
	result["ping"] = true

	return result
}
