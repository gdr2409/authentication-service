package authentication

import (
	"fmt"
	"mydatabase"
	Error "myerrors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var (
	authRepo *gorm.DB
)

// SetupModel creates db connection
func SetupModel() {
	authRepo = mydatabase.Initialize()
	authRepo.AutoMigrate(&AuthToken{})
}

// AuthModule act as container class for authentication
type AuthModule struct{}

// Create new auth token
func (auth AuthModule) Create(id uint, objectType string) (string, *Error.Exception) {
	token := uuid.New()

	tokenToCreate := &AuthToken{
		Token:      token.String(),
		ExpiryAt:   time.Now().AddDate(0, 0, 1),
		ObjectId:   fmt.Sprint(id),
		ObjectType: objectType,
	}

	authRepo.Create(tokenToCreate)

	return tokenToCreate.Token, nil
}
