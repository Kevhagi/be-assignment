package repository

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	"context"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]accountmanagerdto.UserResponse, error)
}
