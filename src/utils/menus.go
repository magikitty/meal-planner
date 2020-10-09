package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/magikitty/menu"
)

// MenuMain displays and handles user input for main menu
func MenuMain() {
	fmt.Println(MenuMessages().MenuInstructions)
	menu.PrintMenu(MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(MenuMainOptions)
	menuMainSelection(menuSelection)
}

func menuMainSelection(menuSelection string) {
	switch menuSelection {
	case MenuMessages().MenuMainOptions["1"]:
		generateDisplayMealPlan()
	case MenuMessages().MenuMainOptions["2"]:
		addNewMeal()
	case MenuMessages().MenuMainOptions["3"]:
		quit()
	}
}

func generateDisplayMealPlan() {
	duration := getMealPlanDuration()
	mealData := ReadDataFromFile(FilePaths().JSONMealsData)
	sumAllPortions := getSumAllMealsPortions(mealData)
	durationValid := checkDurationValid(duration, sumAllPortions)

	if durationValid == true {
		mealPlan := makeMealPlan(duration, mealData.Meals)
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
	fmt.Println("We are going to add a new meal! Yay!")
}

func quit() {
	fmt.Println(MenuMessages().ConfirmQuit)
	userInput := menu.GetUserInput()
	strings.ToLower(userInput)

	if userInput == MenuMessages().QuitYes {
		os.Exit(0)
	}
}
