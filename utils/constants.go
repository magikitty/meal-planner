package utils

/* decide if you prefer this pattern for storing messages in a struct
benefits:
1. messages are not publically exposed variables which can be accidentally changed in code
2. accessing is easy; simply call utils.MenuMessages(). the desired message, example:
	utils.MenuMessages().MenuInstructions for menu instructions
if you adopt this, be sure to:
	a. replace all calls to publically exposed vars with calls to MenuMessages()
	b. remove old publically exposed vars
*/

// MenuMessages function returning struct object with menu messages
func MenuMessages() menuMessages {
	menuMessages := menuMessages{}
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	menuMessages.MenuInstructions = "\nWhat do you want to do? Press the number of your choice."
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.ConfirmQuit = "Are you sure you want to quit? y/n"
	menuMessages.QuitYes = "y"
	menuMessages.QuitNo = "n"
	return menuMessages
}

type menuMessages struct {
	WelcomeMessage   string
	MenuInstructions string
	MenuMainOptions  map[string]string
	ConfirmQuit      string
	QuitYes          string
	QuitNo           string
}

// WelcomeMessage displays when user starts program
var WelcomeMessage = "Welcome to the Meal Planner!"

// MenuInstructions explains how to make a choice in the menu
var MenuInstructions = "\nWhat do you want to do? Press the number of your choice."

// MenuMainOptions contains options for the main menu
var MenuMainOptions = map[string]string{
	"1": "Generate a meal plan",
	"2": "Add a new meal",
	"3": "Quit",
}

// ConfirmQuit is a message for confirmation when quitting
var ConfirmQuit = "Are you sure you want to quit? y/n"

// QuitYes is user's confirmation to quit program
var QuitYes = "y"

// QuitNo is user's confirmation not to quit program
var QuitNo = "n"

// QuitProgram is a bool that is set to true when user wants to quit program
var QuitProgram = false
