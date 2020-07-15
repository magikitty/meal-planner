package utils

import (
	"fmt"
	"menu"
	"strings"
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
		if getQuitInput() == MenuMessages().QuitYes {
			QuitProgram = true
			fmt.Println("Quitting now. Bye bye!") // debugging
		} else {
			MainMenu()
		}
	default:
		fmt.Println("Nothing to do") // debugging
	}
}

func generateMealPlan() {
	fmt.Println("We are going to generate a meal plan! Yay!") // debugging
	GetMealPlanDuration()
}

func addNewMeal() {
	fmt.Println("We are going to add a new meal! Yay!") // debugging
}

// simplify by looping while userInput == "", and returning whatever user inputs
func getQuitInput() (userInput string) {
	for userInput != MenuMessages().QuitYes && userInput != MenuMessages().QuitNo {
		fmt.Println(MenuMessages().ConfirmQuit)
		userInput = GetUserInput()
	}
	return strings.ToLower(userInput)
}
