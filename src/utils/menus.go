package utils

import (
	"fmt"
	"strings"

	"github.com/magikitty/menu"
)

// MainMenu calls functions for the main menu
func MainMenu() {
	fmt.Println(MenuMessages().MenuInstructions)
	menu.PrintMenu(MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(MenuMainOptions)
	callMenuFunctions(menuSelection)
}

func callMenuFunctions(menuSelection string) {
	switch menuSelection {
	case MenuMessages().MenuMainOptions["1"]:
		generateMealPlan()
	case MenuMessages().MenuMainOptions["2"]:
		addNewMeal()
		// Refactor case logic into own function
	case MenuMessages().MenuMainOptions["3"]:
		if getQuitInput() == MenuMessages().QuitYes {
			QuitProgram = true
			// Decide if this is for debugging or now
			fmt.Println("Quitting now. Bye bye!") // debugging
		} else {
			MainMenu()
		}
		// Remove
	default:
		fmt.Println("Nothing to do") // debugging
	}
}

func generateMealPlan() {
	// Rename duration variables to be consistent
	duration := GetMealPlanDuration()
	mealData := ReadDataFromFile(FilePaths().JSONMealsData)
	totalPortions := getTotalPortions(mealData)
	// Rename bool (valid)
	mealPlanDurationViable := checkDurationViability(duration, totalPortions)

	if mealPlanDurationViable == true {
		mealPlan := MakeMealPlan(duration, mealData)
		displayMealPlan(mealPlan)
	}
}

func addNewMeal() {
	fmt.Println("We are going to add a new meal! Yay!") // debugging
}

func getQuitInput() (userInput string) {
	fmt.Println(MenuMessages().ConfirmQuit)
	userInput = GetUserInput()
	return strings.ToLower(userInput)
}
