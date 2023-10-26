package user

import (
	"net/http"

	"github.com/danielmaques/finance_api/handler"
	"github.com/danielmaques/finance_api/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1/user

// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "Request body"
// @Success 200 {object} handler.CreateResponse
// @Failure 400,404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/v1/user [post]
func CreateUserHandler(context *gin.Context) {
	request := CreateUserRequest{}

	if err := context.BindJSON(&request); err != nil {
		handler.Logger.Errorf("error binding request %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, "Invalid request format")
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("error validating request %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		handler.Logger.Errorf("error hashing password %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error hashing password")
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := handler.DB.Create(&user).Error; err != nil {
		handler.Logger.Errorf("error creating user %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error creating user")
		return
	}

	// Não retorne a senha na resposta
	user.Password = ""

	handler.SendSuccess(context, "create user", user)
}
