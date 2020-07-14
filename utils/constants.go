package utils

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
