package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	better "github.com/mitchdennett/tests-make-your-code-inherently-better"
	"github.com/mitchdennett/tests-make-your-code-inherently-better/mock"
)

func TestListRecipes(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/recipes?page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	var store mock.MockRecipeService
	store.ListRecipesFunc = func(page int) ([]*better.Recipe, error) {
		return []*better.Recipe{{ID: 1, Title: "Pasta"}}, nil
	}
	handler := Handler{RecipeService: &store}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
