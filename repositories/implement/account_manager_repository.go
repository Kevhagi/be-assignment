package repositoryimplement

import (
	"be-assignment/prisma/db"
	repository "be-assignment/repositories"

	accountmanagerdto "be-assignment/dtos/account_manager"
	"context"
)

type RepositoryImplement struct {
	DB *db.PrismaClient
}

func RepositoryUser(DB *db.PrismaClient) repository.UserRepository {
	return &RepositoryImplement{DB}
}

func (r *RepositoryImplement) GetUsers(ctx context.Context) ([]accountmanagerdto.UserResponse, error) {
	allUser, err := r.DB.User.FindMany().Exec(ctx)
	if err != nil {
		print("")
	}

	var users []accountmanagerdto.UserResponse
	for _, user := range allUser {
		users = append(users, accountmanagerdto.UserResponse{
			Id:                user.ID,
			Email:             user.Email,
			SupertokensUserId: user.SupertokensUserID,
		})
	}

	return users, err
}
