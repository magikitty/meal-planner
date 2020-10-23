package utils

import "errors"

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
	menuMessages.DisplayMealPlan = "\nHere is your meal plan:"
	menuMessages.DisplayPlanFormatting = "Day %v: %v\n   Ingredients: %v\n   Portion size: %v \n"
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

// UXConstants values used to display any content to user
func UXConstants() uxConstants {
	u := uxConstants{}
	u.AddressHome = PageAddresses().Home
	u.AddressNewPlan = PageAddresses().NewPlan
	u.AddressNewRecipe = PageAddresses().NewRecipe
	u.HomeMessageWelcome = MessagesHome().WelcomeMessage
	u.HomeMessageInstructions = MessagesHome().Instructions
	u.HomeButtonOptionsNewPlan = MessagesHome().OptionNewPlan
	u.HomeButtonOptionsNewRecipe = MessagesHome().OptionNewRecipe
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

type uxConstants struct {
	AddressHome                string
	AddressNewPlan             string
	AddressNewRecipe           string
	HomeMessageWelcome         string
	HomeMessageInstructions    string
	HomeButtonOptionsNewPlan   string
	HomeButtonOptionsNewRecipe string
}
