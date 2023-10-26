package router

import (
	"github.com/danielmaques/finance_api/docs"
	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/handler/transaction"
	"github.com/danielmaques/finance_api/handler/user"
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
		v1.POST("/transactions", transaction.CreateTransactionHandler)
		v1.GET("/transactions", transaction.ListAllTransactionsHandler)
		v1.DELETE("/transactions", transaction.DeleteTransactionHandler)
		v1.PUT("/transactions", transaction.UpdateTransactionHandler)

		v1.GET("/user", user.GetUserHandler)
		v1.DELETE("/user", user.DeleteUserHandler)
		v1.POST("/user", user.CreateUserHandler)
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
