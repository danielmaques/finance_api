package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET /api/v1/user",
	})
}
