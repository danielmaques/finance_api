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

func GetAllTransactionsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET /api/v1/transactions",
	})
}

func GetTransactionHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET /api/v1/transactions",
	})
}

func DeleteTransactionHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE /api/v1/transactions",
	})
}

func UpdateTransactionHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT /api/v1/transactions",
	})
}
