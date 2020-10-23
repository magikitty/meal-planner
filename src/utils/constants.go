package utils

import (
	"errors"
)

var Tab = "   "

// MenuMainOptions contains main menu options
var MenuMainOptions = map[string]string{
	"1": "Generate a meal plan",
	"2": "Add a new meal",
	"3": "Quit",
}

// FilePaths returns struct with file paths
func FilePaths() filePaths {
	filePaths := filePaths{}
	filePaths.JSONMealsData = "./data/meals/meals.json"
	return filePaths
}

// MenuMessages returns struct with menu messages
func MenuMessages() menuMessages {
	menuMessages := menuMessages{}
	menuMessages.ConfirmQuit = "Are you sure you want to quit? y/n"
	menuMessages.DisplayMealPlan = "\nHere is your meal plan:\n"
	menuMessages.DisplayPlanFormatting = "Day %v: %v\n%v%v\n"
	menuMessages.InputNotValid = "Invalid input, please try again."
	menuMessages.MealPlanDuration = "How many days do you want to create a meal plan for?"
	menuMessages.MenuInstructions = "\nWhat do you want to do?\n"
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.QuitYes = "y"
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	return menuMessages
}

// MessagesHome returns strings for home page
func MessagesHome() messagesHome {
	Messages := messagesHome{}
	Messages.WelcomeMessage = MenuMessages().WelcomeMessage
	Messages.Instructions = MenuMessages().MenuInstructions
	Messages.OptionNewPlan = MenuMainOptions["1"]
	Messages.OptionNewRecipe = MenuMainOptions["2"]
	return Messages
}

// PageAddresses contains all page addresses used in e.g. messages and links
func PageAddresses() pageAddress {
	Address := pageAddress{}
	Address.Home = "/"
	Address.NewPlan = "/new-plan"
	Address.NewRecipe = "/new-recipe"
	return Address
}

// UXHomeConstants values used to display any content to user
func UXHomeConstants() uXHomeConstants {
	u := uXHomeConstants{}
	u.AddressHome = PageAddresses().Home
	u.AddressNewPlan = PageAddresses().NewPlan
	u.AddressNewRecipe = PageAddresses().NewRecipe
	u.MessageWelcome = MessagesHome().WelcomeMessage
	u.MessageInstructions = MessagesHome().Instructions
	u.ButtonOptionsNewPlan = MessagesHome().OptionNewPlan
	u.ButtonOptionsNewRecipe = MessagesHome().OptionNewRecipe
	return u
}

func UXNewRecipeConstants() uxNewRecipeConstants {
	u := uxNewRecipeConstants{}
	u.AddressHome = PageAddresses().Home
	u.AddressNewPlan = PageAddresses().NewPlan
	u.AddressNewRecipe = PageAddresses().NewRecipe
	u.MessageNewRecipe = "We are going to add a new meal! Yay!"
	return u
}

// ErrInvalidIndexToRemove returns error
var ErrInvalidIndexToRemove = errors.New("cannot remove item as index invalid")

type filePaths struct {
	JSONMealsData string
}

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

// GetFormatStrings initialization function returns map of string formatting values
func GetFormatStrings() map[string]string {
	var formatStrings = map[string]string{
		"day":           "Day",
		"portionSize":   "Portion size",
		"space":         " ",
		"colon":         ":",
		"bracketOpen":   "(",
		"bracketClosed": ")",
	}
	return formatStrings
}

type messagesHome struct {
	WelcomeMessage  string
	Instructions    string
	OptionNewPlan   string
	OptionNewRecipe string
}

type pageAddress struct {
	Home      string
	NewPlan   string
	NewRecipe string
}

type uXHomeConstants struct {
	AddressHome            string
	AddressNewPlan         string
	AddressNewRecipe       string
	MessageWelcome         string
	MessageInstructions    string
	ButtonOptionsNewPlan   string
	ButtonOptionsNewRecipe string
}

type uxNewRecipeConstants struct {
	AddressHome      string
	AddressNewPlan   string
	AddressNewRecipe string
	MessageNewRecipe string
}
