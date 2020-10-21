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

func FormatMealPlan(mealPlan []Meal, err error) []FormattedMeal {
	var formattedMealPlan []FormattedMeal

	if err != nil {
		fmt.Println(err)
	} else {
		var name, ingredients string
		var portionSize, portionsLeft int

		for i, meal := range mealPlan {
			name, ingredients, portionSize = prettifyMealProperties(meal)
			//fmt.Println("\nday: ", i, "\nname: ", name, "\ningredients: ", ingredients, "\nportion size: ", portionSize)

			if portionsLeft == 0 {
				name, ingredients, portionSize = prettifyMealProperties(meal)
				portionsLeft = portionSize
			} else if portionsLeft != portionSize {
				ingredients = ""
			}

			formattedMealPlan = append(formattedMealPlan,
				FormattedMeal{
				strconv.Itoa(i),
				name,
				ingredients,
				strconv.Itoa(portionSize)})
			portionsLeft--
		}
	}

	return formattedMealPlan
}

func displayMealPlan(mealPlan []Meal, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(utils.MenuMessages().DisplayMealPlan)
		var name, ingredients string
		var portionSize, portionsLeft int

		for i, meal := range mealPlan {
			name, ingredients, portionSize = prettifyMealProperties(meal)
			//fmt.Println("\nday: ", i, "\nname: ", name, "\ningredients: ", ingredients, "\nportion size: ", portionSize)

			if portionsLeft == 0 {
				name, ingredients, portionSize = prettifyMealProperties(meal)
				portionsLeft = portionSize
			} else if portionsLeft != portionSize {
				ingredients = ""
			}

			printMealInPlan(i, name, ingredients, portionSize)
			portionsLeft--
		}
	}
}

func printMealInPlan(i int, name string, ingredients string, portionSize int) {
	fmt.Printf(
		utils.MenuMessages().DisplayPlanFormatting,
		i+1, name,
		utils.Tab, formatIngredients(ingredients, portionSize))
}

func formatIngredients(ingredients string, portionSize int) string {
	if ingredients == "" {
		return ""
	}
	return fmt.Sprintf("Ingredients:\n%v\n%vPortions size: %v\n",
		ingredients, utils.Tab, strconv.Itoa(portionSize))
}

func quit() {
	fmt.Println(utils.MenuMessages().ConfirmQuit)
	userInput := menu.GetUserInput()
	strings.ToLower(userInput)

	if userInput == utils.MenuMessages().QuitYes {
		os.Exit(0)
	}
}
