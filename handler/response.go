package handler

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"mensage": fmt.Sprintf("%s success", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type CreateResponse struct {
	Message string                      `json:"message"`
	Data    schemas.TransactionResponse `json:"data"`
}
