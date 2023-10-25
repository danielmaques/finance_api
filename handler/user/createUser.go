package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "POST /api/v1/user",
	})
}
