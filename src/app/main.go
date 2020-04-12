package main

import (
	"api"
	"authentication"
	"fmt"
	"mydatabase"
	"net/http"
	"user"
)

func main() {
	fmt.Printf("Hello there\n")

	mydatabase.Initialize()
	user.SetupModel()
	authentication.SetupModel()

	routes := api.Routes{}
	sm := routes.Register()

	http.ListenAndServe(":3400", sm)
}
