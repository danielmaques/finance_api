package transaction

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/transactions

// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} CreateResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/transactions/{id} [delete]
func DeleteTransactionHandler(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, errParamsTransactionRequired("id", "string").Error())
		return
	}

	transaction := schemas.Transaction{}

	if err := handler.DB.First(&transaction, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("transaction with id %s not found", id))
		return
	}

	if err := handler.DB.Delete(&transaction).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting transaction with id %s", id))
		return
	}

	handler.SendSuccess(context, "delete transaction", transaction)
}
