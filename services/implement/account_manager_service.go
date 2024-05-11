package serviceimplement

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	repository "be-assignment/repositories"
	service "be-assignment/services"

	"github.com/gin-gonic/gin"
)

type ServiceImplement struct {
	UserRepository repository.UserRepository
}

func ServiceUser(
	UserRepository repository.UserRepository,
) service.UserService {
	return &ServiceImplement{UserRepository}
}

func (s *ServiceImplement) GetUsers(ctx *gin.Context) ([]accountmanagerdto.UserResponse, error) {
	users, err := s.UserRepository.GetUsers(ctx)
	if err != nil {
		print("HOKHOKHOKHOKHOK")
		print(err)
	}
	return users, err
}
