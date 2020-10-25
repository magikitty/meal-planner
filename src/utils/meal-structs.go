package utils

// AllMeals struct represnts unmarshalled JSON data of all meals
type AllMeals struct {
	Meals []Meal
}

// Meal struct for unmarhshalling JSON meals
type Meal struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	PortionSize int          `json:"portionSize"`
}

// MealStringified is string-representation meals for front end presentation
type MealStringified struct {
	DayNumber   string
	Name        string
	Ingredients []string
	PortionSize string
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
