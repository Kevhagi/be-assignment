package controller

import (
	accountmanagerdto "be-assignment/dtos/account_manager"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(ctx *gin.Context) []accountmanagerdto.UserResponse
}
