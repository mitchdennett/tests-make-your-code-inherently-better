package handler

import (
	"net/http"

	better "github.com/mitchdennett/tests-make-your-code-inherently-better"
)

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	RecipeService better.RecipeService
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handleListRecipes(w, r)
}
