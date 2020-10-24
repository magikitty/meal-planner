package generatemeals

import (
	"fmt"
	"os"
	"strconv"
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
		//displayMealPlan(getMealPlan())
		fmt.Print(FormatMealPlan(getMealPlan()))
	case utils.MenuMessages().MenuMainOptions["2"]:
		addNewMeal()
	case utils.MenuMessages().MenuMainOptions["3"]:
		quit()
	}
}

func FormatMealPlan(mealPlan []Meal, err error) []MealStringified {
	var formattedMealPlan []MealStringified

	if err != nil {
		fmt.Println(err)
	} else {
		var name string
		var ingredients []string
		var portionSize, portionsLeft int

		for i, meal := range mealPlan {

			if portionsLeft == 0 {
				name = meal.Name
				ingredients = getIngredientsAsString(meal)
				portionSize = meal.PortionSize
				portionsLeft = portionSize
			} else if portionsLeft != portionSize {
				ingredients = []string{}
			}

			formattedMealPlan = append(formattedMealPlan, formattedMeal(i, name, ingredients, portionSize))
			portionsLeft--
		}
	}

	return formattedMealPlan
}

func formattedMeal(i int, name string, ingredients []string, portionSize int) MealStringified {
	formattedMeal := MealStringified{
		"Day " + strconv.Itoa(i+1) + ":",
		name,
		ingredients,
		"Portion size: " + strconv.Itoa(portionSize)}

	return formattedMeal
}

func quit() {
	fmt.Println(utils.MenuMessages().ConfirmQuit)
	userInput := menu.GetUserInput()
	strings.ToLower(userInput)

	if userInput == utils.MenuMessages().QuitYes {
		os.Exit(0)
	}
}
