package serviceimplement

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	repository "be-assignment/repositories"
	service "be-assignment/services"

	"github.com/gin-gonic/gin"
)

type AccountManagerServiceImplement struct {
	UserRepository repository.UserRepository
}

func ServiceAccountManager(
	UserRepository repository.UserRepository,
) service.UserService {
	return &AccountManagerServiceImplement{UserRepository}
}

func (s *AccountManagerServiceImplement) GetUsers(ctx *gin.Context) ([]accountmanagerdto.UserResponse, error) {
	users, err := s.UserRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}
