package generatemeals

import (
	"github.com/magikitty/meal-planner/src/utils"
	"strconv"
	"strings"
)

func getMealParametersAsString(meal Meal) (string, string, string) {
	name := meal.Name
	ingredients := getIngredientsAsString(meal)
	portionSize := strconv.Itoa(meal.PortionSize)

	return name, ingredients, portionSize
}

func getIngredientsAsString(meal Meal) string {
	var ingredients []string
	var ingredientString string

	for _, ingredient := range meal.Ingredients {
		 if ingredient.Unit != nil {
			 ingredientString = getIngredientUnitsAsString(ingredient)
			 ingredients = append(ingredients, ingredientString)
			} else {
			ingredientString = utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " + ingredient.Name
			ingredients = append(ingredients, ingredientString)
		}
	}
	return strings.Join(ingredients, "\n")
}

func getIngredientUnitsAsString(ingredient Ingredient) string {
	var ingredientString string

	if ingredient.Unit.Unit != "" && ingredient.Unit.Quantity != 0 {
		ingredientString = utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " +
			ingredient.Unit.Name + " (" + strconv.Itoa(ingredient.Unit.Quantity) + " " +
			ingredient.Unit.Unit + ") " + ingredient.Name
	} else {
		ingredientString = utils.Tab + utils.Tab + strconv.Itoa(ingredient.Quantity) + " " +
			ingredient.Unit.Name + " " + ingredient.Name
	}
	return ingredientString
}
