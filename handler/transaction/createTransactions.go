package transaction

import (
	"net/http"

	"github.com/danielmaques/finance_api/handler"
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
// @Success 200 {object} handler.CreateResponse
// @Failure 400,404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/transactions [post]
func CreateTransactionHandler(context *gin.Context) {
	request := CreateTransactionRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("error validating request %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	transaction := schemas.Transaction{
		UserID:      request.UserID,
		Add:         request.Add,
		Category:    request.Category,
		Description: request.Description,
		Amount:      request.Amount,
		Date:        request.Date,
	}

	if err := handler.DB.Create(&transaction).Error; err != nil {
		handler.Logger.Errorf("error creating transaction %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error creating transaction")
		return
	}

	handler.SendSuccess(context, "create transaction", transaction)

}
