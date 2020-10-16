package utils

import "errors"

// MenuMainOptions contains main menu options
var MenuMainOptions = map[string]string{
	"1": "Generate a meal plan",
	"2": "Add a new meal",
	"3": "Quit",
}

// MenuMessages returns struct with menu messages
func MenuMessages() menuMessages {
	menuMessages := menuMessages{}
	menuMessages.ConfirmQuit = "Are you sure you want to quit? y/n"
	menuMessages.DisplayMealPlan = "\nHere is your meal plan:"
	menuMessages.DisplayPlanFormatting = "Day %v: %v\n   Ingredients: %v\n   Portion size: %v \n"
	menuMessages.InputNotValid = "Invalid input, please try again."
	menuMessages.MealPlanDuration = "How many days do you want to create a meal plan for?"
	menuMessages.MenuInstructions = "\nWhat do you want to do? Press the number of your choice."
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.QuitYes = "y"
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	return menuMessages
}

// FilePaths returns struct with file paths
func FilePaths() filePaths {
	FilePaths := filePaths{}
	FilePaths.JSONMealsData = "./data/meals/meals.json"
	return FilePaths
}

// InvalidIndexToRemove returns error
var InvalidIndexToRemove = errors.New("Invalid index. Cannot remove item.")

type menuMessages struct {
	ConfirmQuit           string
	DisplayMealPlan       string
	DisplayPlanFormatting string
	InputNotValid         string
	MealPlanDuration      string
	MenuInstructions      string
	MenuMainOptions       map[string]string
	QuitYes               string
	WelcomeMessage        string
}

type filePaths struct {
	JSONMealsData string
}

// CustomErrors returns struct with custom errors
//func CustomErrors() customErrors {
//	customErrors := customErrors{}
//	customErrors.InvalidIndexToRemove = "Invalid index. Cannot remove item."
//	return customErrors
//}
//
//type customErrors struct {
//	InvalidIndexToRemove string
//}
