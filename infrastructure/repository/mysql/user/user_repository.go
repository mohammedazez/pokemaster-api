package user

import (
	"context"
	"errors"
	domain "pokemaster-api/core/domain/user"
	"pokemaster-api/infrastructure/repository/mysql/transactor"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Users struct {
		ID        string `gorm:"primaryKey;column:id"`
		FullName  string `gorm:"column:fullname"`
		Email     string `gorm:"column:email"`
		Password  string `gorm:"column:password"`
		CreatedAt string `gorm:"column:created_at"`
		UpdatedAt string `gorm:"column:updated_at"`
	}
)

func (Users) TableName() string {
	return "users"
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) getDB(ctx context.Context) *gorm.DB {
	dbWithTx := transactor.ExtractTx(ctx)
	if dbWithTx != nil {
		return dbWithTx
	}
	return repo.db
}

func (repo *Repository) RegisterUser(ctx context.Context, inData *domain.User) error {
	users := mappingInput(inData)

	db := repo.getDB(ctx)
	if err := db.Model(&users).Create(&users).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetAllListUsers() ([]domain.User, error) {
	users := make([]Users, 0)
	query := repo.db.Table("users")

	result := query.Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}

	outData := make([]domain.User, 0)
	for _, value := range users {
		var user domain.User
		user.ID = value.ID
		user.FullName = value.FullName
		user.Email = value.Email
		user.Password = value.Password
		user.CreatedAt = value.CreatedAt
		user.UpdatedAt = value.UpdatedAt
		outData = append(outData, user)
	}

	return outData, nil
}

func mappingInput(user *domain.User) Users {
	var result Users

	id := uuid.New()

	timeNow := time.Now()
	result.ID = id.String()
	result.FullName = user.FullName
	result.Email = user.Email
	result.Password = user.Password
	result.CreatedAt = timeNow.String()

	return result
}
