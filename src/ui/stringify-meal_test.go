package ui

import (
	"testing"

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
	},
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
	expectedMeal := utils.MealStringified{
		DayNumber:   "Day 1: ",
		Name:        "Potato Delight",
		Ingredients: []string{"1 potato"},
		PortionSize: "Portion size: 1",
	}
	actualMeal := stringifiedMeal(0, meal.Name, stringifiedIngredients(meal), meal.PortionSize)

	assert.Equal(t, expectedMeal, actualMeal,
		"Meal was not properly stringified.")
}
