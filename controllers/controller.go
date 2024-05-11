package controller

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	paymentmanagerdto "be-assignment/dtos/payment_manager"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(ctx *gin.Context) []accountmanagerdto.UserResponse
}

type TransactionController interface {
	Send(ctx *gin.Context, transaction paymentmanagerdto.TransactionSendRequestController) (paymentmanagerdto.TransactionSendResponseRepository, error)
}
