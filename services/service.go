package service

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	paymentmanagerdto "be-assignment/dtos/payment_manager"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUsers(ctx *gin.Context) ([]accountmanagerdto.UserResponse, error)
}

type TransactionService interface {
	Send(ctx *gin.Context, transaction paymentmanagerdto.TransactionSendRequestController) (paymentmanagerdto.TransactionSendResponseRepository, error)
	// Withdraw(ctx context.Context, transaction paymentmanagerdto.TransactionRequest) error
}
