package handler

import (
	"net/http"

	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/transactions

// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body CreateTransactionRequest true "Request body"
// @Success 200 {object} CreateResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/transactions [post]
func CreateTransactionHandler(context *gin.Context) {
	request := CreateTransactionRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	transaction := schemas.Transaction{
		Add:         request.Add,
		Category:    request.Category,
		Description: request.Description,
		Amount:      request.Amount,
		Date:        request.Date,
	}

	if err := db.Create(&transaction).Error; err != nil {
		logger.Errorf("error creating transaction %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating transaction")
		return
	}

	sendSuccess(context, "create transaction", transaction)

}
