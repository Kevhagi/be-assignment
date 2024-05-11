package service

import (
	accountmanagerdto "be-assignment/dtos/account_manager"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUsers(ctx *gin.Context) ([]accountmanagerdto.UserResponse, error)
}
