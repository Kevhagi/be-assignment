package repositoryimplement

import (
	"be-assignment/prisma/db"
	repository "be-assignment/repositories"

	accountmanagerdto "be-assignment/dtos/account_manager"
	"context"
)

type AccountManagerRepositoryImplement struct {
	DB *db.PrismaClient
}

func RepositoryAccountManager(DB *db.PrismaClient) repository.UserRepository {
	return &AccountManagerRepositoryImplement{DB}
}

func (r *AccountManagerRepositoryImplement) GetUsers(ctx context.Context) ([]accountmanagerdto.UserResponse, error) {
	allUser, err := r.DB.User.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
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

func (r *AccountManagerRepositoryImplement) AccountIdByUserId(ctx context.Context, userId string) (string, error) {
	account, err := r.DB.Account.FindFirst(
		db.Account.UserID.Equals(userId),
		db.Account.Type.Equals("Debit"),
	).Select(
		"id",
	).Exec(ctx)
	if err != nil {
		return "", err
	}
	return account.ID, nil
}
