package utils

import (
	"testing"

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

func TestHelpers_ConvertStringToInt(t *testing.T) {
	intExpected := 32
	intFromStringSucceess := ConvertStringToInt("32")

	assert.Equal(t, intExpected, intFromStringSucceess, "Failed to convert number string to int")
}

func TestHelpers_GetFileData(t *testing.T) {
	byteArrayExpected := []byte{0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa, 0x7d}
	byteArrayActual := GetFileData("../../data/tests/test-getfiledata.json")

	assert.Equal(t, byteArrayExpected, byteArrayActual, "Failed to get expected byte array from file")
}

func TestHelpers_IsLargerThanZero(t *testing.T) {
	expectedDurationValidBool := true
	actualDurationValidBool := IsLargerThanZero(8)

	expectedDurationInvalidBoolZero := false
	actualDurationInvalidBoolZero := IsLargerThanZero(0)

	expectedDurationInvalidBoolNegative := false
	actualDurationInvalidBoolNegative := IsLargerThanZero(-10)

	assert.Equal(t, expectedDurationValidBool, actualDurationValidBool,
		"Duration of 8 is valid and should have returned true.")
	assert.Equal(t, expectedDurationInvalidBoolZero, actualDurationInvalidBoolZero,
		"Duration of 0 is an invalid duration and should have returned false. Duration has to be more than 0.")
	assert.Equal(t, expectedDurationInvalidBoolNegative, actualDurationInvalidBoolNegative,
		"Duration of -5 is invalid and should have returned false. Duration cannot be a negative number.")
}

func TestGenerateMealPlan_NumberFitsLimit(t *testing.T) {
	expectedNumberFits := true
	actualNumberFits := NumberFitsLimit(4, 4)

	expectedNumberDoesNotFit := true
	actualNumberDoesNotFitTooLarge := NumberFitsLimit(4, 3)

	actualBoolDoesNotFitUnderZero := NumberFitsLimit(-1, 3)

	assert.Equal(t, expectedNumberFits, actualNumberFits,
		"4 Is equal to 4, so it should have passed.")
	assert.NotEqual(t, expectedNumberDoesNotFit, actualNumberDoesNotFitTooLarge,
		"4 is more than 3, so it should have failed.")
	assert.NotEqual(t, expectedNumberDoesNotFit, actualBoolDoesNotFitUnderZero,
		"-1 is below 0, so it should have failed.")
}

func TestGenerateMealPlan_getMealsData(t *testing.T) {
	mealsData := GetFileData("../../data/tests/test-meals.json")
	actualMealsAll := GetAllMealsFromData(mealsData)
	assert.Equal(t, allMeals, actualMealsAll, "Failed to get expected meals.")
}
