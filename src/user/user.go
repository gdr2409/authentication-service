package user

import (
	"authentication"
	"encoding/json"
	"mydatabase"
	Error "myerrors"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type signUpInput struct {
	Name     string `"json.name"`
	Age      int    `"json.age"`
	Email    string `"json.email"`
	Password string `"json.password"`
}

type loginInput struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type logoutInput struct {
	Email string `json: "email"`
}

type signUpResponse struct {
	ID        uint      `json: "id, omitempty"`
	Message   string    `json: "message"`
	CreatedAt time.Time `json: "createdAt, omitempty"`
}

type loginResponse struct {
	ID      uint   `json: "id, omitempty"`
	Token   string `json: "token, omitempty"`
	Message string `json: "message"`
}

type logoutResponse struct {
	Message string `json: "message"`
}

var (
	userRepository *gorm.DB
)

// SetupModel creates db connection
func SetupModel() {
	userRepository = mydatabase.Initialize()
	userRepository.AutoMigrate(&Person{})
}

// SignUp creates new row in Person table
func SignUp(w http.ResponseWriter, r *http.Request) *Error.Exception {
	input := signUpInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid JSON schema",
			Reason:  "INVALID_JSON_SCHEMA",
		}
	}

	existingUsers := []Person{}
	userRepository.Where("email = ?", input.Email).Find(&existingUsers)
	if len(existingUsers) > 0 {
		return &Error.Exception{
			Code:    400,
			Message: "Username already exists",
			Reason:  "USERNAME_ALREADY_EXISTS",
		}
	}

	userToCreate := &Person{
		Name:     input.Name,
		Age:      input.Age,
		Email:    input.Email,
		Password: input.Password,
	}

	userRepository.Create(userToCreate)

	resp := signUpResponse{
		ID:        userToCreate.Model.ID,
		CreatedAt: userToCreate.Model.CreatedAt,
		Message:   "User created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return nil
}

// Login and generates auth token
func Login(w http.ResponseWriter, r *http.Request) *Error.Exception {
	input := loginInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid JSON schema",
			Reason:  "INVALID_JSON_SCHEMA",
		}
	}

	existingUsers := []Person{}
	userRepository.Where("email = ?", input.Email).Find(&existingUsers)

	if len(existingUsers) == 0 {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid username",
			Reason:  "INVALID_USERNAME",
		}
	}

	existingUser := existingUsers[0]

	if input.Password != existingUser.Password {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid Password",
			Reason:  "INVALID_PASSWORD",
		}
	}

	authModule := authentication.AuthModule{}

	authToken, catch := authModule.Create(existingUser.Model.ID, "user")

	if catch != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Error while token creation",
			Reason:  "ERROR_TOKEN_CREATION",
		}
	}

	resp := loginResponse{
		Message: "Logged in successfully",
		Token:   authToken,
		ID:      existingUser.Model.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return nil
}

// Logout and generates auth token
func Logout(w http.ResponseWriter, r *http.Request) *Error.Exception {
	input := logoutInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid JSON schema",
			Reason:  "INVALID_JSON_SCHEMA",
		}
	}

	existingUsers := []Person{}
	userRepository.Where("email = ?", input.Email).First(&existingUsers)

	if len(existingUsers) == 0 {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid username",
			Reason:  "INVALID_USERNAME",
		}
	}

	existingUser := existingUsers[0]

	authModule := authentication.AuthModule{}

	_, catch := authModule.Delete(existingUser.Model.ID, "user")

	if catch != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Error while token deletion",
			Reason:  "ERROR_TOKEN_DELETION",
		}
	}

	resp := logoutResponse{
		Message: "Logged out successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return nil
}
