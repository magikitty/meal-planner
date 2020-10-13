package generatemeals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMeals_mealFitsPlan(t *testing.T) {
	meal := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	expectedBoolFits := true
	gotBoolFits := mealFitsPlan(meal, 5)

	expectedBoolDoesNotFit := false
	gotBoolDoesNotFit := mealFitsPlan(meal, 2)

	assert.Equal(t, expectedBoolFits, gotBoolFits, "Meal portions should be less than or equal to duration")
	assert.Equal(t, expectedBoolDoesNotFit, gotBoolDoesNotFit, "Meal portions should be less than or equal to duration")
}

func TestGenerateMeals_durationValid(t *testing.T) {
	expectedDurationValidBool := true
	gotDurationValidBool := durationValid(8)

	assert.Equal(t, expectedDurationValidBool, gotDurationValidBool, "Duration needs to be more than 0")
}

func TestGenerateMeals_addAllPortionsOfMeal(t *testing.T) {
	soup := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	expectedMealSlice := []Meal{soup, soup, soup, soup}
	actualMealSlice := addAllPortionsOfMeal(soup)

	assert.Equal(t, expectedMealSlice, actualMealSlice, "Meal slice should have number of meals equal to meal's portion property length")
}

func TestGenerateMeals_removeMeal(t *testing.T) {
	meal1 := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	meal2 := Meal{"Curry", []string{"curry, cauliflower"}, 1}
	meal3 := Meal{"Salad", []string{"lettuce, tomato"}, 1}
	meals := []Meal{meal1, meal2, meal3}

	expectedMealSlice := []Meal{meal1, meal3}
	actualMealSlice := removeMeal(meals, 1)

	assert.Equal(t, expectedMealSlice, actualMealSlice, "Meal at given index position should be removed")
}

// PROBLEM: can't find json file with meals
func TestGenerateMeals_makeMealPlan(t *testing.T) {
	mealPlanExpectedLength := 15
	mealPlanGenerated := makeMealPlan(15, "../../data/meals/meals.json")
	mealPlanGeneratedLength := len(mealPlanGenerated)

	assert.Equal(t, mealPlanExpectedLength, mealPlanGeneratedLength, "Meal plan should be same length as duration")
}
