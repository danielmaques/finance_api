package router

import (
	"github.com/danielmaques/finance_api/docs"
	"github.com/danielmaques/finance_api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group("/api/v1")
	{
		v1.GET("/transactions", handler.ListAllTransactionsHandler)
		v1.POST("/transactions", handler.CreateTransactionHandler)
		v1.DELETE("/transactions", handler.DeleteTransactionHandler)
		v1.PUT("/transactions", handler.UpdateTransactionHandler)
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
