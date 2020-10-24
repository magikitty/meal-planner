package utils

import "errors"

/*
Struct initialiser functions returning structs with page-specific values
	- Structs populated by values stored in maps
	- Structs passed to handlers' html.Execute functions for passing to
		pages' HTML
*/

// GetConstantsIndex returns struct with constant values for Home page
func GetConstantsIndex() ConstantsIndex {
	c := ConstantsIndex{}
	c.AddressIndex = PageAddresses()["index"]
	c.AddressNewPlan = PageAddresses()["newPlan"]
	c.AddressNewRecipe = PageAddresses()["newRecipe"]
	c.MessageWelcome = MessagesIndex()["welcome"]
	c.MessageInstructions = MessagesIndex()["instructions"]
	c.ButtonOptionsNewPlan = MessagesIndex()["optionNewPlan"]
	c.ButtonOptionsNewRecipe = MessagesIndex()["optionNewRecipe"]
	c.NavIndex = MessagesGlobal()["index"]
	c.TitlePage = MessagesGlobal()["index"]
	c.TitleApp = MessagesGlobal()["nameApp"]
	return c
}

// GetConstantsNewPlan returns struct with constant values for New Plan page
func GetConstantsNewPlan() ConstantsNewPlan {
	c := ConstantsNewPlan{}
	c.AddressIndex = PageAddresses()["index"]
	c.TitlePage = MessagesGlobal()["titleNewPlan"]
	c.TitleApp = MessagesGlobal()["nameApp"]
	return c
}

// GetConstantsNewRecipe returns struct with constant values for New Recipe page
func GetConstantsNewRecipe() ConstantsNewRecipe {
	c := ConstantsNewRecipe{}
	c.AddressIndex = PageAddresses()["index"]
	c.MessageNewRecipe = "We are going to add a new meal! Yay!"
	c.NavIndex = MessagesGlobal()["index"]
	c.TitlePage = MessagesGlobal()["titleNewRecipe"]
	c.TitleApp = MessagesGlobal()["nameApp"]
	return c
}

/*
Map initialiser functions return maps containing values, acting as constants
	- Maps' valused used both:
		- In backend, e.g. as paths to files and custom errors
		- In frontend, e.g. as text displayed to user and page metadata
*/

// FilePaths initialiser function returns map of all file paths
func FilePaths() map[string]string {
	var filePaths = map[string]string{
		"dataMeal":      "./data/meals/meals.json",
		"pageIndex":     "./html/home.html",
		"pageNewRecipe": "./html/newRecipe.html",
	}
	return filePaths
}

// MessagesIndex initialiser function returns map of home page messages
func MessagesIndex() map[string]string {
	var messagesIndex = map[string]string{
		"welcome":         "Welcome to the Meal Planner!",
		"instructions":    "What do you want to do?",
		"optionNewPlan":   "Make a new meal plan",
		"optionNewRecipe": "Add a new recipe",
	}
	return messagesIndex
}

// MessagesGlobal initialiser function returns map of global page messages
func MessagesGlobal() map[string]string {
	var messagesGlobal = map[string]string{
		"index":          "Home",
		"nameApp":        "Meal Planner",
		"titleNewRecipe": "New Recipe",
		"titleNewPlan":   "New Plan",
	}
	return messagesGlobal
}

// PageAddresses initialiser function returns map of all page addresses
func PageAddresses() map[string]string {
	var pageAddresses = map[string]string{
		"index":     "/",
		"newPlan":   "/new-plan",
		"newRecipe": "/new-recipe",
	}
	return pageAddresses
}

// CustomErrors initialiser function returns map of all custom errors
func CustomErrors() map[string]error {
	var customErrors = map[string]error{
		"invalidIndexToRemove": errors.New("cannot remove item as index invalid"),
	}
	return customErrors
}

/*
Struct definitions for structs used to pass values to application pages
*/

// ConstantsIndex struct for Home page constants
type ConstantsIndex struct {
	AddressIndex           string
	AddressNewPlan         string
	AddressNewRecipe       string
	MessageWelcome         string
	MessageInstructions    string
	ButtonOptionsNewPlan   string
	ButtonOptionsNewRecipe string
	NavIndex               string
	TitlePage              string
	TitleApp               string
}

// ConstantsNewPlan struct for New Plan page constants
type ConstantsNewPlan struct {
	AddressIndex string
	TitlePage    string
	TitleApp     string
}

// ConstantsNewRecipe struct for New Recipe page constants
type ConstantsNewRecipe struct {
	AddressIndex     string
	MessageNewRecipe string
	NavIndex         string
	NameApp          string
	TitlePage        string
	TitleApp         string
}

/*
DEPRECATE
*/

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
	menuMessages.MenuInstructions = "\nWhat do you want to do?\n"
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.QuitYes = "y"
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	return menuMessages
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
