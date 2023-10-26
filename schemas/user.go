package schemas

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type UserResponse struct {
	ID           uint          `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	Transactions []Transaction `json:"transactions"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
}

// HashPassword cria um hash da senha usando bcrypt
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword verifica se a senha fornecida corresponde ao hash armazenado
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
