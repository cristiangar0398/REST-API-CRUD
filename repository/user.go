package repository

import (
	"context"

	"github.com/cristiangar0398/REST-API-CRUD/models"
)

var (
	implementation UserRepository
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
}

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}
