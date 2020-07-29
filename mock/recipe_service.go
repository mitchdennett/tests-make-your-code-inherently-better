package mock

import (
	better "github.com/mitchdennett/tests-make-your-code-inherently-better"
)

type MockRecipeService struct {
	ListRecipesFunc func(page int) ([]*better.Recipe, error)
}

func (s *MockRecipeService) ListRecipes(page int) ([]*better.Recipe, error) {
	return s.ListRecipesFunc(page)
}
