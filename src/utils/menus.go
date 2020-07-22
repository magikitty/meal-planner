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
	case MenuMessages().MenuMainOptions["3"]:
		quit()
	}
}

func generateMealPlan() {
	duration := GetMealPlanDuration()
	mealData := ReadDataFromFile(FilePaths().JSONMealsData)
	sumAllPortions := getSumAllMealsPortions(mealData)
	durationValid := checkDurationValid(duration, sumAllPortions)

	if durationValid == true {
		mealPlan := MakeMealPlan(duration, mealData)
		displayMealPlan(mealPlan)
	}
}

func addNewMeal() {
	fmt.Println("We are going to add a new meal! Yay!") // debugging
}

func quit() {
	if getQuitInput() == MenuMessages().QuitYes {
		QuitProgram = true
	} else {
		MainMenu()
	}
}

func getQuitInput() (userInput string) {
	fmt.Println(MenuMessages().ConfirmQuit)
	userInput = GetUserInput()
	return strings.ToLower(userInput)
}
