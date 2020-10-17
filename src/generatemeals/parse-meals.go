package generatemeals

import (
	"github.com/magikitty/meal-planner/src/utils"
	"strconv"
	"strings"
)

func getMealParametersAsString(meal Meal) (string, string, string) {
	name := meal.Name
	ingredients := getMealIngredientsAsString(meal)
	portionSize := strconv.Itoa(meal.PortionSize)

	return name, ingredients, portionSize
}

func getMealIngredientsAsString(meal Meal) string {
	var ingredients []string
	for _, ingredient := range meal.Ingredients {
		ingredients = append(ingredients, utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " + ingredient.Name)
	}
	return strings.Join(ingredients, "\n")
}
