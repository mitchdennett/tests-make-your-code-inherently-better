package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ListRecipes(e *Env, w http.ResponseWriter, r *http.Request) {
	pages, ok := r.URL.Query()["page"]

	if !ok || len(pages[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pages[0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	itemsList, err := e.RecipeService.ListRecipes(page)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	retJSON, err := json.Marshal(itemsList)
	fmt.Fprintf(w, string(retJSON))
}
