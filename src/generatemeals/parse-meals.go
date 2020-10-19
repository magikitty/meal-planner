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
	var ingredientString string
	for _, ingredient := range meal.Ingredients {
		if ingredient.Unit == nil {
			ingredientString = utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " + ingredient.Name
			ingredients = append(ingredients, ingredientString)
		} else if ingredient.Unit != nil {
			if ingredient.Unit.Unit != "" {
				ingredientString = utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " +
					ingredient.Unit.Name + " (" + strconv.Itoa(ingredient.Unit.Quantity) + " " +
					ingredient.Unit.Unit + ") " + ingredient.Name
				ingredients = append(ingredients, ingredientString)
			} else {
				ingredientString = utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " +
					ingredient.Unit.Name + " " + ingredient.Name
				ingredients = append(ingredients, ingredientString)
			}
		}
	}
	return strings.Join(ingredients, "\n")
}
