package generatemeals

// AllMeals contains new to-be-implemented meals (below)
type AllMeals struct {
	Meals []Meal
}

// Meal to be implemented once new Meal data representation sorted out
type Meal struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	PortionSize int          `json:"portionSize"`
}

// Ingredient is a struct of ingredients for a meal
type Ingredient struct {
	Name     string          `json:"name"`
	Quantity int             `json:"quantity"`
	Unit     *IngredientUnit `json:"unit,omitempty"`
}

// IngredientUnit is a struct representing information about ingredients' units
type IngredientUnit struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity,omitempty"`
	Unit     string `json:"unit,omitempty"`
}
