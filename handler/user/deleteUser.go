package user

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/user

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} handler.CreateResponse
// @Failure 400,404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/user/{id} [delete]
func DeleteUserHandler(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, errParamsUserRequired("id", "string").Error())
		return
	}

	user := schemas.User{}

	if err := handler.DB.First(&user, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("user with id %s not found", id))
		return
	}

	if err := handler.DB.Delete(&user).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id %s", id))
		return
	}

	handler.SendSuccess(context, "delete user", user)
}
