package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "POST /api/v1/transactions",
	})
}
