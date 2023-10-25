package handler

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

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
