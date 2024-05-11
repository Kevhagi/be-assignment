package repository

import (
	accountmanagerdto "be-assignment/dtos/account_manager"
	paymentmanagerdto "be-assignment/dtos/payment_manager"
	"context"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]accountmanagerdto.UserResponse, error)
	AccountIdByUserId(ctx context.Context, userId string) (string, error)
}

type TransactionRepository interface {
	Send(ctx *gin.Context, transaction paymentmanagerdto.TransactionSendRequestRepository) (paymentmanagerdto.TransactionSendResponseRepository, error)
	// Withdraw(ctx context.Context, transaction paymentmanagerdto.TransactionRequest) error
}
