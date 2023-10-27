package house

import (
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/house

// @Summary Create a new house
// @Description Create a new house
// @Tags house
// @Accept json
// @Produce json
// @Param house body CreateHouseRequest true "Request body"
// @Success 200 {object} handler.CreateResponse
// @Failure 400,404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/house [post]
func CreateHouseHandler(context *gin.Context) {
	request := CreateHouseRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	house := schemas.House{
		Terms: request.Terms,
	}

	if err := handler.DB.Create(&house).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, "error creating house")
		return
	}

	handler.SendSuccess(context, "create house", house)
}
