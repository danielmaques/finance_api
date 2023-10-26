package user

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/user

// @Summary Get all user
// @Description Get a list of all user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} schemas.UserResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/user [get]
func GetUserHandler(context *gin.Context) {
	id := context.Query("id")

	user := schemas.User{}

	if id != "" {
		if err := handler.DB.Preload("Transactions").First(&user, "id = ?", id).Error; err != nil {
			handler.SendError(context, http.StatusNotFound, fmt.Sprintf("user with id %s not found", id))
			return
		}
	} else if err := handler.DB.Preload("Transactions").Find(&user).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error getting user: %v", err))
		return
	}

	handler.SendSuccess(context, "get user", user)
}
