package house

import (
	"fmt"
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetHouseHandler(context *gin.Context) {
	id := context.Query("id")

	house := schemas.House{}

	if err := handler.DB.Preload("Transactions").Preload("Users").First(&house, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			handler.SendError(context, http.StatusNotFound, fmt.Sprintf("house with id %s not found", id))
		} else {
			handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error getting house: %v", err))
		}
		return
	}

	handler.SendSuccess(context, "get house", house)
}
