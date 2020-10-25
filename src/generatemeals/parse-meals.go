package generatemeals

import (
	"strconv"

	"github.com/magikitty/meal-planner/src/utils"
)

// StringifyMealPlan returns struct of stringified mealplans
func StringifyMealPlan(mealPlan []utils.Meal, err error) ([]utils.MealStringified, error) {
	// TODO: Throw new custom error defined in constants file instead
	if err != nil {
		return nil, err
	}

	var stringifiedMealPlan []utils.MealStringified
	var name string
	var ingredients []string
	var portionSize, portionsLeft int

	for i, meal := range mealPlan {
		if portionsLeft == 0 {
			name = meal.Name
			ingredients = stringifiedIngredients(meal)
			portionSize = meal.PortionSize
			portionsLeft = portionSize
		} else if portionsLeft != portionSize {
			ingredients = []string{}
		}

		stringifiedMealPlan = append(stringifiedMealPlan, stringifiedMeal(i, name, ingredients, portionSize))
		portionsLeft--
	}

	return stringifiedMealPlan, nil
}

func stringifiedMeal(i int, name string, ingredients []string, portionSize int) utils.MealStringified {
	stringifiedMeal := utils.MealStringified{
		utils.GetFormatStrings()["day"] + utils.GetFormatStrings()["space"] + strconv.Itoa(i+1) +
			utils.GetFormatStrings()["colon"] + utils.GetFormatStrings()["space"],
		name,
		ingredients,
		utils.GetFormatStrings()["portionSize"] + utils.GetFormatStrings()["colon"] +
			utils.GetFormatStrings()["space"] + strconv.Itoa(portionSize)}

	return stringifiedMeal
}

func stringifiedIngredients(meal utils.Meal) []string {
	var ingredients []string
	var ingredientString string

	for _, ingredient := range meal.Ingredients {
		if ingredient.Unit.Name != "" {
			ingredientString = getIngredientUnitsAsString(ingredient)
			ingredients = append(ingredients, ingredientString)
		} else {
			// Format ingredient for recipe string
			ingredientString = strconv.Itoa(ingredient.Quantity) + utils.GetFormatStrings()["space"] + ingredient.Name
			ingredients = append(ingredients, ingredientString)
		}
	}
	return ingredients
}

func getIngredientUnitsAsString(ingredient utils.Ingredient) string {
	var ingredientString string

	// Format ingredient for recipe string
	if ingredient.Unit.Unit != "" && ingredient.Unit.Quantity != 0 {
		ingredientString =
			strconv.Itoa(ingredient.Quantity) +
				utils.GetFormatStrings()["space"] +
				ingredient.Unit.Name +
				utils.GetFormatStrings()["space"] + utils.GetFormatStrings()["bracketOpen"] +
				strconv.Itoa(ingredient.Unit.Quantity) +
				utils.GetFormatStrings()["space"] +
				ingredient.Unit.Unit +
				utils.GetFormatStrings()["bracketClosed"] + utils.GetFormatStrings()["space"] +
				ingredient.Name
	} else {
		ingredientString =
			strconv.Itoa(ingredient.Quantity) +
				utils.GetFormatStrings()["space"] +
				ingredient.Unit.Name +
				utils.GetFormatStrings()["space"] +
				ingredient.Name
	}
	return ingredientString
}
