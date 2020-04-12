package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type input struct {
	Name string `"json.name"`
}

type Hello struct {
	Status  string `json: "status"`
	Message string `json: "message, omitempty"`
}

// SayHello says hello at ping
func SayHello(w http.ResponseWriter, r *http.Request) {
	newInput := input{}

	err := json.NewDecoder(r.Body).Decode(&newInput)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(newInput)

	resp := Hello{
		Status:  "OK",
		Message: "Hello " + newInput.Name,
	}

	fmt.Println(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}
