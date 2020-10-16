package generatemeals

// AllMeals is a struct containing all meals
type AllMeals struct {
	Meals []Meal
}

// Meal is a struct for an individual meal
type Meal struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	PortionSize int      `json:"portionSize"`
}
