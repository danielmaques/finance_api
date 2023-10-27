package house

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateInviteCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func isInviteCodeUnique(inviteCode string) bool {
	var count int64
	handler.DB.Model(&schemas.House{}).Where("invite_code = ?", inviteCode).Count(&count)
	return count == 0
}

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

	for i, user := range request.Users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			handler.Logger.Errorf("error hashing password for user %s: %v", user.Email, err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error hashing password")
			return
		}
		request.Users[i].Password = string(hashedPassword)
	}

	// Gerar um InviteCode único
	inviteCode := generateInviteCode(5)
	for !isInviteCodeUnique(inviteCode) {
		inviteCode = generateInviteCode(5)
	}

	house := schemas.House{
		InviteCode: inviteCode,
		Users:      request.Users,
	}

	if err := handler.DB.Create(&house).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, "error creating house")
		return
	}

	handler.SendSuccess(context, "create house", house)
}
