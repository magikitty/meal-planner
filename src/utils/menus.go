package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/magikitty/menu"
)

// MainMenu displays and handles user input for main menu
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
		mealPlan := MakeMealPlan(duration, mealData.Meals)
		displayMealPlan(mealPlan)
	}
}

func displayMealPlan(mealPlan []Meal) {
	fmt.Println(MenuMessages().DisplayMealPlan)
	for i, meal := range mealPlan {
		fmt.Printf(MenuMessages().DisplayPlanFormatting, i+1, meal.Name, meal.Ingredients, meal.Portions)
	}
}

func addNewMeal() {
	fmt.Println("We are going to add a new meal! Yay!") // debugging
}

func quit() {
	if getQuitInput() == MenuMessages().QuitYes {
		os.Exit(0)
	}
}

func getQuitInput() (userInput string) {
	fmt.Println(MenuMessages().ConfirmQuit)
	userInput = menu.GetUserInput()
	return strings.ToLower(userInput)
}
