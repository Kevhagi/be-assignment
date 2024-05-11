package paymentmanagerroutes

import (
	config "be-assignment/configs"
	controllerImplement "be-assignment/controllers/implement"
	paymentmanagerdto "be-assignment/dtos/payment_manager"
	resultdto "be-assignment/dtos/result"
	"be-assignment/prisma/db"
	repositoryImplement "be-assignment/repositories/implement"
	serviceImplement "be-assignment/services/implement"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func TransactionRoutes(g *gin.Engine, gr *gin.RouterGroup, db *db.PrismaClient) {
	transactionRepository := repositoryImplement.RepositoryPaymentManager(db)
	userRepository := repositoryImplement.RepositoryAccountManager(db)

	transactionService := serviceImplement.ServicePaymentManager(transactionRepository, userRepository)
	transactionController := controllerImplement.ControllerPaymentManager(transactionService)

	gr.POST("/send", config.WrapVerifySession(nil), func(c *gin.Context) {
		transaction := new(paymentmanagerdto.TransactionSendRequestController)

		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(400, &resultdto.ErrorResultJSON{
				Status:  400,
				Message: err.Error(),
			})
			return
		}

		if err := validator.New().Struct(transaction); err != nil {
			print(err)
			c.JSON(400, &resultdto.ErrorResultJSON{
				Status:  400,
				Message: err.Error(),
			})
			return
		}

		transactionRequestData := &paymentmanagerdto.TransactionSendRequestController{
			DestinationAccountID: transaction.DestinationAccountID,
			Amount:               transaction.Amount,
			Currency:             transaction.Currency,
		}
		transactionData, err := transactionController.Send(c, *transactionRequestData)
		if err != nil {
			c.JSON(400, &resultdto.ErrorResultJSON{
				Status:  400,
				Message: err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Hellow Transaction Test",
			"data":    transactionData,
		})
	})

	gr.POST("/withdraw", config.WrapVerifySession(nil), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellow Transaction Test",
		})
	})
}
