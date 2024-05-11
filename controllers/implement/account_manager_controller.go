package controllerimplement

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	service "be-assignment/services"

	"github.com/gin-gonic/gin"
)

type AccountManagerControllerImplement struct {
	UserService service.UserService
}

func ControllerAccountManager(
	UserService service.UserService,
) service.UserService {
	return &AccountManagerControllerImplement{UserService}
}

func (s *AccountManagerControllerImplement) GetUsers(ctx *gin.Context) ([]accountmanagerdto.UserResponse, error) {
	users, err := s.UserService.GetUsers(ctx)
	if err != nil {
		print("HOHOHOHOHO")
		print(err)
	}
	return users, err
}
