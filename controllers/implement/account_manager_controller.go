package controllerimplement

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	service "be-assignment/services"

	"github.com/gin-gonic/gin"
)

type ControllerImplement struct {
	UserService service.UserService
}

func ControllerUser(
	UserService service.UserService,
) service.UserService {
	return &ControllerImplement{UserService}
}

func (s *ControllerImplement) GetUsers(ctx *gin.Context) ([]accountmanagerdto.UserResponse, error) {
	users, err := s.UserService.GetUsers(ctx)
	if err != nil {
		print("HOHOHOHOHO")
		print(err)
	}
	return users, err
}
