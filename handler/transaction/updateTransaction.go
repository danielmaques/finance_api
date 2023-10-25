package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTransactionHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT /api/v1/transactions",
	})
}
