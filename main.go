package main

import (
	"fmt"
	"meal-planner/utils"

	"github.com/magikitty/menu"
)

func main() {
	fmt.Println(utils.WelcomeMessage)
	mainMenu()
}

func mainMenu() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuMainOptions)
	callMenuFunctions(menuSelection)
}

func callMenuFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuMainOptions["1"]:
		generateMealPlan()
	case utils.MenuMainOptions["2"]:
		addNewMeal()
	case utils.MenuMainOptions["3"]:
		getQuitInput()
	default:
		fmt.Println("Nothing to do") // debugging
	}
}

func generateMealPlan() {
	fmt.Println("We are going to generate a meal plan! Yay!") // debugging
}

func addNewMeal() {
	fmt.Println("We are going to add a new meal! Yay!") // debugging
}

func getQuitInput() {
	fmt.Println(utils.ConfirmQuit)
}
