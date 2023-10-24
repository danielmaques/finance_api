package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllTransactionsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET /api/v1/transactions",
	})
}
