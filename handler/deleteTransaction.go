package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTransactionHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE /api/v1/transactions",
	})
}
