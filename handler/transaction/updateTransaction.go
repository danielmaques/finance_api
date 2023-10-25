package transaction

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/transactions

// @Summary Update a transaction
// @Description Update a transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} handler.CreateResponse
// @Failure 400,404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/transactions/{id} [put]
func UpdateTransactionHandler(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, errParamsTransactionRequired("id", "string").Error())
		return
	}

	var updatedTransaction schemas.Transaction
	if err := context.ShouldBindJSON(&updatedTransaction); err != nil {
		handler.SendError(context, http.StatusBadRequest, "Invalid request body")
		return
	}

	transaction := schemas.Transaction{}
	if err := handler.DB.First(&transaction, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("transaction with id %s not found", id))
		return
	}

	if err := handler.DB.Model(&transaction).Updates(updatedTransaction).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error updating transaction with id %s", id))
		return
	}

	handler.SendSuccess(context, "update transaction", transaction)
}
