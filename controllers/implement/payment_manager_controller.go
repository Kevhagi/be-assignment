package controllerimplement

import (
	paymentmanagerdto "be-assignment/dtos/payment_manager"
	resultdto "be-assignment/dtos/result"
	service "be-assignment/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PaymentManagerControllerImplement struct {
	TransactionService service.TransactionService
}

func ControllerPaymentManager(
	TransactionService service.TransactionService,
) service.TransactionService {
	return &PaymentManagerControllerImplement{TransactionService}
}

func (s *PaymentManagerControllerImplement) Send(ctx *gin.Context, transaction paymentmanagerdto.TransactionSendRequestController) (paymentmanagerdto.TransactionSendResponseRepository, error) {
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(400, &resultdto.ErrorResultJSON{
			Status:  400,
			Message: err.Error(),
		})
		return paymentmanagerdto.TransactionSendResponseRepository{}, err
	}

	if err := validator.New().Struct(transaction); err != nil {
		ctx.JSON(400, &resultdto.ErrorResultJSON{
			Status:  400,
			Message: err.Error(),
		})
		return paymentmanagerdto.TransactionSendResponseRepository{}, err
	}

	return s.TransactionService.Send(ctx, transaction)
}
