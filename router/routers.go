package router

import (
	"github.com/danielmaques/finance_api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/transactions", handler.GetAllTransactionsHandler)
		v1.GET("/transactions/{id}", handler.GetTransactionHandler)
		v1.POST("/transactions", handler.CreateTransactionHandler)
		v1.DELETE("/transactions/{id}", handler.DeleteTransactionHandler)
		v1.PUT("/transactions", handler.UpdateTransactionHandler)
	}
}
