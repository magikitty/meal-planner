package generatemeals

import (
	"fmt"
	"os"
	"strings"

	"github.com/magikitty/meal-planner/src/utils"
	"github.com/magikitty/menu"
)

// MenuMain displays and handles user input for main menu
func MenuMain() {
	fmt.Println(utils.MenuMessages().MenuInstructions)
	menu.PrintMenu(utils.MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuMainOptions)
	menuMainSelection(menuSelection)
}

func menuMainSelection(menuSelection string) {
	switch menuSelection {
	case utils.MenuMessages().MenuMainOptions["1"]:
		displayMealPlan(getMealPlan())
	case utils.MenuMessages().MenuMainOptions["2"]:
		addNewMeal()
	case utils.MenuMessages().MenuMainOptions["3"]:
		quit()
	}
}

func displayMealPlan(mealPlan []Meal, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(utils.MenuMessages().DisplayMealPlan)
		for i, meal := range mealPlan {
			// TODO: call function to handle display ingredients to user properly
			fmt.Printf(
				utils.MenuMessages().DisplayPlanFormatting,
				i+1,
				meal.Name,
				utils.Tab,
				getIngredientsAsString(meal),
				utils.Tab,
				meal.PortionSize,
			)
		}
	}
}

// TODO: Print strings returned from parse-meals

func quit() {
	fmt.Println(utils.MenuMessages().ConfirmQuit)
	userInput := menu.GetUserInput()
	strings.ToLower(userInput)

	if userInput == utils.MenuMessages().QuitYes {
		os.Exit(0)
	}
}
