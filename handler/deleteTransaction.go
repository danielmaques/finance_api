package handler

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/transactions

// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param id query string true "Transaction ID"
// @Success 200 {object} CreateResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/transactions [delete]
func DeleteTransactionHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errParamsRequired("id", "string").Error())
		return
	}

	transaction := schemas.Transaction{}

	if err := db.First(&transaction, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("transaction with id %s not found", id))
		return
	}

	if err := db.Delete(&transaction).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting transaction with id %s", id))
		return
	}

	sendSuccess(context, "delete transaction", transaction)
}
