package utils

import "errors"

// MenuMainOptions contains main menu options
var MenuMainOptions = map[string]string{
	"1": "Generate a meal plan",
	"2": "Add a new meal",
	"3": "Quit",
}

// FilePaths initialiser function returns map of all file paths
func FilePaths() map[string]string {
	var filePaths = map[string]string{
		"dataMeal":      "./data/meals/meals.json",
		"pageIndex":     "./html/home.html",
		"pageNewRecipe": "./html/newRecipe.html",
	}
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

// MessagesHome initialiser function returns map of home page messages
func MessagesHome() map[string]string {
	var messagesHome = map[string]string{
		"welcome":         "Welcome to the Meal Planner!",
		"instructions":    "What do you want to do?",
		"optionNewPlan":   "Make a new meal plan",
		"optionNewRecipe": "Add a new recipe",
	}
	return messagesHome
}

// MessagesGlobal initialiser function returns map of global page messages
func MessagesGlobal() map[string]string {
	var messagesGlobal = map[string]string{
		"home":           "Home",
		"nameApp":        "Meal Planner",
		"titleNewRecipe": "New Recipe",
		"titleNewPlan":   "New Plan",
	}
	return messagesGlobal
}

// PageAddresses initialiser function returns map of all page addresses
func PageAddresses() map[string]string {
	var pageAddresses = map[string]string{
		"home":      "/",
		"newPlan":   "/new-plan",
		"newRecipe": "/new-recipe",
	}
	return pageAddresses
}

// GetConstantsHome returns constant values for Home page
func GetConstantsHome() ConstantsHome {
	c := ConstantsHome{}
	c.AddressHome = PageAddresses()["home"]
	c.AddressNewPlan = PageAddresses()["newPlan"]
	c.AddressNewRecipe = PageAddresses()["newRecipe"]
	c.MessageWelcome = MessagesHome()["welcome"]
	c.MessageInstructions = MessagesHome()["instructions"]
	c.ButtonOptionsNewPlan = MessagesHome()["optionNewPlan"]
	c.ButtonOptionsNewRecipe = MessagesHome()["optionNewRecipe"]
	c.NavHome = MessagesGlobal()["home"]
	c.Title = MessagesGlobal()["nameApp"]
	return c
}

// GetConstantsNewPlan returns constant values for New Plan page
func GetConstantsNewPlan() ConstantsNewPlan {
	c := ConstantsNewPlan{}
	c.AddressHome = PageAddresses()["home"]
	c.Title = MessagesGlobal()["titleNewPlan"]
	return c
}

// GetConstantsNewRecipe returns constant values for New Recipe page
func GetConstantsNewRecipe() ConstantsNewRecipe {
	u := ConstantsNewRecipe{}
	u.AddressHome = PageAddresses()["home"]
	u.MessageNewRecipe = "We are going to add a new meal! Yay!"
	u.NavHome = MessagesGlobal()["home"]
	u.Title = MessagesGlobal()["titleNewRecipe"]
	return u
}

// ErrInvalidIndexToRemove returns error
var ErrInvalidIndexToRemove = errors.New("cannot remove item as index invalid")

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

// ConstantsHome struct for Home page constants
type ConstantsHome struct {
	AddressHome            string
	AddressNewPlan         string
	AddressNewRecipe       string
	MessageWelcome         string
	MessageInstructions    string
	ButtonOptionsNewPlan   string
	ButtonOptionsNewRecipe string
	NavHome                string
	Title                  string
}

// ConstantsNewPlan struct for New Plan page constants
type ConstantsNewPlan struct {
	AddressHome string
	Title       string
}

// ConstantsNewRecipe struct for New Recipe page constants
type ConstantsNewRecipe struct {
	AddressHome      string
	MessageNewRecipe string
	NavHome          string
	NameApp          string
	Title            string
}
