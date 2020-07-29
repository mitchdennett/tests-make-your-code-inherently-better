package better

type RecipeService interface {
	ListRecipes(page int) ([]*Recipe, error)
}

type Recipe struct {
	ID    int
	Title string
}
