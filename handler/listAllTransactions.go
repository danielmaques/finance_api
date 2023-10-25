package handler

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

func ListAllTransactionsHandler(context *gin.Context) {
	id := context.Query("id")

	transactions := []schemas.Transaction{}

	if id != "" {
		if err := db.Find(&transactions, "id = ?", id).Error; err != nil {
			sendError(context, http.StatusNotFound, fmt.Sprintf("transaction with id %s not found", id))
			return
		}
	}else if err := db.Find(&transactions).Error; err != nil{
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error listing transactions: %v", err))
		return
	}

	sendSuccess(context, "list transactions", transactions)
}
