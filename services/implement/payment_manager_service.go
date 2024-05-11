package serviceimplement

import (
	paymentmanagerdto "be-assignment/dtos/payment_manager"
	"be-assignment/pkg"
	repository "be-assignment/repositories"
	service "be-assignment/services"

	"github.com/gin-gonic/gin"
)

type PaymentManagerServiceImplement struct {
	TransactionRepository repository.TransactionRepository
	UserRepository        repository.UserRepository
}

func ServicePaymentManager(
	TransactionRepository repository.TransactionRepository,
	UserRepository repository.UserRepository,
) service.TransactionService {
	return &PaymentManagerServiceImplement{
		TransactionRepository,
		UserRepository,
	}
}

func (s *PaymentManagerServiceImplement) Send(ctx *gin.Context, transaction paymentmanagerdto.TransactionSendRequestController) (paymentmanagerdto.TransactionSendResponseRepository, error) {
	cookie, err := ctx.Cookie("sAccessToken")
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return paymentmanagerdto.TransactionSendResponseRepository{}, err
	}

	claims, err := pkg.DecodeJWT(cookie)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return paymentmanagerdto.TransactionSendResponseRepository{}, err
	}

	sourceAccountId, err := s.UserRepository.AccountIdByUserId(ctx, claims.Sub)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return paymentmanagerdto.TransactionSendResponseRepository{}, err
	}

	transactionSendRequest := paymentmanagerdto.TransactionSendRequestRepository{
		SourceAccountID:      sourceAccountId,
		DestinationAccountID: transaction.DestinationAccountID,
		Amount:               transaction.Amount,
		Currency:             transaction.Currency,
	}

	transactionData, err := s.TransactionRepository.Send(ctx, transactionSendRequest)
	if err != nil {
		return paymentmanagerdto.TransactionSendResponseRepository{}, err
	}
	return transactionData, err
}
