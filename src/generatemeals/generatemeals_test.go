package generatemeals

import (
	"github.com/magikitty/meal-planner/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMealPlan_mealFitsPlan(t *testing.T) {
	meal := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	expectedBoolFits := true
	actualBoolFits := mealFitsPlan(meal, 4)

	expectedBoolDoesNotFit := false
	actualBoolDoesNotFit := mealFitsPlan(meal, 3)

	assert.Equal(t, expectedBoolFits, actualBoolFits,
		"Soup has 4 portions and fits into the meal plan with a duration of 4, so it should have returned true.")
	assert.Equal(t, expectedBoolDoesNotFit, actualBoolDoesNotFit,
		"Soup has 4 portions and does not fit into the meal plan with a duration of 3, so it should have returned false.")
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
	soup := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	expectedMealSlice := []Meal{soup, soup, soup, soup}
	actualMealSlice := addAllPortionsOfMeal(soup)

	assert.Equal(t, expectedMealSlice, actualMealSlice,
		"Soup has 4 portions, so returned meal slice should contain soup 4 times.")
}

func TestGenerateMealPlan_removeMeal(t *testing.T) {
	meal1 := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	meal2 := Meal{"Curry", []string{"curry, cauliflower"}, 1}
	meal3 := Meal{"Salad", []string{"lettuce, tomato"}, 1}
	mealSlice1 := []Meal{meal1, meal2, meal3}
	mealSlice2 := []Meal{meal1, meal2, meal3}

	expectedMealSlice, expectedErr := []Meal{meal1, meal3}, error(nil)
	actualMealSlice, actualErr := removeMeal(mealSlice1, 1)

	expectedMealSliceInvalidInput, expectedInvalidError := mealSlice2, utils.InvalidIndexToRemove
	actualMealSliceInvalidInput, actualInvalidError := removeMeal(mealSlice2, 3)

	assert.Equal(t, expectedMealSlice, actualMealSlice,
		"Curry should have been removed from meals and new slice returned without curry.")
	assert.Equal(t, expectedErr, actualErr,
		"Error should be nil. Curry should have been removed from meals and new slice returned without curry.")
	assert.Equal(t, expectedMealSliceInvalidInput, actualMealSliceInvalidInput,
		"Tried to remove a meal at invalid index 3. It should have returned the original slice without removing anything.")
	assert.Equal(t, expectedInvalidError, actualInvalidError,
		"Invalid input of 3 for item index should have returned InvalidIndexToRemove error.")
}

func TestGenerateMealPlan_makeMealPlan(t *testing.T) {
	actualMealPlan := makeMealPlan(11, "../../data/tests/test-make-meal-plan.json")
	expectedMealPlanLength := 11
	actualMealPlanLength := len(actualMealPlan)

	assert.Equal(t, expectedMealPlanLength, actualMealPlanLength,
		"Meal plan was not the right length. makeMealPlan should have returned a meal plan for 11 days.")
}
