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
		displayMealPlanNew(getMealPlan())
	case MenuMessages().MenuMainOptions["2"]:
		addNewMeal()
	case MenuMessages().MenuMainOptions["3"]:
		quit()
	}
}

func displayMealPlanNew(mealPlan []Meal, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(MenuMessages().DisplayMealPlan)
		for i, meal := range mealPlan {
			fmt.Printf(MenuMessages().DisplayPlanFormatting, i+1, meal.Name, meal.Ingredients, meal.Portions)
		}
	}
}

func quit() {
	fmt.Println(MenuMessages().ConfirmQuit)
	userInput := menu.GetUserInput()
	strings.ToLower(userInput)

	if userInput == MenuMessages().QuitYes {
		os.Exit(0)
	}
}
