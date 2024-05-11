package repositoryimplement

import (
	"be-assignment/prisma/db"
	repository "be-assignment/repositories"

	paymentmanagerdto "be-assignment/dtos/payment_manager"

	resultdto "be-assignment/dtos/result"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/shopspring/decimal"
)

type PaymentManagerRepositoryImplement struct {
	DB *db.PrismaClient
}

func RepositoryPaymentManager(DB *db.PrismaClient) repository.TransactionRepository {
	return &PaymentManagerRepositoryImplement{DB}
}

func (r *PaymentManagerRepositoryImplement) Send(ctx *gin.Context, transaction paymentmanagerdto.TransactionSendRequestRepository) (paymentmanagerdto.TransactionSendResponseRepository, error) {
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

	// Process transaction
	r.DB.Account.FindUnique(
		db.Account.ID.Equals(transaction.SourceAccountID),
	).Update(
		db.Account.Balance.Decrement(decimal.NewFromFloat(transaction.Amount)),
	)

	r.DB.Account.FindUnique(
		db.Account.ID.Equals(transaction.DestinationAccountID),
	).Update(
		db.Account.Balance.Increment(decimal.NewFromFloat(transaction.Amount)),
	)

	// Process transaction recordings
	// * Create source transaction
	sourceTransaction, _ := r.DB.Transaction.CreateOne(
		db.Transaction.Type.Set("DEBIT"),
		db.Transaction.Currency.Set(transaction.Currency),
		db.Transaction.Status.Set("DONE"),
		db.Transaction.Source.Link(
			db.Account.ID.Equals(transaction.SourceAccountID),
		),
		db.Transaction.Destination.Link(
			db.Account.ID.Equals(transaction.DestinationAccountID),
		),
		db.Transaction.Amount.Set(decimal.NewFromFloat(transaction.Amount)),
	).Exec(ctx)

	// * Create destination transaction
	r.DB.Transaction.CreateOne(
		db.Transaction.Type.Set("CREDIT"),
		db.Transaction.Currency.Set(transaction.Currency),
		db.Transaction.Status.Set("DONE"),
		db.Transaction.Source.Link(
			db.Account.ID.Equals(transaction.DestinationAccountID),
		),
		db.Transaction.Destination.Link(
			db.Account.ID.Equals(transaction.SourceAccountID),
		),
		db.Transaction.Amount.Set(decimal.NewFromFloat(transaction.Amount)),
	).Exec(ctx)

	transactionSendResponse := paymentmanagerdto.TransactionSendResponseRepository{
		TransactionID:        sourceTransaction.ID,
		Timestamp:            sourceTransaction.Timestamp,
		SourceAccountID:      sourceTransaction.SourceID,
		DestinationAccountID: sourceTransaction.DestinationID,
	}

	return transactionSendResponse, nil
}
