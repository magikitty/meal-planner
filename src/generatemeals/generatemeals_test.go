package generatemeals

import (
	"testing"

	"github.com/magikitty/meal-planner/src/ui"
	"github.com/magikitty/meal-planner/src/utils"
	"github.com/stretchr/testify/assert"
)

var allMeals = utils.AllMeals{
	Meals: []utils.Meal{
		utils.Meal{
			Name: "Potato Delight",
			Ingredients: []utils.Ingredient{
				utils.Ingredient{
					Name:     "potato",
					Quantity: 1,
					Unit: &utils.IngredientUnit{
						Name:     "",
						Quantity: 0,
						Unit:     "",
					},
				},
			},
			PortionSize: 1,
		},
		utils.Meal{
			Name: "Pomato",
			Ingredients: []utils.Ingredient{
				{
					Name:     "tomato chunks",
					Quantity: 2,
					Unit: &utils.IngredientUnit{
						Name:     "can",
						Quantity: 400,
						Unit:     "gram",
					},
				},
				{
					Name:     "potato",
					Quantity: 10,
					Unit: &utils.IngredientUnit{
						Name:     "",
						Quantity: 0,
						Unit:     "",
					},
				},
				{
					Name:     "water",
					Quantity: 1,
					Unit: &utils.IngredientUnit{
						Name:     "litre",
						Quantity: 0,
						Unit:     "",
					},
				},
			},
			PortionSize: 4,
		},
	},
}

func TestGenerateMealPlan_addAllPortionsOfMeal(t *testing.T) {
	meal := allMeals.Meals[1]
	expectedMealSlice := []utils.Meal{meal, meal, meal, meal}
	actualMealSlice := addAllPortionsOfMeal(meal)

	assert.Equal(t, expectedMealSlice, actualMealSlice,
		"Meal has 4 portions, so returned meal slice should contain meal 4 times.")
}

func TestGenerateMealPlan_removeMeal(t *testing.T) {
	meal1 := allMeals.Meals[0]
	meal2 := allMeals.Meals[1]
	mealSlice := []utils.Meal{meal1, meal2}

	expectedMealSlice, _ := []utils.Meal{meal1}, error(nil)
	actualMealSlice, actualErr := removeMeal(mealSlice, 1)

	assert.Equal(t, expectedMealSlice, actualMealSlice,
		"Meal should have been removed from meals and new slice returned without meal.")
	assert.Nil(t, actualErr,
		"Error returned as meal could not be removed - error should have been nil.")
}

func TestGenerateMealPlan_makeMealPlan(t *testing.T) {
	actualMealPlan := makeMealPlan(11, "../../data/tests/test-meals.json")
	expectedMealPlanLength := 11
	actualMealPlanLength := len(actualMealPlan)

	assert.Equal(t, expectedMealPlanLength, actualMealPlanLength,
		"Meal plan was not the right length. makeMealPlan should have returned a meal plan for 11 days.")
}

func TestParseMeals_StringifyMealPlan(t *testing.T) {
	expectedPlan := []utils.MealStringified{
		utils.MealStringified{
			DayNumber:   "Day 1: ",
			Name:        "Potato Delight",
			Ingredients: []string{"1 potato"},
			PortionSize: "Portion size: 1",
		},
		utils.MealStringified{
			DayNumber: "Day 2: ",
			Name:      "Pomato",
			Ingredients: []string{
				"2 can (400 gram) tomato chunks",
				"10 potato",
				"1 litre water",
			},
			PortionSize: "Portion size: 4",
		},
		utils.MealStringified{
			DayNumber:   "Day 3: ",
			Name:        "Pomato",
			Ingredients: []string{},
			PortionSize: "Portion size: 4",
		},
		utils.MealStringified{
			DayNumber:   "Day 4: ",
			Name:        "Pomato",
			Ingredients: []string{},
			PortionSize: "Portion size: 4",
		},
		utils.MealStringified{
			DayNumber:   "Day 5: ",
			Name:        "Pomato",
			Ingredients: []string{},
			PortionSize: "Portion size: 4",
		},
	}
	actualPlan, _ := ui.StringifyMealPlan([]utils.Meal{allMeals.Meals[0], allMeals.Meals[1], allMeals.Meals[1], allMeals.Meals[1], allMeals.Meals[1]}, nil)

	assert.Equal(t, expectedPlan, actualPlan,
		"Plan was not properly stringified.")
}
