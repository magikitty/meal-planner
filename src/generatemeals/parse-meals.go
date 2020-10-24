package generatemeals

import (
	"fmt"
	"strconv"
)

func FormatMealPlan(mealPlan []Meal, err error) []MealStringified {
	var formattedMealPlan []MealStringified

	if err != nil {
		fmt.Println(err)
	} else {
		var name string
		var ingredients []string
		var portionSize, portionsLeft int

		for i, meal := range mealPlan {

			if portionsLeft == 0 {
				name = meal.Name
				ingredients = getIngredientsAsString(meal)
				portionSize = meal.PortionSize
				portionsLeft = portionSize
			} else if portionsLeft != portionSize {
				ingredients = []string{}
			}

			formattedMealPlan = append(formattedMealPlan, formattedMeal(i, name, ingredients, portionSize))
			portionsLeft--
		}
	}

	return formattedMealPlan
}

func formattedMeal(i int, name string, ingredients []string, portionSize int) MealStringified {
	formattedMeal := MealStringified{
		"Day " + strconv.Itoa(i+1) + ":",
		name,
		ingredients,
		"Portion size: " + strconv.Itoa(portionSize)}

	return formattedMeal
}

func getIngredientsAsString(meal Meal) []string {
	var ingredients []string
	var ingredientString string

	for _, ingredient := range meal.Ingredients {
		if ingredient.Unit.Name != "" {
			ingredientString = getIngredientUnitsAsString(ingredient)
			ingredients = append(ingredients, ingredientString)
		} else {
			// Format ingredient for recipe string
			ingredientString = strconv.Itoa(ingredient.Quantity) + " " + ingredient.Name
			ingredients = append(ingredients, ingredientString)
		}
	}
	return ingredients
}

func getIngredientUnitsAsString(ingredient Ingredient) string {
	var ingredientString string

	// Format ingredient for recipe string
	if ingredient.Unit.Unit != "" && ingredient.Unit.Quantity != 0 {
		ingredientString = strconv.Itoa(ingredient.Quantity) + " " +
			ingredient.Unit.Name + " (" + strconv.Itoa(ingredient.Unit.Quantity) + " " +
			ingredient.Unit.Unit + ") " + ingredient.Name
	} else {
		ingredientString = strconv.Itoa(ingredient.Quantity) + " " +
			ingredient.Unit.Name + " " + ingredient.Name
	}
	return ingredientString
}
