package transaction

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/transactions

// @Summary Get all transactions
// @Description Get a list of all transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Success 200 {array} schemas.TransactionResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/transactions [get]
func ListAllTransactionsHandler(context *gin.Context) {
	id := context.Query("id")

	transactions := []schemas.Transaction{}

	if id != "" {
		if err := handler.DB.Find(&transactions, "id = ?", id).Error; err != nil {
			handler.SendError(context, http.StatusNotFound, fmt.Sprintf("transaction with id %s not found", id))
			return
		}
	} else if err := handler.DB.Find(&transactions).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error listing transactions: %v", err))
		return
	}

	handler.SendSuccess(context, "list transactions", transactions)
}
