package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE /api/v1/user",
	})
}