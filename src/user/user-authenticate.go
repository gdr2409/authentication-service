package user

import (
	"authentication"
	"encoding/json"
	Error "myerrors"
	"net/http"
)

type authenticateInput struct {
	Token string `json: "authtoken"`
	ID    uint   `json: "id, omitempty"`
}

type authenticateResponse struct {
	// ID    uint   `json: "id, omitempty"`
	// Name  string `json: "name"`
	// Age   string `json: "age"`
	// Email string `json: "email"`
	Status bool `json: "status"`
}

// Authenticate checks user auth
func Authenticate(w http.ResponseWriter, r *http.Request) *Error.Exception {
	input := authenticateInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Invalid JSON schema",
			Reason:  "INVALID_JSON_SCHEMA",
		}
	}

	authModule := authentication.AuthModule{}

	data, catch := authModule.CheckAuthentication(input.ID, input.Token, "user")

	if data == false || catch != nil {
		return &Error.Exception{
			Code:    400,
			Message: "Unauthorized",
			Reason:  "WRONG_AUTHENTICATION",
		}
	}

	resp := authenticateResponse{
		Status: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return nil
}
