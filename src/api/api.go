package api

import (
	"encoding/json"
	"hello"
	"log"
	Error "myerrors"
	"net/http"
	"user"
)

// Routes will have serve mux to include all routes
type Routes struct {
}

type rootHandler func(http.ResponseWriter, *http.Request) *Error.Exception

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}

// Register routes
func (r *Routes) Register() *http.ServeMux {
	sm := http.NewServeMux()

	sm.Handle("/ping", http.HandlerFunc(hello.SayHello))
	sm.Handle("/signUp", rootHandler(user.SignUp))
	sm.Handle("/login", rootHandler(user.Login))

	return sm
}
