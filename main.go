package main

import (
	"fmt"
	"meal-planner/utils"
	"strings"

	"github.com/magikitty/menu"
)

func main() {
	fmt.Println(utils.WelcomeMessage)

	for utils.QuitProgram == false {
		mainMenu()
	}
}

// move this into new menus.go file
func mainMenu() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuMainOptions)
	callMenuFunctions(menuSelection)
}

// move this into new menus.go file
func callMenuFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuMainOptions["1"]:
		generateMealPlan()
	case utils.MenuMainOptions["2"]:
		addNewMeal()
	case utils.MenuMainOptions["3"]:
		if getQuitInput() == utils.QuitYes {
			utils.QuitProgram = true
			fmt.Println("Quitting now. Bye bye!") // debugging
		} else {
			mainMenu()
		}
	default:
		fmt.Println("Nothing to do") // debugging
	}
}

// move this into new menus.go file
func generateMealPlan() {
	fmt.Println("We are going to generate a meal plan! Yay!") // debugging
	utils.GetMealPlanDuration()
}

// move this into new menus.go file
func addNewMeal() {
	fmt.Println("We are going to add a new meal! Yay!") // debugging
}

// move this into new menus.go file
// simplify by looping while userInput == "", and returning whatever user inputs
func getQuitInput() (userInput string) {
	for userInput == utils.QuitYes && userInput != utils.QuitNo {
		fmt.Println(utils.ConfirmQuit)
		userInput = utils.GetUserInput()
	}
	return strings.ToLower(userInput)
}
