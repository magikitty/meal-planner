package generatemeals

import (
	"testing"

	"github.com/magikitty/meal-planner/src/utils"
	"github.com/stretchr/testify/assert"
)

var allMeals = AllMeals{
	Meals: []Meal{
		Meal{
			Name: "Potato Delight",
			Ingredients: []Ingredient{
				Ingredient{
					Name:     "potato",
					Quantity: 1,
					Unit: &IngredientUnit{
						Name:     "",
						Quantity: 0,
						Unit:     "",
					},
				},
			},
			PortionSize: 1,
		},
		Meal{
			Name: "Pomato",
			Ingredients: []Ingredient{
				{
					Name:     "tomato chunks",
					Quantity: 2,
					Unit: &IngredientUnit{
						Name:     "can",
						Quantity: 400,
						Unit:     "gram",
					},
				},
				{
					Name:     "potato",
					Quantity: 10,
					Unit: &IngredientUnit{
						Name:     "",
						Quantity: 0,
						Unit:     "",
					},
				},
				{
					Name:     "water",
					Quantity: 1,
					Unit: &IngredientUnit{
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

func TestGenerateMealPlan_getMealsData(t *testing.T) {
	mealsData := utils.GetFileData("../../data/tests/test-meals.json")
	actualMealsAll := getAllMealsFromData(mealsData)
	assert.Equal(t, allMeals, actualMealsAll, "Failed to get expected meals.")
}

func TestGenerateMealPlan_mealFitsPlan(t *testing.T) {
	meal := allMeals.Meals[1]
	expectedBoolFits := true
	actualBoolFits := mealFitsPlan(meal, 4)

	expectedBoolDoesNotFit := false
	actualBoolDoesNotFit := mealFitsPlan(meal, 3)

	assert.Equal(t, expectedBoolFits, actualBoolFits,
		"Meal has 4 portions and fits into the meal plan with a duration of 4, so it should have returned true.")
	assert.Equal(t, expectedBoolDoesNotFit, actualBoolDoesNotFit,
		"Meal has 4 portions and does not fit into the meal plan with a duration of 3, so it should have returned false.")
}

func TestGenerateMealPlan_durationValid(t *testing.T) {
	expectedDurationValidBool := true
	actualDurationValidBool := durationValid(8)

	expectedDurationInvalidBoolZero := false
	actualDurationInvalidBoolZero := durationValid(0)

	expectedDurationInvalidBoolNegative := false
	actualDurationInvalidBoolNegative := durationValid(-10)

	assert.Equal(t, expectedDurationValidBool, actualDurationValidBool,
		"Duration of 8 is valid and should have returned true.")
	assert.Equal(t, expectedDurationInvalidBoolZero, actualDurationInvalidBoolZero,
		"Duration of 0 is an invalid duration and should have returned false. Duration has to be more than 0.")
	assert.Equal(t, expectedDurationInvalidBoolNegative, actualDurationInvalidBoolNegative,
		"Duration of -5 is invalid and should have returned false. Duration cannot be a negative number.")
}

func TestGenerateMealPlan_addAllPortionsOfMeal(t *testing.T) {
	meal := allMeals.Meals[1]
	expectedMealSlice := []Meal{meal, meal, meal, meal}
	actualMealSlice := addAllPortionsOfMeal(meal)

	assert.Equal(t, expectedMealSlice, actualMealSlice,
		"Meal has 4 portions, so returned meal slice should contain meal 4 times.")
}

func TestGenerateMealPlan_removeMeal(t *testing.T) {
	meal1 := allMeals.Meals[0]
	meal2 := allMeals.Meals[1]
	mealSlice := []Meal{meal1, meal2}

	expectedMealSlice, _ := []Meal{meal1}, error(nil)
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

func TestParseMeals_StringifiedIngredients(t *testing.T) {
	meal := allMeals.Meals[0]

	expectedIngredients := []string{"1 potato"}
	actualIngredients := stringifiedIngredients(meal)

	assert.Equal(t, expectedIngredients, actualIngredients,
		"Ingredients were not properly stringified.")
}

func TestParseMeals_StringifyMeal(t *testing.T) {
	meal := allMeals.Meals[0]
	expectedMeal := MealStringified{
		DayNumber:   "Day 1: ",
		Name:        "Potato Delight",
		Ingredients: []string{"1 potato"},
		PortionSize: "Portion size: 1",
	}
	actualMeal := stringifiedMeal(0, meal.Name, stringifiedIngredients(meal), meal.PortionSize)

	assert.Equal(t, expectedMeal, actualMeal,
		"Meal was not properly stringified.")
}

func TestParseMeals_StringifyMealPlan(t *testing.T) {
	expectedPlan := []MealStringified{
		MealStringified{
			DayNumber:   "Day 1: ",
			Name:        "Potato Delight",
			Ingredients: []string{"1 potato"},
			PortionSize: "Portion size: 1",
		},
		MealStringified{
			DayNumber: "Day 2: ",
			Name:      "Pomato",
			Ingredients: []string{
				"2 can (400 gram) tomato chunks",
				"10 potato",
				"1 litre water",
			},
			PortionSize: "Portion size: 4",
		},
		MealStringified{
			DayNumber:   "Day 3: ",
			Name:        "Pomato",
			Ingredients: []string{},
			PortionSize: "Portion size: 4",
		},
		MealStringified{
			DayNumber:   "Day 4: ",
			Name:        "Pomato",
			Ingredients: []string{},
			PortionSize: "Portion size: 4",
		},
		MealStringified{
			DayNumber:   "Day 5: ",
			Name:        "Pomato",
			Ingredients: []string{},
			PortionSize: "Portion size: 4",
		},
	}
	actualPlan, _ := StringifyMealPlan([]Meal{allMeals.Meals[0], allMeals.Meals[1], allMeals.Meals[1], allMeals.Meals[1], allMeals.Meals[1]}, nil)

	assert.Equal(t, expectedPlan, actualPlan,
		"Plan was not properly stringified.")
}
