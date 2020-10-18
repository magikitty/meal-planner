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
		displayMealPlan(GetMealPlan())
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
			fmt.Printf(utils.MenuMessages().DisplayPlanFormatting, i+1, meal.Name, meal.Ingredients, meal.PortionSize)
		}
	}
}

// StringMealPlan to return string representation of meal plan for frontend
func StringMealPlan(mealPlan []Meal) (string, error) {

	var mealPlanString []string
	var mealString string
	mealPlanString = append(mealPlanString, utils.MenuMessages().DisplayMealPlan)
	for i, meal := range mealPlan {
		mealString = fmt.Sprintf(utils.MenuMessages().DisplayPlanFormatting, i+1, meal.Name, meal.Ingredients, meal.PortionSize)
		mealPlanString = append(mealPlanString, mealString)
	}
	return strings.Join(mealPlanString, "\n"), nil
}

func quit() {
	fmt.Println(utils.MenuMessages().ConfirmQuit)
	userInput := menu.GetUserInput()
	strings.ToLower(userInput)

	if userInput == utils.MenuMessages().QuitYes {
		os.Exit(0)
	}
}
