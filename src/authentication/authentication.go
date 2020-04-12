package authentication

import (
	"cache"
	"encoding/json"
	"fmt"
	"mydatabase"
	Error "myerrors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var (
	authRepo *gorm.DB
	lru      *cache.LRU
)

// SetupModel creates db connection
func SetupModel() {
	authRepo = mydatabase.Initialize()
	authRepo.AutoMigrate(&AuthToken{})
	lru = cache.New(3)
}

type CachedAuthToken struct {
	Token      string    `json: "token"`
	ExpiryAt   time.Time `json: "expiry_at"`
	ObjectId   uint      `json: "object_id"`
	ObjectType string    `json: "object_type"`
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

// Delete old auth token
func (auth AuthModule) Delete(id uint, objectType string) (bool, *Error.Exception) {

	authRepo.Where("object_id = ? AND object_type = ?", id, objectType).Delete(&AuthToken{})

	return true, nil
}

// CheckAuthentication validates auth token
func (auth AuthModule) CheckAuthentication(id uint, token string, objectType string) (bool, *Error.Exception) {
	cachedData := lru.Get(id)

	if cachedData == "" {
		fmt.Println("No matching data in cache")
		existingToken := []AuthToken{}

		authRepo.Where("object_id = ? AND object_type = ? AND token = ? AND expiry_at > ?", id, objectType, token, time.Now()).First(&existingToken)

		if len(existingToken) > 0 {
			// Save this data to cache
			authToken, err := json.Marshal(CachedAuthToken{
				Token:      existingToken[0].Token,
				ExpiryAt:   existingToken[0].ExpiryAt,
				ObjectId:   id,
				ObjectType: existingToken[0].ObjectType,
			})

			if err != nil {
				fmt.Println("Error while caching data")
			} else {
				lru.Put(id, string(authToken))
			}

			return true, nil
		}
	} else {
		// Validate token from cache
		fmt.Println("Found data in cache")
		authToken := &CachedAuthToken{}
		err := json.Unmarshal([]byte(cachedData), authToken)

		fmt.Println(authToken)

		if err != nil {
			return false, &Error.Exception{
				Code:    400,
				Message: "Error while retrieving from cache",
				Reason:  "CACHE_ERROR",
			}
		}

		if authToken.Token != token ||
			authToken.ObjectId != id ||
			authToken.ObjectType != objectType ||
			authToken.ExpiryAt.Before(time.Now()) {
			return false, &Error.Exception{
				Code:    400,
				Message: "Error while retrieving from cache",
				Reason:  "CACHE_ERROR",
			}
		}

		return true, nil
	}

	return false, nil
}
