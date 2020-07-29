package main

import (
	"log"
	"net/http"

	"github.com/mitchdennett/tests-make-your-code-inherently-better/database"
	"github.com/mitchdennett/tests-make-your-code-inherently-better/handler"
)

func main() {

	db := database.SetupDatabase()
	store := database.DB{DB: db}

	http.Handle("/recipes", handler.Handler{RecipeService: &store})

	// Logs the error if ListenAndServe fails.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
